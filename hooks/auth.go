package hooks

import (
	"errors"
	"fmt"
	"reflect"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

var errTestHook = errors.New("error")

type Hooks struct {
	mqtt.HookBase
	fail bool
	err error
}

func (h *Hooks) OnPacketRead(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	if h.fail {
		if h.err != nil {
			fmt.Println("fail connect with error")
			return pk, h.err
		}
		fmt.Println("fail connect without error")
		return pk, errTestHook
	}
	// fmt.Printf("client Properties : %v\n client Net : %v\n client ID : %v\n", cl.Properties, cl.Net, cl.ID)
	// fmt.Printf("Packet fields - Connect: %v\n, Properties: %v\n, Payload: %v\n, ReasonCodes: %v\n, Filters: %v\n, TopicName: %v\n, Origin: %v\n, FixedHeader: %v\n, Created: %v\n, Expiry: %v\n, Mods: %v\n, PacketID: %v\n, ProtocolVersion: %v\n, SessionPresent: %v\n, ReasonCode: %v\n, ReservedBit: %v\n, Ignore: %v\n, UsernameFlag: %v\n, PasswordFlag: %v\n",
		// pk.Connect, pk.Properties, pk.Payload, pk.ReasonCodes, pk.Filters, pk.TopicName, pk.Origin, pk.FixedHeader, pk.Created, pk.Expiry, pk.Mods, pk.PacketID, pk.ProtocolVersion, pk.SessionPresent, pk.ReasonCode, pk.ReservedBit, pk.Ignore, pk.Connect.UsernameFlag, pk.Connect.PasswordFlag)
	// if cl.ID == nil {
	// 	cl.ID = []byte("new_id")
	// }
	// fmt.Println("payload type is", reflect.TypeOf(pk.Payload))

	if string(pk.Connect.Username) == "" {
		pk.Connect.UsernameFlag = false
	}
	if string(pk.Connect.Password) == "" {
		pk.Connect.PasswordFlag = false
	}
	// fmt.Printf("username: %v\n password: %v\n",	 pk.Connect.Username, pk.Connect.Password)
	return pk, nil
}

func (h *Hooks) ID() string {
	return "modified"
}

func (h *Hooks) Init(config any) error {
	if config != nil {
		return errTestHook
	}
	return nil
}

func (h *Hooks) Provides(b byte) bool {
	return true
}

func (h *Hooks) Stop() error {
	if h.fail {
		return errTestHook
	}

	return nil
}

func (h *Hooks) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	if h.fail {
		return errTestHook
	}

	return nil
}

func (h *Hooks) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	return true
}

func (h *Hooks) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	return true
}

func (h *Hooks) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	if h.fail {
		if h.err != nil {
			return pk, h.err
		}

		return pk, errTestHook
	}
	
	
	return pk, nil
}


func (h *Hooks) OnAuthPacket(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	if h.fail {
		if h.err != nil {
			return pk, h.err
		}

		return pk, errTestHook
	}

	return pk, nil
}

func (h *Hooks) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	// fmt.Println("payload is", pk.Payload)
	// fmt.Println("payload type is", reflect.TypeOf(pk.Payload))
	return pk
}