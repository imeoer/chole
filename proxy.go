package main

import (
	"io"
	"log"
	"net"
)

type Proxy struct {
	inited  bool
	checked bool
	from    net.Conn
	to      net.Conn
	init    func(io.ReadWriter)
	valid   func([]byte) bool
}

func (proxy Proxy) pipe(src, dst io.ReadWriter, direct bool) bool {
	buff := make([]byte, 0xffff)
	for {
		if direct && !proxy.inited {
			proxy.inited = true
			if proxy.init != nil {
				proxy.init(src)
				_, err := src.Read(buff)
				if err != nil {
					break
				}
				log.Println(buff)
			}
		}
		size, err := src.Read(buff)
		if err != nil {
			break
		}
		data := buff[:size]
		if direct && !proxy.checked {
			proxy.checked = true
			if proxy.valid != nil && !proxy.valid(data) {
				return false
			}
		}
		size, err = dst.Write(data)
		if err != nil {
			break
		}
	}
	return true
}

func (proxy Proxy) Start() {
	go func() {
		proxy.pipe(proxy.from, proxy.to, true)
		proxy.from.Close()
		proxy.to.Close()
	}()
	go proxy.pipe(proxy.to, proxy.from, false)
}
