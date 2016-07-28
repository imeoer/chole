package main

import (
	"io"
	"log"
	"net"
	"time"
)

const KEEP_ALIVE_TIMEOUT = 5 * time.Second

type Proxy struct {
	isServer bool
	inited   bool
	checked  bool
	usedChan chan bool
	from     net.Conn
	to       net.Conn
	init     func(io.ReadWriter)
	valid    func([]byte) bool
	closed   func(bool)
}

func (proxy Proxy) pipe(src, dst io.ReadWriter, direct bool) {
	if direct && !proxy.isServer {
		counter.Up()
	}
	buff := make([]byte, 0xffff)
	defer func() {
		if direct && !proxy.isServer {
			counter.Down()
		}
		proxy.from.Close()
		proxy.to.Close()
		if proxy.closed != nil {
			proxy.closed(direct)
		}
		proxy.usedChan <- true
	}()
	for {
		if direct && !proxy.inited {
			proxy.inited = true
			if proxy.init != nil {
				proxy.init(src)
				_, err := src.Read(buff)
				if err != nil {
					log.Println(err)
					break
				}
			}
		}
		size, err := src.Read(buff)
		// log.Println("Transport: ", size)
		if err != nil {
			break
		}
		data := buff[:size]
		if direct && !proxy.checked {
			proxy.checked = true
			proxy.usedChan <- true
			if proxy.valid != nil && !proxy.valid(data) {
				break
			}
		}
		size, err = dst.Write(data)
		if err != nil {
			break
		}
	}
}

func (proxy Proxy) Start(isServer bool) chan bool {
	proxy.isServer = isServer
	proxy.usedChan = make(chan bool)

	fromConn := proxy.from.(*net.TCPConn)
	toConn := proxy.to.(*net.TCPConn)
	fromConn.SetKeepAlive(true)
	toConn.SetKeepAlive(true)
	fromConn.SetKeepAlivePeriod(KEEP_ALIVE_TIMEOUT)
	toConn.SetKeepAlivePeriod(KEEP_ALIVE_TIMEOUT)

	go proxy.pipe(proxy.from, proxy.to, true)
	go proxy.pipe(proxy.to, proxy.from, false)

	return proxy.usedChan
}
