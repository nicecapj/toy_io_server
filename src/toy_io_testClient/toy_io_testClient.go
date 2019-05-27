package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func RecoverError() {
	err := recover()
	log.Fatalln(err)
}

func ProcessError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	defer RecoverError()
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	ProcessError(err)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("To Send Message :")
		text, err := reader.ReadString('\n')
		ProcessError(err)
		conn.Write([]byte(text))
		message, err := bufio.NewReader(conn).ReadString('\n')
		ProcessError(err)
		fmt.Println("Message Received :", string(message))
	}
}
