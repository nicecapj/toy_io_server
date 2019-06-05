package network

import (
	PROTOCOL "packet_protocol"
)

type Header struct {
	PacketSize int32
	PacketID   PROTOCOL.ProtocolID
}

const PacketHeaderLen = 8 //성능을 위해 sizeof하지 않는다
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
