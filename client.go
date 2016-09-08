package main

import (
	"log"
	"net"
)

type Client struct {
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
	toConn := client.connect("127.0.0.1:" + APP_SERVER_PORT)
	proxy := Proxy{
		from: fromConn,
		to:   toConn,
		valid: func(data []byte) bool {
			// domain := ParseDomain(data)
			// log.Println(domain)
			return true
		},
	}
	<-proxy.Start(false)
}

func (client *Client) Start(port string) {
	var manage ManageClient
	manage = ManageClient{
		port: port,
		onConnect: func(conn net.Conn) {
			SendPacket(conn, []byte("chole.io"))
		},
		onData: func(data []byte) {
			event := string(data)
			if event == "new" {
				client.newConnect()
			}
		},
	}
	manage.Start()
}
