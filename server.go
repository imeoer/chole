package main

import (
	"log"
	"net"
)

type Server struct {
}

func (server Server) listen(isFrom bool, port string) chan net.Conn {
	client, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	connPool := make(chan net.Conn)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("SERVER", err)
			}
			defer client.Close()
		}()
		for {
			conn, err := client.Accept()
			if err != nil {
				panic(err)
			}
			connPool <- conn
		}
	}()
	return connPool
}

func (server Server) Start() {
	fromChan := server.listen(true, CLIENT_SERVER_PORT)
	toChan := server.listen(false, PROXY_SERVER_PORT)
	go func() {
		for {
			proxy := Proxy{
				from: <-fromChan,
				to:   <-toChan,
				valid: func(data []byte) bool {
					// domain := ParseDomain(data)
					return true
				},
			}
			<-proxy.Start()
		}
	}()
}
