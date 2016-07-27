package main

import (
// "log"
)

const CLIENT_SERVER_PORT = "8002"
const PROXY_SERVER_PORT = "8001"
const APP_SERVER_PORT = "8000"

func main() {
	counter = &Counter{}

	server := Server{}
	server.Start()

	client := Client{}
	client.Start()
}
