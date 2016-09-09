package main

const PROXY_SERVER_PORT = "7521"
const MANAGER_SERVER_PORT = "7520"

var counter *Counter

func main() {
	counter = &Counter{}

	server := Server{}
	<-server.Start()

	config := new(Config)
	config.Parse()

	status := make(chan bool, len(config.Rules)-1)

	for name, rule := range config.Rules {
		client := Client{
			name: name,
			in:   rule.In,
			out:  rule.Out,
		}
		status <- (<-client.Start())
	}
}
