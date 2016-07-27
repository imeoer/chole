package main

import (
	"log"
	"net"
)

const CONN_LIMIT = 200

type Server struct {
	fromLimit chan bool
	toLimit chan bool
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
		<- server.fromLimit
	} else {
		<- server.toLimit
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
			connPool <- conn
		}
	}()
	return connPool
}

func (server Server) Start() {
	server.fromLimit = make(chan bool, CONN_LIMIT)
	server.toLimit = make(chan bool, CONN_LIMIT)
	fromChan := server.listen(true, CLIENT_SERVER_PORT)
	toChan := server.listen(false, PROXY_SERVER_PORT)
	go func() {
		for {
			fromConn := <- fromChan
			toConn := <- toChan
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
