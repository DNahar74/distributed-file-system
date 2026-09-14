package p2p

import "github.com/DNahar74/distributed-file-system/message"

// Peer is any remote node
type Peer interface {
	Close() error
}

// Transport is anything that handles communication between the nodes in a server.
// This can be anything like TCP, UDP, WS, etc.
type Transport interface {
	ListenAndAccept() error
	Dial(string) error
	Consume() <-chan message.Message
}

// HandshakeFunc is the type every handshake func needs to follow
type HandshakeFunc func(Peer) error
