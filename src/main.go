package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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

	commandArray := strings.Split(string(buf), "\r\n")
	if commandArray[2] == "PING" || commandArray[2] == "ping" {
		conn.Write([]byte("+PONG\r\n"))
	} else if (commandArray[2] == "ECHO" || commandArray[2] == "echo") {
		if (len(commandArray)>4){
		msg := fmt.Sprintf("+%s\r\n", commandArray[4])
		conn.Write([]byte(msg))
		}else{
			conn.Write([]byte("+NO MESSAGE WITH ECHO!\r\n"))
		}
	} else{
		conn.Write([]byte("+WRONG COMMAND!\r\n"))
	}
	

	
}