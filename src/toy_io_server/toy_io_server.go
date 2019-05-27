package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func handleSession(conn net.Conn) {
	buffer := make([]byte, 4096)

	for {
		readnSize, err := conn.Read(buffer)
		if ProcessError(err) {
			return
		}

		wroteSize, err := conn.Write(buffer[:readnSize])
		if ProcessError(err) {
			return
		}

		fmt.Println("Send : " + strconv.Itoa(wroteSize))
	}
}

func ProcessError(err error) bool {
	if err != nil {
		//panic(err)
		log.Fatalln(err)
		return true
	}
	return false
}

func RecoverError() {
	err := recover()
	fmt.Println(err)
}

func main() {
	defer recover()

	listener, err := net.Listen("tcp", ":6666")
	ProcessError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if ProcessError(err) {
			continue
		}
		defer conn.Close()

		go handleSession(conn)
	}
}
