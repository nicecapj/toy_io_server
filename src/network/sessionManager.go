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

// GetSessionManager is singleton
func GetSessionManager() *SessionManager {
	onceSessionManager.Do(func() {
		sessionManagerInstace = &SessionManager{}
		sessionManagerInstace.Init()
	})

	return sessionManagerInstace
}

// Init ...
func (sessionManager *SessionManager) Init() {

	sessionManager.sessionList = make(map[net.Conn]*Session)
}

// CreateSession make new sessionManager
func (sessionManager *SessionManager) CreateSession(Conn net.Conn) *Session {
	//newSession := new(Session)
	newSession := &Session{}

	if newSession != nil {

		log.Printf("New session : %s", Conn.RemoteAddr())
		//log.Printf("New session : %s", Conn.LocalAddr())

		newSession.InitConnection(Conn)
		sessionManager.AddSession(newSession)
	}

	return newSession
}

// AddSession is ...
func (sessionManager *SessionManager) AddSession(session *Session) bool {

	sessionManager.Lock()
	_, ok := sessionManager.sessionList[session.Conn]
	if ok == false {
		sessionManager.sessionList[session.Conn] = session

		userCount := len(sessionManager.sessionList)
		log.Printf("UserCount : %d", userCount)

		sessionManager.Unlock()
		return true
	}

	sessionManager.sessionList[session.Conn] = session
	sessionManager.Unlock()
	return true
}

// RemoveSession is ...
func (sessionManager *SessionManager) RemoveSession(session *Session) bool {
	if sessionManager.FindSession(session.Conn) == false {
		return false
	}

	log.Printf("user: %s removed from sessionManager\n", session.Name)

	sessionManager.Lock()
	delete(sessionManager.sessionList, session.Conn)
	sessionManager.Unlock()

	return true
}

// FindSession ...
func (sessionManager *SessionManager) FindSession(session net.Conn) bool {

	sessionManager.Lock()
	_, ok := sessionManager.sessionList[session]
	sessionManager.Unlock()

	return ok
}

// Broadcast is ...
func (sessionManager *SessionManager) Broadcast(protocolID PROTOCOL.ProtocolID, pb proto.Message) {
	for _, session := range sessionManager.sessionList {
		if session != nil {
			session.SendPacket(protocolID, pb)
		}
	}
}
