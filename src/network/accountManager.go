package network

import (
	"log"
	"sync"
)

//AccountManager ...
type AccountManager struct {
	sync.Mutex
	userList      map[string]bool
	uidList       map[string]int64
	beginUIDIndex int64
}

var AccountManagerInstace *AccountManager
var onceAccountManager sync.Once

func GetAccountManager() *AccountManager {
	onceAccountManager.Do(func() {
		AccountManagerInstace = &AccountManager{}
		AccountManagerInstace.Init()
	})

	return AccountManagerInstace
}

// Init ...
func (this *AccountManager) Init() {

	this.userList = make(map[string]bool)
	//this.userList = []string{}
	this.uidList = make(map[string]int64)
	this.beginUIDIndex = 2222
}

// AddUser ...
func (this *AccountManager) AddUser(name string) bool {

	this.Lock()
	_, ok := this.userList[name]
	if ok == false {
		this.userList[name] = true

		userCount := len(this.userList)
		log.Printf("UserCount : %d", userCount)

		this.Unlock()
		return true
	}

	this.userList[name] = true
	this.Unlock()
	return true
}

func (this *AccountManager) RemoveUser(name string) bool {
	if this.FindUser(name) == false {
		return false
	}

	delete(this.userList, name)
	log.Println("user: %s removed from this", name)

	return true
}

// FindUser ...
func (this *AccountManager) FindUser(name string) bool {

	this.Lock()
	_, ok := this.userList[name]
	this.Unlock()

	return ok
}

//GetUID ...
func (this *AccountManager) GetUID(name string) int64 {

	this.Lock()
	if val, ok := this.uidList[name]; ok {
		this.Unlock()
		return val
	}

	count := len(this.uidList)
	count = int(this.beginUIDIndex) + count
	this.uidList[name] = int64(count)
	this.Unlock()
	return int64(count)
}
