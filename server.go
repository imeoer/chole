package main

import (
	"log"
	"net"
)

type Server struct {
}

func (server Server) start(isFrom bool, port string) chan net.Conn {
	client, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	connPool := make(chan net.Conn)
	go func() {
    defer client.Close()
		for {
			conn, err := client.Accept()
			if err != nil {
				log.Fatal(err)
			}
			connPool <- conn
		}
	}()
	return connPool
}

func (server Server) Init() {
	fromChan := server.start(true, CLIENT_SERVER_PORT)
	toChan := server.start(false, PROXY_SERVER_PORT)
	go func() {
		for {
			proxy := Proxy{}
			proxy.Start(<- fromChan, <- toChan, func(data []byte) bool {
				parseDomain(data)
				return true
			})
		}
	}()
}
