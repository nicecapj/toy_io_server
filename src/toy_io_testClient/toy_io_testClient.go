package main

import (
	Network "Network"
	"fmt"
	"log"
	"math/rand"
	"net"
	LobbyPacket "packet_lobby"
	PROTOCOL "packet_protocol"
	"time"
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
	session.SendPacket(PROTOCOL.ProtocolID_LoginReq, req)

	recvBuffer := make([]byte, 4096)
	readn, err := session.Recv(recvBuffer)
	processError(err)
	log.Printf("recv : %d\n", readn)
	session.HandlePacket(recvBuffer)

	// processError(err)
	// log.Printf("recv : %d\n", readn)

	// loginRes := &LobbyPacket.LoginRes{}
	// err = proto.Unmarshal(recvBuffer[:readn], loginRes)
	// processError(err)

	// //log.Printf("Recv : loginRes : %d\n", loginRes.GetRetCode())
	// log.Printf("%s\n", loginRes.String())
}

func recoverError() {
	err := recover()
	log.Fatalln(err)
}

//GetRandomName ...
func GetRandomName() string {
	alpha := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "z", "y", "z",
	}

	seedValue := time.Now().Unix()
	rand.Seed(seedValue)

	name := ""
	for i := 0; i < 10; i++ {
		name = name + alpha[rand.Intn(len(alpha))]
	}

	return name
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
		name := GetRandomName()
		clientSession.reqestLogin(name)
		//time.Sleep(1 * time.Second)
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
