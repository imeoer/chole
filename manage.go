package main

import (
	"log"
	"net"
)

type ManageServer struct {
	port      string
	onConnect func()
	onData    func(net.Conn, []byte)
}

type ManageClient struct {
	port      string
	conn      net.Conn
	onConnect func(net.Conn)
	onData    func([]byte)
}

func (server *ManageServer) Start() chan bool {
	status := make(chan bool)
	connect, err := net.Listen("tcp", ":"+server.port)
	if err != nil {
		log.Fatal("MANAGE SERVER", err)
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
					buff := RecvPacket(conn)
					if buff == nil {
						break
					}
					server.onData(conn, buff)
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
		log.Fatal("MANAGE CLIENT", err)
	}
	go func() {
		defer conn.Close()
		status <- true
		client.conn = conn
		client.onConnect(conn)
		for {
			buff := RecvPacket(conn)
			if err != nil || buff == nil {
				break
			}
			client.onData(buff)
		}
	}()
	return status
}
