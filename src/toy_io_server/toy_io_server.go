package main

import (
	Network "Network"
	"fmt"
	"log"
	"net"
	LobbyPacket "packet_lobby"
	ReturnCode "packet_returncode"

	"strconv"

	"github.com/golang/protobuf/proto"
)

// ServerSession is...
type ServerSession struct {
	Network.Session
}

func handleLogin(sessionManager *Network.SessionManager, session *Network.Session, buffer []byte, size int) {
	//이렇게 명시 안하고, 해더보고 파악하려면?
	req := &LobbyPacket.LoginReq{}
	err := proto.Unmarshal(buffer[:size], req)
	processError(err)
	log.Printf(req.String())

	name := req.GetName()
	log.Printf("Logged : %s", name)

	//var res *LobbyPacket.LoginRes
	//res = new(LobbyPacket.LoginRes) new할 필요없이 &구조체{} 형식으로 선언하면 된다
	res := &LobbyPacket.LoginRes{}

	if sessionManager.FindUser(name) {
		res.RetCode = ReturnCode.ReturnCode_retExist
		res.Uid = sessionManager.GetUID(name)
	} else {
		sessionManager.AddUser(name)
		res.RetCode = ReturnCode.ReturnCode_retOK
		res.Uid = sessionManager.GetUID(name)
	}
	log.Printf(res.String())

	packet, err := proto.Marshal(res)
	processError(err)
	log.Printf("Packet size : %d", len(packet))

	fmt.Println(res.String())
	wroteSize, err := session.Send(packet)
	processError(err)

	fmt.Println("Send : " + strconv.Itoa(wroteSize))
	log.Printf(res.String())
}

//-----------------------------------------------------------------------------
func handleSession(sessionManager *Network.SessionManager, session *Network.Session) {

	buffer := make([]byte, 4096)
	defer session.Close()

	for {
		//readnSize, err := session.conn.Read(buffer)
		readSize, err := session.Recv(buffer)
		if err != nil {
			log.Fatalln(err)
			return //네트웍 연결 끊어진 경우
		}

		if readSize > 0 {
			handleLogin(sessionManager, session, buffer, readSize)
			//processError(err)
		}
	}
}

func processError(err error) {
	if err != nil {
		panic(err)
		//log.Fatalln(err)
	}
}

func recoverError() {
	err := recover()
	fmt.Println(err)
}

//-----------------------------------------------------------------------------
func main() {
	defer recoverError()

	sessionManager := Network.SessionManager{}
	sessionManager.Init()

	listener, err := net.Listen("tcp", ":6666")
	processError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		processError(err)

		session := Network.CreateSession(conn)

		//defer session.conn.Close()

		go handleSession(&sessionManager, session)
	}
}
