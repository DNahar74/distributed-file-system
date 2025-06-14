package p2p

import "testing"

func TestNewTCPTransportAddress(t *testing.T) {
	listenAddr := ":3000"
	tr := NewTCPTransport(listenAddr).(*TCPTransport)

	if tr.listenAddress != listenAddr {
		t.Fatalf("The listen addresses do not match transport address. \n ListenAddress = %s \n TransportAddress = %s", listenAddr, tr.listenAddress)
	}

	err := tr.ListenAndAccept()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewTCPTransportIncorrectAddr(t *testing.T) {
	listenAddr := "3000"
	tr := NewTCPTransport(listenAddr).(*TCPTransport)

	if tr.listenAddress != listenAddr {
		t.Fatalf("The listen addresses do not match transport address. \n ListenAddress = %s \n TransportAddress = %s", listenAddr, tr.listenAddress)
	}

	err := tr.ListenAndAccept()
	if err == nil {
		t.Fatalf("The listen address should not be accepted")
	}
}
