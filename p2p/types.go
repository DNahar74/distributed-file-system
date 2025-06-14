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
