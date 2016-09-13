package main

import "os"

const (
	PROXY_SERVER_PORT   = "7521"
	MANAGER_SERVER_PORT = "7520"
)

var counter *Counter

func main() {
	counter = &Counter{}
	args := os.Args

	if len(args) > 1 && args[1] == "-s" {
		server := Server{}
		server.Start()
	} else {
		config := new(Config)
		config.Parse()

		status := make(chan bool, len(config.Rules)-1)

		for name, rule := range config.Rules {
			client := Client{
				server: config.Server,
				name:   name,
				in:     rule.In,
				out:    rule.Out,
			}
			status <- (<-client.Start())
		}
	}
}
