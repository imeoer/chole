package main

import (
	"io"
	"log"
	"net"
)

type Proxy struct {
	inited  bool
	checked bool
	usedChan chan bool
	from    net.Conn
	to      net.Conn
	init    func(io.ReadWriter)
	valid   func([]byte) bool
}

var connCount int

func (proxy Proxy) pipe(src, dst io.ReadWriter, direct bool) {
	connCount++
	log.Println(connCount)
	buff := make([]byte, 0xffff)
	defer func() {
		proxy.from.Close()
		proxy.to.Close()
		connCount--
		log.Println(connCount)
		proxy.usedChan <- true
	}()
	for {
		if direct && !proxy.inited {
			proxy.inited = true
			if proxy.init != nil {
				proxy.init(src)
				_, err := src.Read(buff)
				if err != nil {
					break
				}
			}
		}
		size, err := src.Read(buff)
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

func (proxy Proxy) Start() (chan bool) {
	proxy.usedChan = make(chan bool)
	go proxy.pipe(proxy.from, proxy.to, true)
	go proxy.pipe(proxy.to, proxy.from, false)
	return proxy.usedChan
}
