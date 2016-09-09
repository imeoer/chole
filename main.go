package main

const CLIENT_SERVER_PORT = "8003"
const PROXY_SERVER_PORT = "8002"
const MANAGER_SERVER_PORT = "8001"
const APP_SERVER_PORT = "8000"

var counter *Counter

func main() {
	counter = &Counter{}

	server := Server{}
	<-server.Start()

	config := new(Config)
	config.Parse()

	status := make(chan bool, 0)

	for name, rule := range config.Rules {
		client := Client{
			name: name,
			in:   rule.In,
			out:  rule.Out,
		}
		status <- (<-client.Start())
	}
}
