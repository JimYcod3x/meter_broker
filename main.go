package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/mochi-mqtt/server/v2"

	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
)

// const testSub = "J200002335C2S"
const testPub = "J200002335S2C"
var server *mqtt.Server



type modifiedHookBase struct {
	mqtt.HookBase
}

func (h *modifiedHookBase) ID() string {
	return "mqtt-hook"
}

func (h *modifiedHookBase) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnectAuthenticate,
		mqtt.OnACLCheck,
		mqtt.OnSessionEstablish,
		mqtt.OnDisconnect,
		mqtt.OnPublished,
	}, []byte{b})
}

func (h *modifiedHookBase) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	username := string(pk.Connect.Username)
	// password := string(pk.Connect.Password)
	if username == "" || len(username) == 0 {
		return true
	}
	return true
}

func (h *modifiedHookBase) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	username := string(cl.Properties.Username)
	if username == "" || len(username) == 0 {
		return true
	}
	return true
}

func (h *modifiedHookBase) OnSessionEstablish(cl *mqtt.Client, pk packets.Packet) {
	username := string(cl.Properties.Username)
	if username == "" || len(username) == 0 {
		return
	}
}

func (h *modifiedHookBase) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	if err != nil {
		h.Log.Info("client disconnected", "client", cl.ID, "expire", expire, "error", err)
	} else {
		h.Log.Info("client disconnected", "client", cl.ID, "expire", expire)
	}

}

func (h *modifiedHookBase) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	fmt.Printf("mqtt server OnPublished info topic: %s, msg: %s", pk.TopicName, string(pk.Payload))
}

// func PublishMsg(topic string, msg []byte) bool {
// 	err := server.Publish("J200002335S2C", msg, false, 0)
// 	if err != nil {
// 		log.Fatal("server.Publish", "error", err)
// 		return false
// 	}
// 	return true
// }

func SubscribeMsg(topic string, subscriptionId int, callback func(topic string, msg []byte)) {
	callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
		server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
		callback(pk.TopicName, pk.Payload)
	}
	err := server.Subscribe(topic, subscriptionId, callbackFn)
	if err != nil {
		log.Fatal("server.Subscribe", "error", err)
	}
}

func UnsubscribeMsg(topic string, subscriptionId int) {
	err := server.Unsubscribe(topic, subscriptionId)
	if err != nil {
		log.Fatal("server.Unsubscribe", "error", err)
	}
}

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

	err := server.AddHook(new(modifiedHookBase), nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		// Inline subscriptions can also receive retained messages on subscription.
		_ = server.Publish(testPub, []byte("retained message"), true, 0)
		// _ = server.Publish("direct/alternate/retained", []byte("some other retained message"), true, 0)

		// Subscribe to a filter and handle any received messages via a callback function.
		// callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
		// 	server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
		// }
		// server.Log.Info("inline client subscribing")
		// _ = server.Subscribe(testTopic, 1, callbackFn)
		// _ = server.Subscribe(testTopic, 2, callbackFn)
	}()

	go func() {
		for range time.Tick(time.Second * 3) {
			err := server.Publish(testPub, []byte("scheduled message"), false, 0)
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
