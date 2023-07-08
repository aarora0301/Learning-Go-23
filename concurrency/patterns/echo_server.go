package main

import (
	"io"
	"log"
	"net"
)

const listenAddr1 = "localhost:4000"

func main() {
	l, err := net.Listen("tcp", listenAddr1)
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// blocks on single connection
		// keep waiting to get input from first connection
		io.Copy(c, c)
	}
}
