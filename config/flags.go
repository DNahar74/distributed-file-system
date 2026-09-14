package config

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

// Flags contains the port to start the connection and the peers addresses
type Flags struct {
	Start    string
	Connect []string
}

// LoadFlags parses and provides the validated flags
func LoadFlags() (*Flags, error) {
	c := &Flags{}

	startPtr := flag.String("start", ":3000", "This stores the host's port address")
	connectPtr := flag.String("connect", ":3001", "This stores comma separated peer URIs")

	flag.Parse()

	// Check if flags are provided

	inputFlags := make(map[string]bool)

	flag.Visit(func(f *flag.Flag) {
		inputFlags[f.Name] = true
	})

	if !inputFlags["start"] {
		return nil, errors.New("Initialization URL not provided")
	}

	// Get Valid conn. addresses for p2p conn.

	startURL, err := getValidURL(*startPtr)
	if err != nil {
		return nil, err
	}
	
	c.Start = startURL

	if inputFlags["connect"] {
		peers := strings.Split(*connectPtr, ",")
		for i, v := range peers {
			peerURL, err := getValidURL(v)
			if err != nil {
				return nil, err
			}
	
			peers[i] = peerURL
		}
	
		c.Connect = peers
	}

	return c, nil
}

func getValidURL(s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty URL string")
	}

	if strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") {
    return "", fmt.Errorf("HTTP/HTTPS URLs are not supported for raw P2P transport connections")
	}

	host, port := "", ""
	var err error

	if strings.Contains(s, "://") {
		u, err := url.Parse(s)
		if err != nil {
			return "", err
		}

		if u.Scheme != "tcp" && u.Scheme != "p2p" {
			return "", fmt.Errorf("unsupported transport layer protocol: %s", u.Scheme)
		}

		if u.Port() == "" {
			return "", errors.New("No specified port")
		}

		host, port, err = net.SplitHostPort(u.Host)
		if err != nil {
			return "", err
		}
	} else if strings.HasPrefix(s, ":") {
		// 2. Handle shorthand local ports (e.g., ":3000")
		host = "127.0.0.1" // Fallback to local loopback interface for P2P testing
		port = s[1:]
	} else {
		// 3. Handle raw host:port combinations (e.g., "localhost:3000", "192.168.1.5:4000")
		host, port, err = net.SplitHostPort(s)
		if err != nil {
			return "", err
		}
	}

	portVal, err := strconv.Atoi(port)
	if err != nil || portVal < 1 || portVal > 65535 {
		return "", fmt.Errorf("invalid network port number: %s", port)
	}

	return net.JoinHostPort(host, port), nil
}