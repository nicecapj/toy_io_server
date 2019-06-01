package network

import (
	"log"
	"net"
)

//-----------------------------------------------------------------------------
type Session struct {
	Conn net.Conn
}

func CreateSession(Conn net.Conn) *Session {
	newSession := new(Session)

	if newSession != nil {

		log.Printf("New session : %s", Conn.RemoteAddr())
		//log.Printf("New session : %s", Conn.LocalAddr())
	}

	newSession.Conn = Conn
	return newSession
}

func (session *Session) Send(packet []byte) (int, error) {
	readSize, err := session.Conn.Write([]byte(packet))
	return readSize, err
}

func (session *Session) Recv(packet []byte) (int, error) {
	readSize, err := session.Conn.Read([]byte(packet))
	return readSize, err
}
