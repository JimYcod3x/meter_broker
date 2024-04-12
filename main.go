package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"meter_broker/hooks"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	mqtt "github.com/mochi-mqtt/server/v2"

	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
)
const (
	testSub = "J230008542C2S"
	testPub = "J200002335S2C"
)

var server = mqtt.New(&mqtt.Options{
	InlineClient: true,
})


func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
	
	err := server.AddHook(new(hooks.Hooks), nil)
	if err != nil {
		log.Fatal(err)
	}

	tcp := listeners.NewTCP(listeners.Config{
		ID: "t2",
		Address: ":1883",
	})



	err = server.Subscribe(testSub, 1, func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
		server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
		fmt.Println("payload type is", reflect.TypeOf(pk.Payload))
		fmt.Println("payload is", hex.EncodeToString([]byte(pk.Payload)))
		fmt.Println("payload is", pk.Payload)
	})

	if err != nil {
		log.Fatal(err)
	}

	err = server.AddListener(tcp)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}



	<- done
	server.Log.Warn("caught signal, stopping...")
	_ = server.Close()
	server.Log.Info("main.go finished")
}