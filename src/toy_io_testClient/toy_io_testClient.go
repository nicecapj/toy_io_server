package main

import (
	Network "Network"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"util"
)

func main() {
	defer util.RecoverError()
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	util.ProcessError(err)

	session := Network.CreateSession(conn)
	clientSession := &ClientSession{}
	clientSession.Init(session)

	defer clientSession.Close()

	for {
		//send
		name := util.GetRandomName()
		clientSession.reqestLogin(name)

		//recv
		recvBuffer := make([]byte, 4096)

		buffer := clientSession.PoolBuffer.Get().(*bytes.Buffer)
		defer func() {
			buffer.Reset()
			clientSession.PoolBuffer.Put(buffer)
		}()

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
