package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	mqtt "github.com/mochi-mqtt/server/v2"

	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
)
const (
	testSub = "J200002335C2S"
	testPub = "J200002335S2C"
)
var errTestHook = errors.New("error")

type hooks struct {
	mqtt.HookBase
	fail bool
	err error
}

func (h *hooks) OnPacketRead(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	if h.fail {
		if h.err != nil {
			return pk, h.err
		}

		return pk, errTestHook
	}
	fmt.Println(cl.ID, pk.Connect.Username, pk.Connect.Password)
	if pk.Connect.Password == nil {
		newPassword := "new_password"
		pk.Connect.Password = []byte(newPassword)
	}
	
	fmt.Println(cl.ID, pk.Connect.Username, pk.Connect.Password)
	return pk, nil
}

func (h *hooks) ID() string {
	return "modified"
}

func (h *hooks) Init(config any) error {
	if config != nil {
		return errTestHook
	}
	return nil
}

func (h *hooks) Provides(b byte) bool {
	return true
}

func (h *hooks) Stop() error {
	if h.fail {
		return errTestHook
	}

	return nil
}

func (h *hooks) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	if h.fail {
		return errTestHook
	}

	return nil
}

func (h *hooks) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	return true
}

func (h *hooks) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	return true
}

func (h *hooks) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	if h.fail {
		if h.err != nil {
			return pk, h.err
		}

		return pk, errTestHook
	}

	return pk, nil
}


func (h *hooks) OnAuthPacket(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	if h.fail {
		if h.err != nil {
			return pk, h.err
		}

		return pk, errTestHook
	}

	return pk, nil
}

func (h *hooks) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	return pk
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
	err := server.AddHook(new(hooks), nil)
	if err != nil {
		log.Fatal(err)
	}

	tcp := listeners.NewTCP(listeners.Config{
		ID: "t2",
		Address: ":1883",
	})

	callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
		server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
	}

	server.Subscribe(testSub, 1, callbackFn)

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