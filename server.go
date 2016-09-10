package main

import (
	"net"
	"time"
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
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		Fatal("Server", err)
	}
	go func() {
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				Error("Server", err)
				continue
			}
			if isFrom {
				server.newConnect(port)
				connection := server.clients[port]
				connection.from <- conn
			} else {
				packet := RecvPacket(conn)
				if packet != nil && packet.event == "REQUEST_PIPE" {
					fromPort := packet.data
					connection := server.clients[fromPort]
					connection.to <- conn
				}
			}
		}
	}()
}

func (server *Server) newConnect(port string) {
	if connection, ok := server.clients[port]; ok {
		SendPacket(connection.manage, "REQUEST_COMING", port)
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
		select {
		case toConn := <-connection.to:
			server.newProxy(fromConn, toConn)
		case <-time.After(time.Second * 5):
			Error("SERVER", "Reset Connection")
			TryClose(fromConn)
			continue
		}
	}
}

func (server *Server) newManage(port string) {
	manage := ManageServer{
		port: port,
		onData: func(conn net.Conn, packet *Packet) {
			if packet != nil && packet.event == "REQUEST_PORT" {
				port := packet.data
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
