package p2p

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/DNahar74/distributed-file-system/encoding"
	"github.com/DNahar74/distributed-file-system/message"
)

// TCPTransportOptions gives the transport options
type TCPTransportOptions struct {
	listenAddress string           // Address on which this node listens (Eg: port 3000 is exposed)
	handshake     HandshakeFunc    // Handshake for incoming connections
	decoder       encoding.Decoder // Handles the decoding of incoming messages
}

// NewTCPTransportOptions return transport options for the TCP server
// To be configured into main for options of decoder
func NewTCPTransportOptions(listenAddr string) *TCPTransportOptions {
	return &TCPTransportOptions{
		listenAddress: listenAddr,
		handshake: TCPHandshake,
		decoder: encoding.NOPDecoder{},
	}
}

// TCPTransport represents transport over TCP
type TCPTransport struct {
	options  TCPTransportOptions
	listener net.Listener // Listener for incoming connections

	mu    sync.RWMutex      // RWMutex to access peers safely during concurrent access
	peers map[net.Addr]Peer // Peers to track connected peer nodes identified by their network address
}

// NewTCPTransport returns a TCPTransport instance for the given address
func NewTCPTransport(options TCPTransportOptions) Transport {
	return &TCPTransport{
		options:   options,
		peers:     make(map[net.Addr]Peer),
	}
}

// ListenAndAccept starts a TCP server on the caller's address
func (t *TCPTransport) ListenAndAccept() error {
	listener, err := net.Listen("tcp", t.options.listenAddress)
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
	err := t.options.handshake(p)
	if err != nil {
		conn.Close()
		log.Printf("Error completing handshake with connection %+v", p)
		return
	}

	t.mu.Lock()
	t.peers[conn.RemoteAddr()] = p
	t.mu.Unlock()

	log.Printf("New connection: %+v", p)
	badRequests := 0

	msg := message.Message{}
	// Read loop
	for {
		err = t.options.decoder.Decode(conn, &msg)
		if err != nil {
			log.Println("Error decoding the command:", err)
			badRequests++
			if badRequests >= 5 {
				conn.Close()
				return
			}
			continue
		}
		badRequests = 0
		log.Printf("Message is: %+v", msg)
	}
}
