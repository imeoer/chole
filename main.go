package main

import ()

const CLIENT_SERVER_PORT = "8003"
const PROXY_SERVER_PORT = "8002"
const MANAGER_SERVER_PORT = "8001"
const APP_SERVER_PORT = "8000"

var counter *Counter

func main() {
	counter = &Counter{}

	server := Server{}
	<-server.Start()

	client := Client{}
	client.Start(MANAGER_SERVER_PORT)
}
