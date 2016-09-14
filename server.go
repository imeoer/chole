package main

import (
	"net"
	"strings"
	"time"
)

type Connection struct {
	manage   net.Conn
	listener net.Listener
	from     chan net.Conn
	to       chan net.Conn
}

type Server struct {
	clients map[string]Connection
}

func (server *Server) listen(isFrom bool, port string, block bool) net.Listener {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		Fatal("SERVER", err)
	}
	handler := func() {
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				if isFrom {
					break
				} else {
					Error("SERVER", err)
					continue
				}
			}
			go func() {
				if isFrom {
					if ok := server.newConnect(conn); !ok {
						conn.Close()
					}
				} else {
					packet := RecvPacket(conn)
					if packet != nil && packet.event == "REQUEST_PROXY" {
						if ok := server.tryProxy(conn, packet.data); !ok {
							conn.Close()
						}
					}
				}
			}()
		}
	}
	if block {
		handler()
	} else {
		go handler()
	}
	return listener
}

func (server *Server) newConnect(conn net.Conn) bool {
	_, reqPort, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		return false
	}
	for addr, connection := range server.clients {
		if strings.HasSuffix(addr, reqPort) {
			SendPacket(connection.manage, "REQUEST_COMING", reqPort)
			connection.from <- conn
			return true
		}
	}
	return false
}

func (server *Server) tryProxy(conn net.Conn, addr string) bool {
	if connection, ok := server.clients[addr]; ok {
		connection.to <- conn
		return true
	}
	return false
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
			Error("SERVER", "reset connection")
			TryClose(fromConn)
			continue
		}
	}
}

func (server *Server) newManage(port string) {
	manage := ManageServer{
		port: port,
		onEvent: func(conn net.Conn, event string, data string) {
			if event == "REQUEST_PORT" {
				reqPort := data
				for addr := range server.clients {
					if strings.HasSuffix(addr, reqPort) {
						SendPacket(conn, "REQUEST_PORT_REJECT", reqPort)
						return
					}
				}
				connection := Connection{
					manage:   conn,
					from:     make(chan net.Conn),
					to:       make(chan net.Conn),
					listener: server.listen(true, reqPort, false),
				}
				remoteAddr := conn.RemoteAddr().String()
				server.clients[remoteAddr+":"+reqPort] = connection
				go server.waitProxy(connection)
			}
		},
		onClose: func(conn net.Conn) {
			remoteAddr := conn.RemoteAddr().String()
			for addr, connection := range server.clients {
				if strings.Index(addr, remoteAddr) == 0 {
					connection.listener.Close()
					delete(server.clients, addr)
					Log("CLOSE", remoteAddr)
				}
			}
		},
	}
	<-manage.Start()
}

func (server *Server) Start() {
	server.clients = make(map[string]Connection)
	server.newManage(MANAGER_SERVER_PORT)
	server.listen(false, PROXY_SERVER_PORT, true)
}
