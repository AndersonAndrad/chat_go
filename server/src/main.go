package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting chat go server...")
	ln, err := net.Listen("tcp", ":8000")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Chat go channel started...")
	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	_, err := conn.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Receive, %s", buf)
}

func ConnectServer() {
	conn, err := net.Dial("tcp", "localhost:5050")

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Write([]byte("Hello go server, serve is online"))

	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Close()
}
