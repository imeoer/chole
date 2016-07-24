package main

import (
	"log"
	"net"
)

type Client struct {
}

func connect(port string) net.Conn {
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func connectPool(port string, size int) chan net.Conn {
	connPool := make(chan net.Conn, size)
	go func() {
		for {
			connPool <- connect(port)
		}
	}()
	return connPool
}

func (client Client) Start() {
	fromChan := connectPool(PROXY_SERVER_PORT, 50)
	toChan := connectPool(APP_SERVER_PORT, 50)
	for {
		fromConn := <-fromChan
		toConn := <-toChan
		proxy := Proxy{
			from: fromConn,
			to: toConn,
			valid: func(data []byte) bool {
				// domain := ParseDomain(data)
				return true
			},
		}
		closedChan := proxy.Start()
		<- closedChan
	}
}
