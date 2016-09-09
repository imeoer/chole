package main

import (
	"log"
	"net"
)

type Connection struct {
	manage net.Conn
	from   chan net.Conn
	to     chan net.Conn
}

type Server struct {
	clients map[string]Connection
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
				connection := server.clients[port]
				connection.from <- conn
			} else {
				toPort := RecvPacket(conn)
				if toPort != nil {
					connection := server.clients[string(toPort)]
					connection.to <- conn
				}
			}
		}
	}()
}

func (server *Server) newConnect(port string) {
	if connection, ok := server.clients[port]; ok {
		SendPacket(connection.manage, []byte(port))
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

func (server *Server) waitProxy(connection Connection) {
	for {
		fromConn := <-connection.from
		toConn := <-connection.to
		server.newProxy(fromConn, toConn)
	}
}

func (server *Server) newManage(port string) {
	manage := ManageServer{
		port: port,
		onData: func(conn net.Conn, buff []byte) {
			port := string(buff)
			if port != "" {
				connection := Connection{
					manage: conn,
					from:   make(chan net.Conn),
					to:     make(chan net.Conn),
				}
				server.clients[port] = connection
				go server.waitProxy(connection)
				server.listen(true, port)
			}
		},
	}
	<-manage.Start()
}

func (server *Server) Start() chan bool {
	status := make(chan bool)
	server.clients = make(map[string]Connection)
	go func() {
		server.newManage(MANAGER_SERVER_PORT)
		server.listen(false, PROXY_SERVER_PORT)
		status <- true
	}()
	return status
}
