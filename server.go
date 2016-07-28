package main

import (
	"log"
	"net"
)

const MAX_CONNECTION_COUNT = 200

type Server struct {
	fromLimit chan bool
	toLimit   chan bool
}

func (server Server) acquire(isFrom bool) {
	if isFrom {
		server.fromLimit <- true
	} else {
		server.toLimit <- true
	}
}

func (server Server) release(isFrom bool) {
	if isFrom {
		<-server.fromLimit
	} else {
		<-server.toLimit
	}
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
			server.release(isFrom)
			defer client.Close()
		}()
		for {
			server.acquire(isFrom)
			conn, err := client.Accept()
			if err != nil {
				panic(err)
			}
			if isFrom {
				newConn <- true
			}
			connPool <- conn
		}
	}()
	return connPool
}

func (server Server) Start() {
	server.fromLimit = make(chan bool, MAX_CONNECTION_COUNT)
	server.toLimit = make(chan bool, MAX_CONNECTION_COUNT)
	fromChan := server.listen(true, CLIENT_SERVER_PORT)
	toChan := server.listen(false, PROXY_SERVER_PORT)
	go func() {
		for {
			fromConn := <-fromChan
			toConn := <-toChan
			proxy := Proxy{
				from: fromConn,
				to:   toConn,
				valid: func(data []byte) bool {
					// domain := ParseDomain(data)
					return true
				},
				closed: func(isFrom bool) {
					server.release(isFrom)
				},
			}
			<-proxy.Start(true)
		}
	}()
}
