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
	newPassword := "new_password"
	pk.Connect.Password = []byte(newPassword)
	return pk, nil
}

func (h *hooks) ID() string {
	return "myHook"
}
func (h *hooks) Provides(b byte) bool {
	return true
}

func (h *hooks) Init(config interface{}) error {
	fmt.Println("init hook")
	return nil
}

func (h *hooks) OnStarted() {
	fmt.Println("client connected")
}

func (h *hooks) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	fmt.Println("client test connected")
	return true
}

func (h *hooks) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool     { return true }

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