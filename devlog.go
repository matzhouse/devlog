package main

import (
	"log"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", "localhost:1222")
	if err != nil {
		log.Println(err)
		return
	}

	for {

		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		go handleConnection(conn)
	}

}
