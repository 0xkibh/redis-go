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

	buf := make([]byte, 128)

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("error reading from client: ",err.Error())
		os.Exit(1)
	}

	parsedCommand := parseCommand(string(buf))
	fmt.Println(parsedCommand[1:])

	switch parsedCommand[0] {
	case "ping":
		conn.Write([]byte(encodeSimpleString("PONG")))
	case "echo":
		if (len(parsedCommand)>1){
			conn.Write([]byte(encodeBulkString(parsedCommand[1:])))
		}else{
			conn.Write([]byte(encodeSimpleString("NO MESSAGE WITH ECHO!")))
		}
	default:
		conn.Write([]byte(encodeSimpleString("WRONG COMMAND!")))
	}
	
}