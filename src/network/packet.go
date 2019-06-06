package network

import (
	PROTOCOL "packet_protocol"
)

//Header is packet header
type Header struct {
	PacketSize int32
	PacketID   PROTOCOL.ProtocolID
}

//PacketHeaderLen used for performance(call sizeof(Header))
const PacketHeaderLen = 8

//MaxPacketSize is limited size of packet
const MaxPacketSize = 4096

//MarshalPacket...
//func SetHeader(PacketID PROTOCOL.ProtocolID, ) (bytes.Buffer, error) {
// 	header := Header{}
// 	header.PacketID = PacketID

// 	headerSize := unsafe.Sizeof(header)
// 	packetBuffer := make([]byte, headerSize)

// 	packetStream, err := proto.Marshal(pb)
// 	packetBuffer = append(packetBuffer, packetStream...)

// 	return packetBuffer, err
//}

// //UnmarshalPacket...
// func GetHeader(packetStream []byte) {

// }
