package main

import (
	"log"
	"os"
	"time"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
)

const testTopic = "/test"

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	go func() {
		<-sigs
		done <- true
	}()

	options := &mqtt.Options{
		InlineClient: true,
	}

	server := mqtt.New(options)

	err := server.AddHook(new(auth.AllowHook), nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		// Inline subscriptions can also receive retained messages on subscription.
		_ = server.Publish("direct/retained", []byte("retained message"), true, 0)
		_ = server.Publish("direct/alternate/retained", []byte("some other retained message"), true, 0)

		// Subscribe to a filter and handle any received messages via a callback function.
		callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
			server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
		}
		server.Log.Info("inline client subscribing")
		_ = server.Subscribe(testTopic, 1, callbackFn)
		_ = server.Subscribe(testTopic, 2, callbackFn)
	}()

	go func() {
		for range time.Tick(time.Second * 3) {
			err := server.Publish("direct/publish", []byte("scheduled message"), false, 0)
			if err != nil {
				server.Log.Error("server.Publish", "error", err)
			}
			server.Log.Info("main.go issued direct message to direct/publish")
		}
	}()

	tcp := listeners.NewTCP(listeners.Config{
		ID:      "t1",
		Address: ":1883",
	})
	err = server.AddListener(tcp)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-done
	server.Log.Warn("caught signal, stopping...")
	_ = server.Close()
	server.Log.Info("main.go finished")
}
