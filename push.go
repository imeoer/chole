package main

import (
	"github.com/InkProject/ink.go"
	"golang.org/x/net/websocket"
)

type Push struct {
	conn *websocket.Conn
}

func (push *Push) Start() {
	web := ink.New()
	web.Get("/push", func(ctx *ink.Context) {
		websocket.Handler(func(ws *websocket.Conn) {
			push.conn = ws
			for {
				buff := make([]byte, 0xffff)
				_, err := ws.Read(buff)
				if err != nil {
					break
				}
			}
		}).ServeHTTP(ctx.Res, ctx.Req)
		ctx.Stop()
	})
	web.Get("*", ink.Static("./web/dist"))

	Log("", "web console running: http://localhost:"+CONSOLE_SERVER_PORT+"/")
	web.Listen(":" + CONSOLE_SERVER_PORT)
}

func (push *Push) Send(data []byte) {
	if push.conn != nil {
		push.conn.Write(data)
	}
}
