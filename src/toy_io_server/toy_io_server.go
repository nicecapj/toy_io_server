package main

import (
	Network "network"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"util"
)

//-----------------------------------------------------------------------------
func handleSession(user *User) {

	recvBuffer := make([]byte, 4096)
	defer user.Close()

	buffer := user.Session.PoolBuffer.Get().(*bytes.Buffer)
	defer func() {
		buffer.Reset()
		user.Session.PoolBuffer.Put(buffer)
	}()

	for {
		//readnSize, err := session.conn.Read(buffer)
		readSize, err := user.Recv(recvBuffer)
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
		} else {
			//이 경우는 뭐가 있을까
		}

		if readSize == 0 {
			continue
		}

		if Network.PacketHeaderLen > readSize {
			continue
		}

		{
			buffer.Write(recvBuffer[:readSize])

			bufferArray := buffer.Bytes()

			header, err := Network.GetHeader(bufferArray[:])
			if err != nil {
				log.Panicln("read header")
			}

			if header.PacketSize > Network.MaxPacketSize {
				panic("packet size larger than maxsize")
			}

			//channel style
			//user.DispatchPacket(header.PacketID, bufferArray[Network.PacketHeaderLen:header.PacketSize])

			//immediately process style
			user.HandlePacket(header.PacketID, bufferArray[Network.PacketHeaderLen:header.PacketSize])

			buffer.Next(int(header.PacketSize))
		}
	}
}

//-----------------------------------------------------------------------------
func main() {
	//defer recoverError()

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Printf("GOMAX Proc : %d\n", runtime.NumCPU())

	sessionManager := Network.GetSessionManager()

	listener, err := net.Listen("tcp", ":6666")
	util.ProcessError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		util.ProcessError(err)

		session := sessionManager.CreateSession(conn)
		user := &User{}
		user.Init(session)

		go handleSession(user)
	}
}
