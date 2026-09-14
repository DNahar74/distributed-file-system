package main

import (
	"fmt"
	"log"

	"github.com/DNahar74/distributed-file-system/config"
	"github.com/DNahar74/distributed-file-system/p2p"
)

func main() {
	fmt.Println("Hello World")
	flags, err := config.LoadFlags()
	if err != nil {
		log.Fatal(err)
	}

	tcpOpts := p2p.NewTCPTransportOptions(flags.Start)
	tr := p2p.NewTCPTransport(*tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	} else {
		for _, v := range flags.Connect {
			err := tr.Dial(v)
			if err != nil {
				fmt.Println("Error Dialing peer:", v)
				fmt.Println(err)
			}
		}
		select{}
	}
}
