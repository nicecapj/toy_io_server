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

var accountManagerInstace *AccountManager
var onceAccountManager sync.Once

// GetAccountManager is singleton
func GetAccountManager() *AccountManager {
	onceAccountManager.Do(func() {
		accountManagerInstace = &AccountManager{}
		accountManagerInstace.Init()
	})

	return accountManagerInstace
}

// Init ...
func (accountManager *AccountManager) Init() {

	accountManager.userList = make(map[string]bool)
	//accountManager.userList = []string{}
	accountManager.uidList = make(map[string]int64)
	accountManager.beginUIDIndex = 2222
}

// AddUser ...
func (accountManager *AccountManager) AddUser(name string) bool {

	accountManager.Lock()
	_, ok := accountManager.userList[name]
	if ok == false {
		accountManager.userList[name] = true

		userCount := len(accountManager.userList)
		log.Printf("UserCount : %d", userCount)

		accountManager.Unlock()
		return true
	}

	accountManager.userList[name] = true
	accountManager.Unlock()
	return true
}

// RemoveUser is ...
func (accountManager *AccountManager) RemoveUser(name string) bool {
	if accountManager.FindUser(name) == false {
		return false
	}

	delete(accountManager.userList, name)
	log.Printf("user: %s removed from accountManager\n", name)

	return true
}

// FindUser ...
func (accountManager *AccountManager) FindUser(name string) bool {

	accountManager.Lock()
	_, ok := accountManager.userList[name]
	accountManager.Unlock()

	return ok
}

//GetUID ...
func (accountManager *AccountManager) GetUID(name string) int64 {

	accountManager.Lock()
	if val, ok := accountManager.uidList[name]; ok {
		accountManager.Unlock()
		return val
	}

	count := len(accountManager.uidList)
	count = int(accountManager.beginUIDIndex) + count
	accountManager.uidList[name] = int64(count)
	accountManager.Unlock()
	return int64(count)
}
