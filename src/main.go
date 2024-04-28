package main

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("Creating Redis Server.")
	var l,err = net.Listen("tcp",":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close() // close connection once finished

	// infinite loop to read and write response
	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}

		conn.Write([]byte("+PONG\r\n"))
	}
}