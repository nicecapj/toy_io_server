package main

import (
	Network "Network"
	"fmt"
	"log"
	"net"
)

// ServerSession is...
type ServerSession struct {
	Network.Session
}

func handleLogin(sessionManager *Network.SessionManager, session *Network.Session, buffer []byte, size int) {
	//이렇게 명시 안하고, 해더보고 파악하려면?
	session.HandlePacket(buffer)

	//var res *LobbyPacket.LoginRes
	//res = new(LobbyPacket.LoginRes) new할 필요없이 &구조체{} 형식으로 선언하면 된다
	//res := &LobbyPacket.LoginRes{}

	// if sessionManager.FindUser(name) {
	// 	res.RetCode = ReturnCode.ReturnCode_retExist
	// 	res.Uid = sessionManager.GetUID(name)
	// } else {
	// 	sessionManager.AddUser(name)
	// 	res.RetCode = ReturnCode.ReturnCode_retOK
	// 	res.Uid = sessionManager.GetUID(name)
	// }

	//log.Printf(res.String())
	//session.SendPacket(PROTOCOL.ProtocolID_LoginReS, res)
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
