package main

import (
	"fmt"
	"log"

	"github.com/DNahar74/distributed-file-system/p2p"
)

func main() {
	fmt.Println("Hello World")
	tcpOpts := p2p.NewTCPTransportOptions(":3000")
	tr := p2p.NewTCPTransport(*tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	} else {
		select{}
	}
}
