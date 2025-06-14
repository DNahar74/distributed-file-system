package p2p

import (
	"net"
)

// TCPPeer a peer that is connected over TCP
type TCPPeer struct {
	conn     net.Conn // The ref to this peer's connection
	outbound bool
}

// NewTCPPeer returns an instance of a new Peer over TCP
func NewTCPPeer(conn net.Conn, outbound bool) Peer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}
