package main

import (
	"crypto/aes"
	"crypto/cipher"
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
	testSub = "J23P000078C2S"
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
		// hexPayload := hex.EncodeToString([]byte(pk.Payload))
		decrptedPaylod := Decrypt(pk.Payload)
		fmt.Println(pk.Payload)
		fmt.Println(decrptedPaylod)
		if err != nil {
			fmt.Println("error decrypting payload: ", err)
		}
		server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", reflect.TypeOf(pk.Payload) )
		// fmt.Println("payload type is", reflect.TypeOf(pk.Payload))
		// fmt.Println("payload is", decrptedPaylod)
		// fmt.Println("length of payload ", len(decrptedPaylod))
		fmt.Println("decrpt payload is", decrptedPaylod)
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

func Decrypt(plaintext []byte) string {
	bKey := []byte("69aF7&3KY0_kk89@")
	bIV := []byte("420#abA%,ZfE79@M")
	
	bPlaintext := []byte(plaintext)
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}