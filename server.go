package main

import (
	"log"
	"net"
)

type Connection struct {
	port string
	conn net.Conn
}

type Server struct {
	clients  map[string]net.Conn
	fromChan chan Connection
	toChan   chan Connection
}

func (server *Server) listen(isFrom bool, port string) {
	client, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
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
			if isFrom {
				server.newConnect(port)
				server.fromChan <- Connection{conn: conn, port: port}
			} else {
				toPort := RecvPacket(conn)
				if toPort != nil {
					server.toChan <- Connection{conn: conn, port: string(toPort)}
				}
			}
		}
	}()
}

func (server *Server) newConnect(port string) {
	if conn, ok := server.clients[port]; ok {
		SendPacket(conn, []byte(port))
	}
}

func (server *Server) newProxy(from net.Conn, to net.Conn) {
	proxy := Proxy{
		from: from,
		to:   to,
		valid: func(data []byte) bool {
			// domain := ParseDomain(data)
			// log.Println(domain)
			return true
		},
		closed: func(isFrom bool) {
		},
	}
	<-proxy.Start(true)
}

func (server *Server) newManage(port string) {
	manage := ManageServer{
		port: port,
		onData: func(conn net.Conn, buff []byte) {
			port := string(buff)
			if port != "" {
				server.clients[port] = conn
				server.listen(true, port)
			}
		},
	}
	<-manage.Start()
}

func (server *Server) Start() chan bool {
	status := make(chan bool)
	server.fromChan = make(chan Connection)
	server.toChan = make(chan Connection)
	server.clients = make(map[string]net.Conn)
	go func() {
		server.newManage(MANAGER_SERVER_PORT)
		server.listen(false, PROXY_SERVER_PORT)
		go func() {
			for {
				fromConn := (<-server.fromChan).conn
				toConn := (<-server.toChan).conn
				server.newProxy(fromConn, toConn)
			}
		}()
		status <- true
	}()
	return status
}
