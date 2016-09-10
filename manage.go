package main

import (
	"net"
)

type ManageServer struct {
	port      string
	onConnect func()
	onData    func(net.Conn, *Packet)
}

type ManageClient struct {
	port      string
	conn      net.Conn
	onConnect func(net.Conn)
	onData    func(net.Conn, *Packet)
}

func (server *ManageServer) Start() chan bool {
	status := make(chan bool)
	connect, err := net.Listen("tcp", ":"+server.port)
	if err != nil {
		Fatal("MANAGE SERVER", err)
	}
	go func() {
		defer connect.Close()
		status <- true
		for {
			conn, err := connect.Accept()
			if err != nil {
				panic(err)
			}
			if server.onConnect != nil {
				server.onConnect()
			}
			go func() {
				for {
					packet := RecvPacket(conn)
					if packet == nil {
						continue
					}
					server.onData(conn, packet)
				}
			}()
		}
	}()
	return status
}

func (client *ManageClient) Start() chan bool {
	status := make(chan bool)
	conn, err := net.Dial("tcp", ":"+client.port)
	if err != nil {
		Fatal("MANAGE CLIENT", err)
	}
	go func() {
		defer conn.Close()
		status <- true
		client.conn = conn
		client.onConnect(conn)
		for {
			packet := RecvPacket(conn)
			if packet == nil {
				continue
			}
			client.onData(conn, packet)
		}
	}()
	return status
}
