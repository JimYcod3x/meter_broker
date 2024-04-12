package main

import (
	"net"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/mochi-mqtt/server/v2/system"
)




func Subscribe() {
	client := server.Subscribe(filter string, subscriptionId int, handler mqtt.InlineSubFn)
	
}


	if c != nil {
		cl.Net = ClientConnection{
			Conn:   c,
			bconn:  bufio.NewReaderSize(c, o.options.ClientNetReadBufferSize),
			Remote: c.RemoteAddr().String(),
		}
	}

	return cl
}

func  callbackFn(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
	server.Log.Info("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
}

