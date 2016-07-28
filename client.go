package main

import (
	"log"
	"net"
)

type Client struct {
}

func connect(addr string) net.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("CLIENT: ", err)
		return nil
	}
	return conn
}

func connectPool(addr string, size int) chan net.Conn {
	connPool := make(chan net.Conn, size)
	go func() {
		for {
			conn := connect(addr)
			if conn != nil {
				connPool <- conn
			}
		}
	}()
	return connPool
}

func (client Client) Start() {
	for {
		<-newConn
		fromConn := connect(":" + PROXY_SERVER_PORT)
		toConn := connect(":" + APP_SERVER_PORT)
		proxy := Proxy{
			from: fromConn,
			to:   toConn,
			valid: func(data []byte) bool {
				// domain := ParseDomain(data)
				return true
			},
		}
		<-proxy.Start(false)
	}
}
