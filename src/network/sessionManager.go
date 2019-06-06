package network

import (
	"log"
	"sync"
)

//SessionManager ...
type SessionManager struct {
	sync.Mutex
	userList      map[string]bool
	uidList       map[string]int64
	beginUIDIndex int64
}

var instace *SessionManager
var once sync.Once

func GetSessionManager() *SessionManager {
	once.Do(func() {
		instace = &SessionManager{}
		instace.Init()
	})

	return instace
}

// Init ...
func (sessionManager *SessionManager) Init() {

	sessionManager.userList = make(map[string]bool)
	//sessionManager.userList = []string{}
	sessionManager.uidList = make(map[string]int64)
	sessionManager.beginUIDIndex = 2222
}

// AddUser ...
func (sessionManager *SessionManager) AddUser(name string) bool {

	sessionManager.Lock()
	_, ok := sessionManager.userList[name]
	if ok == false {
		sessionManager.userList[name] = true

		userCount := len(sessionManager.userList)
		log.Printf("UserCount : %d", userCount)

		sessionManager.Unlock()
		return true
	}

	sessionManager.userList[name] = true
	sessionManager.Unlock()
	return true
}

func (sessionManager *SessionManager) RemoveUser(name string) bool {
	if sessionManager.FindUser(name) == false {
		return false
	}

	delete(sessionManager.userList, name)
	log.Println("user: %s removed from sessionManager", name)

	return true
}

// FindUser ...
func (sessionManager *SessionManager) FindUser(name string) bool {

	sessionManager.Lock()
	_, ok := sessionManager.userList[name]
	sessionManager.Unlock()

	return ok
}

//GetUID ...
func (sessionManager *SessionManager) GetUID(name string) int64 {

	sessionManager.Lock()
	if val, ok := sessionManager.uidList[name]; ok {
		sessionManager.Unlock()
		return val
	}

	count := len(sessionManager.uidList)
	count = int(sessionManager.beginUIDIndex) + count
	sessionManager.uidList[name] = int64(count)
	sessionManager.Unlock()
	return int64(count)
}
