package main

import (
	"log"
	"net"

	"github.com/matzhouse/devlog/message"
)

func handleConnection(conn net.Conn) {

	err := message.Parse(conn)

	if err != nil {
		log.Println(err)
		return
	}

}
