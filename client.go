package main

import (
	"net"
)

type Client struct {
	server string
	name   string
	in     string
	out    string
}

func (client *Client) connect(addr string) net.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		Error("Client", err)
		return nil
	}
	return conn
}

func (client *Client) newConnect() bool {
	fromConn := client.connect(client.server + ":" + PROXY_SERVER_PORT)
	toConn := client.connect(client.in)
	if fromConn == nil || toConn == nil {
		TryClose(fromConn)
		TryClose(toConn)
		return false
	}
	proxy := Proxy{
		from: fromConn,
		to:   toConn,
		init: func(fromConn net.Conn) {
			SendPacket(fromConn, "REQUEST_PIPE", client.out)
		},
		valid: func(data []byte) bool {
			// domain := ParseDomain(data)
			// log.Println(domain)
			return true
		},
	}
	<-proxy.Start(false)
	return true
}

func (client *Client) Start() chan bool {
	manage := ManageClient{
		port: MANAGER_SERVER_PORT,
		onConnect: func(conn net.Conn) {
			if client.out != "" {
				SendPacket(conn, "REQUEST_PORT", client.out)
			}
		},
		onData: func(conn net.Conn, packet *Packet) {
			if packet.event == "REQUEST_COMING" {
				client.newConnect()
			}
		},
	}
	return manage.Start()
}
