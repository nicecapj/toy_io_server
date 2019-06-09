package main

import (
	Network "Network"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"util"
)

//-----------------------------------------------------------------------------
func handleSession(user *User) {

	buffer := make([]byte, 4096)
	defer user.Close()

	for {
		//readnSize, err := session.conn.Read(buffer)
		readSize, err := user.Recv(buffer)
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
			user.HandlePacket(buffer)
		}
	}
}

//-----------------------------------------------------------------------------
func main() {
	//defer recoverError()

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Printf("GOMAX Proc : %d\n", runtime.NumCPU())

	//sessionManager := Network.GetSessionManager()

	listener, err := net.Listen("tcp", ":6666")
	util.ProcessError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		util.ProcessError(err)

		session := Network.CreateSession(conn)
		user := &User{}
		user.Init(session)

		go handleSession(user)
	}
}
