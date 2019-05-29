package main

import (
	"fmt"
	"log"
	"net"
	. "packet_lobby"
	"strconv"

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

func (session *Session) Send(packet []byte) (int, error) {
	readSize, err := session.conn.Write([]byte(packet))
	return readSize, err
}

func (session *Session) Recv(packet []byte) (int, error) {
	readSize, err := session.conn.Read([]byte(packet))
	return readSize, err
}

//-----------------------------------------------------------------------------
type SessionManager struct {
	userList []string
}

func (sessionManager *SessionManager) Init() {

	sessionManager.userList = make([]string, 64)
}

func (sessionManager *SessionManager) AddUser(name string) bool {

	for id, item := range sessionManager.userList {
		if item == "" {
			sessionManager.userList[id] = name
			return true
		}
	}

	sessionManager.userList = append(sessionManager.userList, name)
	sessionManager.userList[len(sessionManager.userList)] = name

	return true
}

func (sessionManager *SessionManager) FindUser(name string) bool {

	for _, item := range sessionManager.userList {
		if item == name {
			return true
		}
	}

	return false
}

func HandleLogin(sessionManager *SessionManager, session *Session, buffer []byte, size int) {
	//이렇게 명시 안하고, 해더보고 파악하려면?
	req := LoginReq{}
	err := proto.Unmarshal(buffer[:size], &req)
	ProcessError(err)
	fmt.Printf(req.String())

	name := req.GetName()
	log.Printf("Logged : %s", name)

	res := LoginRes{}

	if sessionManager.FindUser(name) {
		res.RetCode = 0
	} else {
		sessionManager.AddUser(name)
		res.RetCode = 1
	}

	packet, err := proto.Marshal(&res)
	ProcessError(err)

	fmt.Println(res.String())
	wroteSize, err := session.conn.Write(packet)
	ProcessError(err)

	fmt.Println("Send : " + strconv.Itoa(wroteSize))
}

//-----------------------------------------------------------------------------
func handleSession(sessionManager *SessionManager, session *Session) {
	buffer := make([]byte, 4096)

	defer session.conn.Close()

	for {
		//readnSize, err := session.conn.Read(buffer)
		readSize, err := session.Recv(buffer)
		if readSize > 0 {
			HandleLogin(sessionManager, session, buffer, readSize)
		}
		ProcessError(err)
	}
}

func ProcessError(err error) {
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

	sessionManager := SessionManager{}
	sessionManager.Init()

	listener, err := net.Listen("tcp", ":6666")
	ProcessError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		ProcessError(err)

		session := CreateSession(conn)

		//defer session.conn.Close()

		go handleSession(&sessionManager, session)
	}
}
