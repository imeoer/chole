package main

import (
	"log"
	"net"
)

type Client struct {
}

func connect(port string) net.Conn {
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func connectPool(port string) chan net.Conn {
	connPool := make(chan net.Conn, 10)
	go func() {
		for {
			connPool <- connect(port)
		}
	}()
	return connPool
}

func (client Client) Init() {
	proxyChan := connectPool(PROXY_SERVER_PORT)
	appChan := connectPool(APP_SERVER_PORT)
  for {
		proxy := Proxy{}
		proxy.Start(<- proxyChan, <- appChan, func(data []byte) bool {
			parseDomain(data)
			return true
		})
  }
}
