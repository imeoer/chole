package main

import (
	"log"
	"net"
)

type Client struct {
	name string
	in   string
	out  string
}

func (client *Client) connect(addr string) net.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("CLIENT: ", err)
		return nil
	}
	return conn
}

func (client *Client) newConnect() {
	fromConn := client.connect(":" + PROXY_SERVER_PORT)
	toConn := client.connect("127.0.0.1:" + client.in)
	proxy := Proxy{
		from: fromConn,
		to:   toConn,
		init: func(fromConn net.Conn) {
			SendPacket(fromConn, []byte(client.out))
		},
		valid: func(data []byte) bool {
			// domain := ParseDomain(data)
			// log.Println(domain)
			return true
		},
	}
	<-proxy.Start(false)
}

func (client *Client) Start() chan bool {
	manage := ManageClient{
		port: MANAGER_SERVER_PORT,
		onConnect: func(conn net.Conn) {
			if client.out != "" {
				SendPacket(conn, []byte(client.out))
			}
		},
		onData: func(data []byte) {
			event := string(data)
			if event != "" {
				client.newConnect()
			}
		},
	}
	return manage.Start()
}
