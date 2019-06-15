package network

import (
	"log"
	"net"
	"sync"
)

//SessionManager ...
type SessionManager struct {
	sync.Mutex
	sessionList map[net.Conn]interface{}
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

	this.sessionList = make(map[net.Conn]interface{})
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

	log.Println("user: %s removed from this", session.Name)
	delete(this.sessionList, session.Conn)

	return true
}

// FindUser ...
func (this *SessionManager) FindSession(session net.Conn) bool {

	this.Lock()
	_, ok := this.sessionList[session]
	this.Unlock()

	return ok
}
