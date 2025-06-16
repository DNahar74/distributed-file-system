package p2p

import (
	"testing"
)

func TestNewTCPTransportAddress(t *testing.T) {
	listenAddr := ":3000"
	tcpOpts := NewTCPTransportOptions(listenAddr)
	tr := NewTCPTransport(*tcpOpts).(*TCPTransport)

	if tr.options.listenAddress != listenAddr {
		t.Fatalf("The listen addresses do not match transport address. \n ListenAddress = %s \n TransportAddress = %s", listenAddr, tr.options.listenAddress)
	}

	err := tr.ListenAndAccept()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewTCPTransportIncorrectAddr(t *testing.T) {
	listenAddr := "3000"
	tcpOpts := NewTCPTransportOptions(listenAddr)
	tr := NewTCPTransport(*tcpOpts).(*TCPTransport)

	if tr.options.listenAddress != listenAddr {
		t.Fatalf("The listen addresses do not match transport address. \n ListenAddress = %s \n TransportAddress = %s", listenAddr, tr.options.listenAddress)
	}

	err := tr.ListenAndAccept()
	if err == nil {
		t.Fatalf("The listen address should not be accepted")
	}
}
