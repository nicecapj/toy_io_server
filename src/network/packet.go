package network

import (
	PROTOCOL "packet_protocol"
)

type Header struct {
	packetID   PROTOCOL.ProtocolID
	packetSize int32
}

const PacketHeaderLen = 8 //성능을 위해 sizeof하지 않는다
const MaxPacketSize = 4096

//MarshalPacket...
//func SetHeader(packetID PROTOCOL.ProtocolID, ) (bytes.Buffer, error) {
// 	header := Header{}
// 	header.packetID = packetID

// 	headerSize := unsafe.Sizeof(header)
// 	packetBuffer := make([]byte, headerSize)

// 	packetStream, err := proto.Marshal(pb)
// 	packetBuffer = append(packetBuffer, packetStream...)

// 	return packetBuffer, err
//}

// //UnmarshalPacket...
// func GetHeader(packetStream []byte) {

// }
