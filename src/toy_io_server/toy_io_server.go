package main

import (
	Network "Network"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
)

// ServerSession is...
type ServerSession struct {
	Network.Session
}

func handleLogin(session *Network.Session, buffer []byte, size int) {
	session.HandlePacket(buffer)
}

//-----------------------------------------------------------------------------
func handleSession(session *Network.Session) {

	buffer := make([]byte, 4096)
	defer session.Close()

	for {
		//readnSize, err := session.conn.Read(buffer)
		readSize, err := session.Recv(buffer)
		if err != nil && err != io.EOF {
			//log.Panicln(err)
			fmt.Println(err)
			return
		} else if err == io.EOF {
			log.Printf("Close connection\n")
			return
		} else if err != nil {
			fmt.Println(err)
			return
		}

		if readSize > 0 {
			handleLogin(session, buffer, readSize)
			//processError(err)
		}
	}
}

func processError(err error) {
	if err != nil {
		panic(err)
		//log.Fatalln(err)
	}
}

func recoverError() {
	err := recover()
	fmt.Println(err)
}

//-----------------------------------------------------------------------------
func main() {
	//defer recoverError()

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Printf("GOMAX Proc : %d\n", runtime.NumCPU())

	//sessionManager := Network.GetSessionManager()

	listener, err := net.Listen("tcp", ":6666")
	processError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		processError(err)

		session := Network.CreateSession(conn)

		go handleSession(session)
	}
}
