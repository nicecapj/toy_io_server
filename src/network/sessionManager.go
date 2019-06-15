package network

import (
	"log"
	"net"
	PROTOCOL "packet_protocol"
	"sync"

	"github.com/golang/protobuf/proto"
)

//SessionManager ...
type SessionManager struct {
	sync.Mutex
	//sessionList map[net.Conn]interface{}
	sessionList map[net.Conn]*Session
}

var sessionManagerInstace *SessionManager
var onceSessionManager sync.Once

func GetSessionManager() *SessionManager {
	onceSessionManager.Do(func() {
		sessionManagerInstace = &SessionManager{}
		sessionManagerInstace.Init()
	})

	return sessionManagerInstace
}

// Init ...
func (this *SessionManager) Init() {

	this.sessionList = make(map[net.Conn]*Session)
}

// CreateSession make new this
func (this *SessionManager) CreateSession(Conn net.Conn) *Session {
	//newSession := new(Session)
	newSession := &Session{}

	if newSession != nil {

		log.Printf("New session : %s", Conn.RemoteAddr())
		//log.Printf("New session : %s", Conn.LocalAddr())

		newSession.InitConnection(Conn)
		this.AddSession(newSession)
	}

	return newSession
}

// AddUser ...
func (this *SessionManager) AddSession(session *Session) bool {

	this.Lock()
	_, ok := this.sessionList[session.Conn]
	if ok == false {
		this.sessionList[session.Conn] = session

		userCount := len(this.sessionList)
		log.Printf("UserCount : %d", userCount)

		this.Unlock()
		return true
	}

	this.sessionList[session.Conn] = session
	this.Unlock()
	return true
}

func (this *SessionManager) RemoveSession(session *Session) bool {
	if this.FindSession(session.Conn) == false {
		return false
	}

	log.Printf("user: %s removed from sessionManager\n", session.Name)

	this.Lock()
	delete(this.sessionList, session.Conn)
	this.Unlock()

	return true
}

// FindUser ...
func (this *SessionManager) FindSession(session net.Conn) bool {

	this.Lock()
	_, ok := this.sessionList[session]
	this.Unlock()

	return ok
}

func (this *SessionManager) Broadcast(protocolID PROTOCOL.ProtocolID, pb proto.Message) {
	for _, session := range this.sessionList {
		if session != nil {
			session.SendPacket(protocolID, pb)
		}
	}
}
