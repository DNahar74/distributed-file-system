package p2p

// Peer is any remote node
type Peer interface {
}

// Transport is anything that handles communication between the nodes in a server.
// This can be anything like TCP, UDP, WS, etc.
type Transport interface {
	ListenAndAccept() error
	Dial(string) error
}

// HandshakeFunc is the type every handshake func needs to follow
type HandshakeFunc func(Peer) error
