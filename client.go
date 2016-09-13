package main

import (
	"net"
)

type Client struct {
	server string
	name   string
	in     string
	out    string
	proxys []Proxy
	manage ManageClient
}

func (client *Client) connect(addr string) net.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		Error("CLIENT", err)
		return nil
	}
	return conn
}

func (client *Client) newConnect(conn net.Conn) bool {
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
			localAddr := conn.LocalAddr().String()
			SendPacket(fromConn, "REQUEST_PROXY", localAddr+":"+client.out)
		},
		valid: func(data []byte) bool {
			// domain := ParseDomain(data)
			// log.Println(domain)
			return true
		},
	}
	client.proxys = append(client.proxys, proxy)
	<-proxy.Start(false)
	return true
}

func (client *Client) Close() {
	client.manage.Close()
	for _, proxy := range client.proxys {
		proxy.Close()
	}
}

func (client *Client) Start() chan bool {
	manage := ManageClient{
		port: MANAGER_SERVER_PORT,
		onConnect: func(conn net.Conn) {
			SendPacket(conn, "REQUEST_PORT", client.out)
		},
		onEvent: func(conn net.Conn, event string, data string) {
			switch event {
			case "REQUEST_COMING":
				client.newConnect(conn)
				break
			case "REQUEST_PORT_REJECT":
				Error("CLIENT", "requested port "+client.out+" has been used")
				conn.Close()
				break
			}
		},
	}
	client.proxys = make([]Proxy, 0)
	client.manage = manage
	return manage.Start()
}
