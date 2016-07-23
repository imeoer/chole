package main

import (
	"log"
	"io"
	"net"
)

type Proxy struct {
	checked bool
	from   net.Conn
	to   net.Conn
	inspect func(data []byte) bool
}

func (proxy Proxy) pipe(src, dst io.ReadWriter) bool {
	buff := make([]byte, 0xffff)
	for {
		size, err := src.Read(buff)
		if err != nil {
			break
		}
		data := buff[:size]
		if !proxy.checked {
			proxy.checked = true
			if !proxy.inspect(data) {
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

func (proxy Proxy) Start(from net.Conn, to net.Conn, inspect func(data []byte) bool) {
	proxy.inspect = inspect
	go func() {
		proxy.pipe(from, to)
		from.Close()
		to.Close()
		log.Println("done")
	}()
	go proxy.pipe(to, from)
}
