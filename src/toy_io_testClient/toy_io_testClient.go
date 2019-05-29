package main

import (
	"fmt"
	"log"
	"net"
	. "packet_lobby"

	"github.com/golang/protobuf/proto"
)

//-----------------------------------------------------------------------------
type Session struct {
	conn net.Conn
}

func CreateSession(conn net.Conn) *Session {
	newSession := new(Session)
	newSession.conn = conn
	return newSession
}

func (session *Session) ReqestLogin(userName string) {
	packet := &LoginReq{
		Name: userName,
	}

	data, err := proto.Marshal(packet)
	if err != nil {
		log.Fatal("marshaling error: ", err)
		panic(err)
	}
	session.conn.Write([]byte(data))
}

func (session *Session) Send(packet []byte) (int, error) {
	readSize, err := session.conn.Write([]byte(packet))
	return readSize, err
}

func (session *Session) Recv(packet []byte) (int, error) {
	readSize, err := session.conn.Read([]byte(packet))
	return readSize, err
}

//-----------------------------------------------------------------------------

func RecoverError() {
	err := recover()
	log.Fatalln(err)
}

func ProcessError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	defer RecoverError()
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	ProcessError(err)

	session := CreateSession(conn)

	fmt.Println("input name")

	for {
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Println("To Send Message :")
		// text, err := reader.ReadString('\n')
		// ProcessError(err)

		text := "hello"

		req := LoginReq{}
		req.Name = text
		packet, err := proto.Marshal(&req)
		ProcessError(err)

		writeSize, err := session.Send(packet)
		ProcessError(err)
		log.Printf("send : %d\n", writeSize)

		recvBuffer := make([]byte, 4096)
		readn, err := session.Recv(recvBuffer)

		loginRes := LoginRes{}
		proto.Unmarshal(recvBuffer[:readn], &loginRes)

		log.Printf("Recv : loginRes : %d\n", loginRes.GetRetCode())
	}
}
