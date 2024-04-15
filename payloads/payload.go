package payload

import "meter_broker/meters"

type Payload struct {
	Meter meters.Meter
	DataPacket string
}