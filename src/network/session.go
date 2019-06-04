package network

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	LobbyPacket "packet_lobby"
	PROTOCOL "packet_protocol"
	ReturnCode "packet_returncode"
	"sync"

	"github.com/golang/protobuf/proto"
)

// Session is basic struct for network communication
type Session struct {
	Conn        net.Conn
	isConnected bool
	poolBuffer  sync.Pool
}

// CreateSession make new session
func CreateSession(Conn net.Conn) *Session {
	//newSession := new(Session)
	newSession := &Session{}

	if newSession != nil {

		log.Printf("New session : %s", Conn.RemoteAddr())
		//log.Printf("New session : %s", Conn.LocalAddr())

		newSession.InitConnection(Conn)
	}

	return newSession
}

//InitConnection ...
func (session *Session) InitConnection(conn net.Conn) {
	session.Conn = conn
	session.isConnected = true

	session.poolBuffer = sync.Pool{
		New: func() interface{} {
			nb := new(bytes.Buffer)
			return nb
		},
	}
}

func (session *Session) SendPacket(protocolID PROTOCOL.ProtocolID, pb proto.Message) {

	body, err := proto.Marshal(pb)
	if err != nil {
		panic(err)
	}

	var header Header
	header.packetID = protocolID
	header.packetSize = PacketHeaderLen + int32(len(body))

	buffer, err := session.SetHeader(header)
	if err != nil {
		log.Fatalln("set header")
	}

	buffer.Write(body)

	session.Send(buffer.Bytes())
}

func (session *Session) SetHeader(header Header) (*bytes.Buffer, error) {
	buffer := session.poolBuffer.Get().(*bytes.Buffer)

	err := binary.Write(buffer, binary.LittleEndian, header)
	if err != nil {
		log.Fatalln("write header")
	}

	return buffer, err
}

// Recv ...
//func (session *Session) HandlePacket() (int, error) {
func (session *Session) HandlePacket(packet []byte) {
	header, err := GetHeader(packet)
	if err != nil {
		log.Fatalln("read header")
	}

	switch header.packetID {
	case PROTOCOL.ProtocolID_LoginReq:
		{
			req := &LobbyPacket.LoginReq{}
			err = proto.Unmarshal(packet[PacketHeaderLen:MaxPacketSize], req)
			log.Printf("%s\n", req.String)

			res := &LobbyPacket.LoginRes{}
			res.RetCode = ReturnCode.ReturnCode_retExist
			res.Uid = 1234 //sessionManager.GetUID(name)
			session.SendPacket(PROTOCOL.ProtocolID_LoginReS, res)
		}
	case PROTOCOL.ProtocolID_LoginReS:
		{
			res := &LobbyPacket.LoginRes{}
			err = proto.Unmarshal(packet[PacketHeaderLen:MaxPacketSize], res)
			log.Printf("%s\n", res.String)
		}

	}
}

func GetHeader(stream []byte) (Header, error) {
	var header Header
	buffer := bytes.NewReader(stream[:PacketHeaderLen])

	err := binary.Read(buffer, binary.LittleEndian, &header)
	if err != nil {
		log.Fatalln("read header")
	}

	return header, err
}

//Send ...
func (session *Session) Send(packet []byte) (int, error) {
	readSize, err := session.Conn.Write([]byte(packet))
	if err != nil {
		log.Fatalln(err)
	}
	return readSize, err
}

// Recv ...
func (session *Session) Recv(packet []byte) (int, error) {
	readSize, err := session.Conn.Read([]byte(packet))

	if err != nil {
		session.isConnected = false
		log.Fatalln(err)
	}

	return readSize, err
}

// IsConnected ...
func (session *Session) IsConnected() bool {
	//conn에 이미 이런게 있지 않을까?
	return session.isConnected
}

//Close ...
func (session *Session) Close() {
	session.Conn.Close()
}
