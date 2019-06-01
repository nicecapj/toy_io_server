package network

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
