package main

import (
	Network "Network"
	"fmt"
	"log"
	"net"
	LobbyPacket "packet_lobby"

	"github.com/golang/protobuf/proto"
)

//ClientSession is ...
type ClientSession struct {
	Network.Session
}

// func (session *ClientSession) Init(conn net.Conn) {
// 	session.Session.Conn = Network.CreateSession(conn)
// }

//-----------------------------------------------------------------------------
func (session *ClientSession) reqestLogin(userName string) {

	req := &LobbyPacket.LoginReq{}
	req.Name = userName
	packet, err := proto.Marshal(req)
	processError(err)

	writeSize, err := session.Send(packet)
	processError(err)
	log.Printf("send : %d\n", writeSize)
	log.Printf(req.String())

	recvBuffer := make([]byte, 4096)
	readn, err := session.Recv(recvBuffer)
	processError(err)
	log.Printf("recv : %d\n", readn)

	loginRes := &LobbyPacket.LoginRes{}
	err = proto.Unmarshal(recvBuffer[:readn], loginRes)
	processError(err)

	//log.Printf("Recv : loginRes : %d\n", loginRes.GetRetCode())
	log.Printf(loginRes.String())
}

func recoverError() {
	err := recover()
	log.Fatalln(err)
}

// processError call panic when error occurs
func processError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	defer recoverError()
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	processError(err)

	clientSession := ClientSession{}
	clientSession.Session.InitConnection(conn)

	defer clientSession.Session.Conn.Close()

	fmt.Println("input name")

	for {
		clientSession.reqestLogin("hello_abc")
	}

	// //recv
	// go func(client *ClientSession) {
	// 	buffer := make([]byte, 4096)
	// 	for {
	// 		//readSize, err := client.Recv(buffer)
	// 		if !client.IsConnected() {
	// 			return
	// 		}

	// 		readSize, _ := client.Recv(buffer)
	// 		//processError(err)

	// 		if readSize > 0 {
	// 			client.reqestLogin("hello")
	// 		}
	// 	}
	// }(&clientSession)

	// //send
	// //recv에서 처리하고 보낼 패킷을 컨테이너에 저장해서 여기서 보내줘야 정석
	// //현재는 recv하면 바로 처리해서 send한다
	// go func(client *ClientSession) {

	// }(&clientSession)
}
