package main

import (
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


// func decrypt(hexPayload, hexKey, hexIV string) ([]byte, error) {
// 	payload, err := hex.DecodeString(hexPayload)
// 	if err != nil {
// 			return nil, fmt.Errorf("error decoding payload: %v", err)
// 	}

// 	key, err := hex.DecodeString(hexKey)
// 	if err != nil {
// 			return nil, fmt.Errorf("error decoding key: %v", err)
// 	}

// 	iv, err := hex.DecodeString(hexIV)
// 	if err != nil {
// 			return nil, fmt.Errorf("error decoding IV: %v", err)
// 	}

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 			return nil, err
// 	}

// 	if len(payload) < aes.BlockSize {
// 			return nil, fmt.Errorf("ciphertext too short")
// 	}

// 	mode := cipher.NewCBCDecrypter(block, iv)

// 	mode.CryptBlocks(payload, payload)

// 	// Remove PKCS#7 padding
// 	pad := int(payload[len(payload)-1])
// 	if pad < 1 || pad > aes.BlockSize {
// 			return nil, fmt.Errorf("invalid padding")
// 	}
// 	payload = payload[:len(payload)-pad]

// 	return payload, nil
// }

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
		// decrptedPaylod, err := decrypt(hexPayload, "69aF7&3KY0_kk89@", "420#abA%,ZfE79@M")
		if err != nil {
			fmt.Println("error decrypting payload: ", err)
		}
		server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", reflect.TypeOf(pk.Payload) )
		// fmt.Println("payload type is", reflect.TypeOf(pk.Payload))
		// fmt.Println("payload is", hexPayload)
		// fmt.Println("length of payload ", len(hexPayload))
		// fmt.Println("decrpt payload is", decrptedPaylod)
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