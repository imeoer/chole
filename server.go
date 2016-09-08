package main

import (
	"log"
	"net"
)

type Server struct {
	clients map[string]net.Conn
}

func (server *Server) listen(isFrom bool, port string) chan net.Conn {
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
			if isFrom {
				server.newConnect("chole.io")
			}
			connPool <- conn
		}
	}()
	return connPool
}

func (server *Server) newConnect(domain string) {
	if conn, ok := server.clients[domain]; ok {
		SendPacket(conn, []byte("new"))
	}
}

func (server *Server) startManage(port string) {
	manage := ManageServer{
		port: port,
		onData: func(conn net.Conn, buff []byte) {
			domain := string(buff)
			if domain != "" {
				server.clients[domain] = conn
			}
		},
	}
	<-manage.Start()
}

func (server *Server) Start() chan bool {
	status := make(chan bool)
	server.clients = make(map[string]net.Conn)
	go func() {
		server.startManage(MANAGER_SERVER_PORT)
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
						// log.Println(domain)
						return true
					},
					closed: func(isFrom bool) {
					},
				}
				<-proxy.Start(true)
			}
		}()
		status <- true
	}()
	return status
}
