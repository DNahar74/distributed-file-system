package message

import "net"

// Message represents any message between two nodes
type Message struct {
	Sender  net.Addr
	Payload []byte
}
