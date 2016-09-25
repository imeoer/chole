package main

import "net"

type Client struct {
	id      string
	server  string
	name    string
	in      string
	out     string
	flow    uint64
	proxys  *SafeMap
	manage  *ManageClient
	onClose func(string)
	onEvent func(string, string, interface{})
}

func (client *Client) connect(addr string) net.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		Error("CLIENT", err)
		return nil
	}
	return conn
}

func (client *Client) newConnect(conn net.Conn) chan bool {
	fromConn := client.connect(client.server + ":" + PROXY_SERVER_PORT)
	toConn := client.connect(client.in)
	if fromConn == nil || toConn == nil {
		TryClose(fromConn)
		TryClose(toConn)
		return nil
	}
	proxy := Proxy{
		id:   UUID(),
		from: fromConn,
		to:   toConn,
		init: func(fromConn net.Conn) {
			remoteAddr := client.manage.remoteAddr
			SendPacket(fromConn, "REQUEST_PROXY", remoteAddr)
		},
		valid: func(data []byte) bool {
			// domain := ParseDomain(data)
			// log.Println(domain)
			return true
		},
		onData: func(id string, data []byte) {
			client.onData(id, data)
		},
		closed: func(id string) {
			client.onProxyClose(id)
		},
	}
	client.addProxy(&proxy)
	return proxy.Start(false)
}

func (client *Client) getConns() int {
	return client.proxys.Len()
}

func (client *Client) getFlow() uint64 {
	return client.flow
}

func (client *Client) addProxy(proxy *Proxy) {
	client.proxys.Set(proxy.id, proxy)
	client.onEvent(client.id, "CONNECTIONS", client.getConns())
}

func (client *Client) removeProxy(id string) {
	proxy := client.proxys.Get(id)
	if proxy != nil {
		client.proxys.Set(id, nil)
		client.onEvent(client.id, "CONNECTIONS", client.getConns())
	}
}

func (client *Client) onProxyClose(id string) {
	client.removeProxy(id)
}

func (client *Client) onData(id string, data []byte) {
	client.flow = client.flow + uint64(len(data))
	client.onEvent(client.id, "DATA", client.flow)
}

func (client *Client) Close() {
	TryClose(client.manage.conn)
	for _, pProxy := range client.proxys.Data() {
		proxy := pProxy.(*Proxy)
		proxy.Close()
		client.removeProxy(proxy.id)
	}
}

func (client *Client) Start() chan bool {
	status := make(chan bool)
	manage := ManageClient{
		server: client.server + ":" + MANAGER_SERVER_PORT,
		onConnect: func(conn net.Conn) {
			SendPacket(conn, "REQUEST_PORT", client.out)
		},
		onEvent: func(conn net.Conn, event string, data string) {
			switch event {
			case "REQUEST_COMING":
				<-client.newConnect(conn)
				break
			case "REQUEST_PORT_ACCEPT":
				client.manage.remoteAddr = data
				status <- true
				break
			case "REQUEST_PORT_REJECT":
				Error("CLIENT", "requested port "+client.out+" has been used")
				conn.Close()
				status <- false
				break
			}
		},
		onClose: func(net.Conn) {
			client.onClose(client.id)
		},
	}
	client.proxys = NewSafeMap()
	client.manage = &manage
	<-manage.Start()
	return status
}
