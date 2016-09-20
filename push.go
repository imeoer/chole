package main

import (
	"io"

	"github.com/InkProject/ink.go"
	"golang.org/x/net/websocket"
)

type Push struct {
}

func (push *Push) Start() {
	web := ink.New()
	web.Get("/push", func(ctx *ink.Context) {
		websocket.Handler(func(ws *websocket.Conn) {
			io.Copy(ws, ws)
		}).ServeHTTP(ctx.Res, ctx.Req)
	})
	web.Get("*", ink.Static("./web/dist"))

	Log("", "web console running: http://localhost:"+CONSOLE_SERVER_PORT+"/")
	web.Listen(":" + CONSOLE_SERVER_PORT)
}
