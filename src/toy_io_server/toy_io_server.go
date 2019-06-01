package main

import (
	Network "Network"
	"fmt"
	"log"
	"net"
	LobbyPacket "packet_lobby"
	"strconv"

	"github.com/golang/protobuf/proto"
)

//-----------------------------------------------------------------------------
type ServerSession struct {
	Network.Session
}

func HandleLogin(sessionManager *Network.SessionManager, session *Network.Session, buffer []byte, size int) {
	//이렇게 명시 안하고, 해더보고 파악하려면?
	req := LobbyPacket.LoginReq{}
	err := proto.Unmarshal(buffer[:size], &req)
	processError(err)
	fmt.Printf(req.String())

	name := req.GetName()
	log.Printf("Logged : %s", name)

	res := LobbyPacket.LoginRes{}

	if sessionManager.FindUser(name) {
		res.RetCode = 0
	} else {
		sessionManager.AddUser(name)
		res.RetCode = 1
	}

	packet, err := proto.Marshal(&res)
	processError(err)

	fmt.Println(res.String())
	wroteSize, err := session.Conn.Write(packet)
	processError(err)

	fmt.Println("Send : " + strconv.Itoa(wroteSize))
}

//-----------------------------------------------------------------------------
func handleSession(sessionManager *Network.SessionManager, session *Network.Session) {
	buffer := make([]byte, 4096)

	defer session.Conn.Close()

	for {
		//readnSize, err := session.conn.Read(buffer)
		readSize, err := session.Recv(buffer)
		if readSize > 0 {
			HandleLogin(sessionManager, session, buffer, readSize)
		}
		processError(err)
	}
}

func processError(err error) {
	if err != nil {
		panic(err)
		//log.Fatalln(err)
	}
}

func RecoverError() {
	err := recover()
	fmt.Println(err)
}

//-----------------------------------------------------------------------------
func main() {
	defer recover()

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
