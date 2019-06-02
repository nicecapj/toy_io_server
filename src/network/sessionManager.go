package network

//SessionManager ...
type SessionManager struct {
	userList      []string
	uidList       map[string]int64
	beginUIDIndex int64
}

// Init ...
func (sessionManager *SessionManager) Init() {

	sessionManager.userList = make([]string, 64)
	sessionManager.uidList = make(map[string]int64)
	sessionManager.beginUIDIndex = 2222
}

// AddUser ...
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

// FindUser ...
func (sessionManager *SessionManager) FindUser(name string) bool {

	for _, item := range sessionManager.userList {
		if item == name {
			return true
		}
	}

	return false
}

//GetUID ...
func (sessionManager *SessionManager) GetUID(name string) int64 {

	if val, ok := sessionManager.uidList[name]; ok {
		return val
	}

	count := len(sessionManager.uidList)
	count = int(sessionManager.beginUIDIndex) + count
	sessionManager.uidList[name] = int64(count)
	return int64(count)
}
