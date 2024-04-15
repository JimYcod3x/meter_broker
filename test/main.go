package main

import (
	// "encoding/hex"
	"fmt"
	// "net"
	"meter_broker/meters"
)


// 	func mainencrypt() {
//     // Hex string
//     hexString := "6a004a23500000780000000000000000"

//     // Convert hex string to byte slice
//     hexBytes, _ := hex.DecodeString(hexString)

//     // Extract the first byte
//     firstByte := hexBytes[0]

//     // Extract the first 3 bits using left shift
//     first3Bits := firstByte >> (8 - 3) // Shift right by 5 bits to keep only the first 3 bits
// 		then4Bits := firstByte >> (8 - 5)
// 		net.ListenPacket(network string, address string)
//     // Display the first 3 bits
//     fmt.Printf("First 3 bits in binary: %03b\n", first3Bits)
//     fmt.Printf("Then 4 bits in binary: %03b\n", then4Bits)
// }


func main() {
	fmt.Printf("%03b\n", meters.ElectricityMeter)
	fmt.Printf("%03b\n", meters.WaterMeter)
	fmt.Printf("%03b\n", meters.IoT)
	fmt.Printf("%03b\n", meters.GasMeter)
	fmt.Printf("%03b\n", meters.HeatMeter)
	fmt.Printf("%03b\n", meters.PV)
} 