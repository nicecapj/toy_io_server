package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	PROTOCOL "packet_protocol"
	"reflect"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

// Session is basic struct for network communication
type Session struct {
	sync.Mutex
	Conn        net.Conn
	isConnected bool
	PoolBuffer  sync.Pool

	recvChan chan Packet
	sendChan chan Packet
	exitChan chan struct{}

	Name string
	UID  int64
}

//InitConnection ...
func (session *Session) InitConnection(conn net.Conn) {
	session.Conn = conn
	session.isConnected = true

	session.recvChan = make(chan Packet, 2)
	session.sendChan = make(chan Packet, 2)
	session.exitChan = make(chan struct{})

	session.PoolBuffer = sync.Pool{
		New: func() interface{} {
			nb := new(bytes.Buffer)
			return nb
		},
	}
}

// SendPacket ...
func (session *Session) SendPacket(protocolID PROTOCOL.ProtocolID, pb proto.Message) {

	if !session.IsConnected() {
		return
	}

	body, err := proto.Marshal(pb)
	if err != nil {
		log.Panicln(err)
	}

	log.Print(reflect.TypeOf(pb))

	log.Printf("Send : %s\n", pb.String())

	var header Header
	header.PacketID = protocolID
	header.PacketSize = PacketHeaderLen + int32(len(body))

	buffer, err := session.SetHeader(header)
	if err != nil {
		log.Panicln("set header")
	}

	defer func() {
		buffer.Reset()
		session.PoolBuffer.Put(buffer)
	}()

	buffer.Write(body)

	packet := Packet{}
	packet.Header.PacketID = protocolID
	//packet.Header.PacketSize = int32(len(buffer))
	packet.MessageStream = buffer.Bytes()

	//channel style
	//session.sendChan <- packet

	//immediately process style
	session.Send(buffer.Bytes())
}

// SetHeader ...
func (session *Session) SetHeader(header Header) (*bytes.Buffer, error) {
	buffer := session.PoolBuffer.Get().(*bytes.Buffer)

	err := binary.Write(buffer, binary.LittleEndian, header)
	if err != nil {
		log.Panicln("write header")
	}

	return buffer, err
}

// DispatchPacket ...
func (session *Session) DispatchPacket(protocolID PROTOCOL.ProtocolID, buffer []byte) {
	packet := Packet{}
	packet.Header.PacketID = protocolID
	packet.Header.PacketSize = int32(len(buffer))
	packet.MessageStream = buffer[:]
	session.recvChan <- packet
}

// GetHeader ...
func GetHeader(stream []byte) (Header, error) {

	var header Header
	buffer := bytes.NewReader(stream[:PacketHeaderLen])

	err := binary.Read(buffer, binary.LittleEndian, &header)
	if err != nil {
		log.Panicln("read header")
	}

	return header, err
}

//Send ...
func (session *Session) Send(packet []byte) (int, error) {
	readSize, err := session.Conn.Write([]byte(packet))
	if err != nil {
		log.Panicln(err)
	}
	return readSize, err
}

// Recv ...
func (session *Session) Recv(packet []byte) (int, error) {
	readSize, err := session.Conn.Read([]byte(packet))

	if err != nil {
		session.isConnected = false

		if err == io.ErrClosedPipe {

		} else if err == io.EOF {

		} else {
			//특정 상황에 에러가 알고 싶으면 아래 주석 활성화
			//log.Panic(err)
			fmt.Println(err)
		}
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

	SessionManager := GetSessionManager()
	SessionManager.RemoveSession(session)
}

func (session *Session) OnTick(delta time.Duration) {
	//log.Printf("Tick : %s", session.Name)

	//sync lock! it make stop sessionManager`s ontick

	// //channel style
	// for sendPacket := range session.`sendChan` {
	// 	log.Printf("channel send : %d\n", sendPacket.Header.PacketID)
	// 	session.Send(sendPacket.MessageStream)
	// }

	// for recvPacket := range session.recvChan {
	// 	log.Printf("channel recv : %d\n", recvPacket.Header.PacketID)
	// }
}
