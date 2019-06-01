package main

import (
	Network "Network"
	"fmt"
	"log"
	"net"
	LobbyPacket "packet_lobby"

	"github.com/golang/protobuf/proto"
)

//-----------------------------------------------------------------------------
type ClientSession struct {
	Network.Session
}

//-----------------------------------------------------------------------------
func (session *ClientSession) ReqestLogin(userName string) {
	packet := &LobbyPacket.LoginReq{
		Name: userName,
	}

	data, err := proto.Marshal(packet)
	if err != nil {
		log.Fatal("marshaling error: ", err)
		panic(err)
	}
	session.Conn.Write([]byte(data))
}

//-----------------------------------------------------------------------------
// RecoverError recover panic and print err
func RecoverError() {
	err := recover()
	log.Fatalln(err)
}

// ProcessError call panic when error occurs
func ProcessError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	defer RecoverError()
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	ProcessError(err)

	session := Network.CreateSession(conn)

	fmt.Println("input name")

	for {
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Println("To Send Message :")
		// text, err := reader.ReadString('\n')
		// ProcessError(err)

		text := "hello"

		req := LobbyPacket.LoginReq{}
		req.Name = text
		packet, err := proto.Marshal(&req)
		ProcessError(err)

		writeSize, err := session.Send(packet)
		ProcessError(err)
		log.Printf("send : %d\n", writeSize)

		recvBuffer := make([]byte, 4096)
		readn, err := session.Recv(recvBuffer)

		loginRes := LobbyPacket.LoginRes{}
		proto.Unmarshal(recvBuffer[:readn], &loginRes)

		log.Printf("Recv : loginRes : %d\n", loginRes.GetRetCode())
	}
}
