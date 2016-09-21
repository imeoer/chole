package main

import (
	"github.com/InkProject/ink.go"
	"golang.org/x/net/websocket"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type Package struct {
	Id    string      `msgpack:"id"`
	Event string      `msgpack:"event"`
	Data  interface{} `msgpack:"data"`
}

type Push struct {
	conn *websocket.Conn
}

func (push *Push) Start() {
	web := ink.New()
	web.Get("/push", func(ctx *ink.Context) {
		websocket.Handler(func(ws *websocket.Conn) {
			push.conn = ws
			Log("PUSH", "new connection")
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

func (push *Push) Send(id string, event string, data interface{}) bool {
	if push.conn != nil {
		pkg, err := msgpack.Marshal(&Package{
			Id:    id,
			Event: event,
			Data:  data,
		})
		if err != nil {
			return false
		}
		err = websocket.Message.Send(push.conn, pkg)
		if err == nil {
			return true
		} else {
			Error("PUSH", err)
		}
	}
	return false
}
