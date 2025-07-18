package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("cdPlayer")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
		//go handleConnection(conn)
	}
}
