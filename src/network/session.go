package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	PROTOCOL "packet_protocol"
	"sync"

	"github.com/golang/protobuf/proto"
)

// Session is basic struct for network communication
type Session struct {
	sync.Mutex
	Conn        net.Conn
	isConnected bool
	PoolBuffer  sync.Pool

	Name string
}

//InitConnection ...
func (this *Session) InitConnection(conn net.Conn) {
	this.Conn = conn
	this.isConnected = true

	this.PoolBuffer = sync.Pool{
		New: func() interface{} {
			nb := new(bytes.Buffer)
			return nb
		},
	}
}

func (this *Session) SendPacket(protocolID PROTOCOL.ProtocolID, pb proto.Message) {

	body, err := proto.Marshal(pb)
	if err != nil {
		log.Panicln(err)
	}

	var header Header
	header.PacketID = protocolID
	header.PacketSize = PacketHeaderLen + int32(len(body))

	buffer, err := this.SetHeader(header)
	if err != nil {
		log.Panicln("set header")
	}

	defer func() {
		buffer.Reset()
		this.PoolBuffer.Put(buffer)
	}()

	buffer.Write(body)

	this.Send(buffer.Bytes())
}

func (this *Session) SetHeader(header Header) (*bytes.Buffer, error) {
	buffer := this.PoolBuffer.Get().(*bytes.Buffer)

	err := binary.Write(buffer, binary.LittleEndian, header)
	if err != nil {
		log.Panicln("write header")
	}

	return buffer, err
}

func (this *Session) DispatchPacket(protocolID PROTOCOL.ProtocolID, buffer []byte) {
}

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
func (this *Session) Send(packet []byte) (int, error) {
	readSize, err := this.Conn.Write([]byte(packet))
	if err != nil {
		log.Panicln(err)
	}
	return readSize, err
}

// Recv ...
func (this *Session) Recv(packet []byte) (int, error) {
	readSize, err := this.Conn.Read([]byte(packet))

	if err != nil {
		this.isConnected = false

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
func (this *Session) IsConnected() bool {
	//conn에 이미 이런게 있지 않을까?
	return this.isConnected
}

//Close ...
func (this *Session) Close() {
	this.Conn.Close()

	SessionManager := GetSessionManager()
	SessionManager.RemoveSession(this)
}
