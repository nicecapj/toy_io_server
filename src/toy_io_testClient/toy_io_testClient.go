package main

import (
	Network "Network"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"util"
)

func testPacketEnterAndLeaveRoom(client *ClientSession) {
	timer := time.NewTimer(1000 * time.Millisecond)
	go func() {
		<-timer.C

		if client.roomID != 0 {
			client.RequestRoomLeaveReq()
		} else {
			client.RequestRoomEnterReq()
		}

		go testPacketEnterAndLeaveRoom(client)
	}()
}

func main() {
	defer util.RecoverError()
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	util.ProcessError(err)

	//sessionManager := Network.GetSessionManager() //no need for client
	//session := sessionManager.CreateSession(conn)
	session := &Network.Session{}
	if session != nil {
		session.InitConnection(conn)
	}

	clientSession := &ClientSession{}
	clientSession.Init(session)

	defer clientSession.Close()

	//send
	name := util.GetRandomName()
	clientSession.RequestLoginReq(name)

	//recv
	recvBuffer := make([]byte, 4096)

	buffer := clientSession.PoolBuffer.Get().(*bytes.Buffer)
	defer func() {
		buffer.Reset()
		clientSession.PoolBuffer.Put(buffer)
	}()

	timer := time.NewTimer(30 * 60 * time.Second)
	go func() {
		<-timer.C
		testPacketEnterAndLeaveRoom(clientSession)
	}()

	//recv
	func(buffer *bytes.Buffer) {
		for {
			//readnSize, err := session.conn.Read(buffer)
			readSize, err := session.Recv(recvBuffer)
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

			if readSize == 0 {
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

				clientSession.DispatchPacket(header.PacketID, bufferArray[Network.PacketHeaderLen:header.PacketSize])

				buffer.Next(int(header.PacketSize))
			}
		}
	}(buffer)

	// //recv
	// go func(session *ClientSession) {
	// 	recvBuffer := make([]byte, 4096)

	// 	buffer := session.Session.PoolBuffer.Get().(*bytes.Buffer)
	// 	defer func() {
	// 		buffer.Reset()
	// 		session.Session.PoolBuffer.Put(buffer)
	// 	}()

	// 	for {
	// 		//readnSize, err := session.conn.Read(buffer)
	// 		readSize, err := session.Recv(recvBuffer)
	// 		if err != nil && err != io.EOF {
	// 			//log.Panicln(err)
	// 			fmt.Println(err)
	// 			return
	// 		} else if err == io.EOF {
	// 			log.Printf("Close connection\n")
	// 			return
	// 		} else if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		} else {
	// 			continue
	// 		}

	// 		if readSize == 0 {
	// 			continue
	// 		}

	// 		{
	// 			buffer.Write(recvBuffer[:readSize])

	// 			bufferArray := buffer.Bytes()

	// 			header, err := Network.GetHeader(bufferArray[:])
	// 			if err != nil {
	// 				log.Panicln("read header")
	// 			}

	// 			if header.PacketSize > Network.MaxPacketSize {
	// 				panic("packet size larger than maxsize")
	// 			}

	// 			session.DispatchPacket(header.PacketID, bufferArray[Network.PacketHeaderLen:header.PacketSize])

	// 			buffer.Next(int(header.PacketSize))
	// 		}
	// 	}
	// }(clientSession)

	// //send
	// //fmt.Println("input name")
	// go func(client *ClientSession) {
	// 	name := GetRandomName()
	// 	clientSession.reqestLogin(name)
	// 	//time.Sleep(1 * time.Second)
	// }(clientSession)
}
