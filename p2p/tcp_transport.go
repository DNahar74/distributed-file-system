package p2p

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// TCPTransport represents transport over TCP
type TCPTransport struct {
	listenAddress string       // Address on which this node listens (Eg: port 3000 is exposed)
	listener      net.Listener // Listener for incoming connections

	mu    sync.RWMutex      // RWMutex to access peers safely during concurrent access
	peers map[net.Addr]Peer // Peers to track connected peer nodes identified by their network address
}

// NewTCPTransport returns a TCPTransport instance for the given address
func NewTCPTransport(listenAddr string) Transport {
	return &TCPTransport{
		listenAddress: listenAddr,
		peers:         make(map[net.Addr]Peer),
	}
}

// ListenAndAccept starts a TCP server on the caller's address
func (t *TCPTransport) ListenAndAccept() error {
	listener, err := net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	t.listener = listener

	go t.startAcceptingConnections()
	return nil
}

// Dial tries to connect the caller to a TCP server on the given address
func (t *TCPTransport) Dial(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	p := NewTCPPeer(conn, true)
	t.mu.Lock()
	t.peers[conn.RemoteAddr()] = p
	t.mu.Unlock()

	return nil
}

// Peers returns an array of all connected peers
func (t *TCPTransport) Peers() []Peer {
	t.mu.RLock()
	defer t.mu.RUnlock()

	pArr := make([]Peer, 0, len(t.peers))

	for _, peer := range t.peers {
		fmt.Println("Remote conn address:", peer.(*TCPPeer).conn.RemoteAddr())
		pArr = append(pArr, peer)
	}

	return pArr
}

func (t *TCPTransport) startAcceptingConnections() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Println("Error accepting connection")
		}

		go t.handleConnection(conn)
	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	p := NewTCPPeer(conn, false)
	t.mu.Lock()
	t.peers[conn.RemoteAddr()] = p
	t.mu.Unlock()

	// Send a message via channel to this connection such that it adds us into it's peer network
	log.Printf("New connection: %v", conn.RemoteAddr())
}
