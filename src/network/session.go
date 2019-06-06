package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
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

	Name string
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
		log.Panicln(err)
	}

	var header Header
	header.PacketID = protocolID
	header.PacketSize = PacketHeaderLen + int32(len(body))

	buffer, err := session.SetHeader(header)
	if err != nil {
		log.Panicln("set header")
	}

	defer func() {
		buffer.Reset()
		session.poolBuffer.Put(buffer)
	}()

	buffer.Write(body)

	session.Send(buffer.Bytes())
}

func (session *Session) SetHeader(header Header) (*bytes.Buffer, error) {
	buffer := session.poolBuffer.Get().(*bytes.Buffer)

	err := binary.Write(buffer, binary.LittleEndian, header)
	if err != nil {
		log.Panicln("write header")
	}

	return buffer, err
}

// Recv ...
//func (session *Session) HandlePacket(packet []byte) {
func (session *Session) HandlePacket(bufferArray []byte) {

	//풀을 여기서는 안써도 되지 않을까. 괜히 복사나 한번 더 일어나지
	// buffer := session.poolBuffer.Get().(*bytes.Buffer)
	// defer func() {
	// 	buffer.Reset()
	// 	session.poolBuffer.Put(buffer)
	// }()
	// buffer.Write(packet[:MaxPacketSize])

	// bufferArray := buffer.Bytes()
	header, err := GetHeader(bufferArray[:MaxPacketSize])
	if err != nil {
		log.Panicln("read header")
	}

	switch header.PacketID {
	case PROTOCOL.ProtocolID_LoginReq:
		{
			req := &LobbyPacket.LoginReq{}
			err = proto.Unmarshal(bufferArray[PacketHeaderLen:MaxPacketSize], req)
			log.Printf("%s\n", req.String())

			session.OnHandleLoginReq(req)
		}
	case PROTOCOL.ProtocolID_LoginRes:
		{
			res := &LobbyPacket.LoginRes{}
			err = proto.Unmarshal(bufferArray[PacketHeaderLen:MaxPacketSize], res)
			log.Printf("%s\n", res.String())
		}

	}
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
	SessionManager.RemoveUser(session.Name)
}

func (session *Session) OnHandleLoginReq(req *LobbyPacket.LoginReq) {
	sessionManager := GetSessionManager()

	name := req.GetName()
	session.Name = name

	res := &LobbyPacket.LoginRes{}
	if sessionManager.FindUser(name) {
		res.RetCode = ReturnCode.ReturnCode_retExist
		res.Uid = sessionManager.GetUID(name)
	} else {
		sessionManager.AddUser(name)
		res.RetCode = ReturnCode.ReturnCode_retOK
		res.Uid = sessionManager.GetUID(name)
	}
	session.SendPacket(PROTOCOL.ProtocolID_LoginRes, res)
}
