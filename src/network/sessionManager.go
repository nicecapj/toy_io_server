package network

import (
	"log"
	"sync"
)

//SessionManager ...
type SessionManager struct {
	sync.Mutex
	userList      []string
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

	sessionManager.userList = make([]string, 64)
	//sessionManager.userList = []string{}
	sessionManager.uidList = make(map[string]int64)
	sessionManager.beginUIDIndex = 2222
}

// AddUser ...
func (sessionManager *SessionManager) AddUser(name string) bool {

	sessionManager.Lock()
	for id, item := range sessionManager.userList {
		if item == "" {
			sessionManager.userList[id] = name
			sessionManager.Unlock()
			return true
		}
	}

	sessionManager.userList = append(sessionManager.userList, name)

	userCount := len(sessionManager.userList)

	log.Printf("UserCount : %d", userCount)
	sessionManager.userList[userCount-1] = name
	sessionManager.Unlock()

	return true
}

// FindUser ...
func (sessionManager *SessionManager) FindUser(name string) bool {

	sessionManager.Lock()
	for _, item := range sessionManager.userList {
		if item == name {
			sessionManager.Unlock()
			return true
		}
	}
	sessionManager.Unlock()
	return false
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
