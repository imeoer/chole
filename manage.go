package main

import (
	"net"
)

type ManageServer struct {
	port      string
	onConnect func()
	onEvent   func(net.Conn, string, string)
	onClose   func(net.Conn)
}

type ManageClient struct {
	conn       net.Conn
	server     string
	remoteAddr string
	onConnect  func(net.Conn)
	onEvent    func(net.Conn, string, string)
	onClose    func(net.Conn)
}

func (server *ManageServer) Start() chan bool {
	status := make(chan bool)
	connect, err := net.Listen("tcp", ":"+server.port)
	if err != nil {
		Fatal("MANAGE", err)
	}
	go func() {
		defer connect.Close()
		status <- true
		for {
			conn, err := connect.Accept()
			if err != nil {
				Error("MANAGE", err)
				continue
			}
			if server.onConnect != nil {
				go server.onConnect()
			}
			go func() {
				defer func() {
					if server.onClose != nil {
						server.onClose(conn)
					}
					conn.Close()
				}()
				for {
					packet := RecvPacket(conn)
					if packet == nil {
						Error("MANAGE", "client disconnected")
						break
					}
					go server.onEvent(conn, packet.event, packet.data)
				}
			}()
		}
	}()
	return status
}

func (client *ManageClient) Close() {
	TryClose(client.conn)
}

func (client *ManageClient) Start() chan bool {
	status := make(chan bool)
	conn, err := net.Dial("tcp", client.server)
	if err != nil {
		Fatal("MANAGE", err)
	}
	client.conn = conn
	go func() {
		defer func() {
			if client.onClose != nil {
				client.onClose(conn)
			}
			conn.Close()
		}()
		status <- true
		go client.onConnect(conn)
		for {
			packet := RecvPacket(conn)
			if packet == nil {
				Error("MANAGE", "server disconnected")
				break
			}
			go client.onEvent(conn, packet.event, packet.data)
		}
	}()
	return status
}
