package main

import (
	"fmt"
	"net"
	"os"
)

func main(){
	fmt.Println("Creating Redis Server.")
	var l,err = net.Listen("tcp",":6379")
	if err != nil {
		fmt.Println(err)
		return
	}


	// infinite loop to read and write response
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleRequest(conn) // goroutine
		
	}
}

func handleRequest(conn net.Conn){
	defer conn.Close() // close connection once finished

	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("error reading from client: ",err.Error())
		os.Exit(1)
	}

	conn.Write([]byte("+PONG\r\n"))
}