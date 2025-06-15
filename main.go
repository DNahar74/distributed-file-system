package main

import (
	"fmt"
	"log"

	"github.com/DNahar74/distributed-file-system/p2p"
)

func main() {
	fmt.Println("Hello World")
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	} else {
		select{}
	}
}
