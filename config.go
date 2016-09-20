package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/go-fsnotify/fsnotify"
	"gopkg.in/yaml.v2"
)

type Rule struct {
	Out string `yaml:"out"`
	In  string `yaml:"in"`
}

type Config struct {
	Server  string           `yaml:"server"`
	Rules   map[string]*Rule `yaml:"rules"`
	Clients map[string]*Client
}

func (rule Rule) getId(name string) string {
	return fmt.Sprintf("%s:%s:%s", name, rule.In, rule.Out)
}

func (config *Config) parse() {
	content, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		Fatal("CONFIG", err)
	}
	if err = yaml.Unmarshal(content, &config); err != nil {
		Fatal("CONFIG", err)
	}
	if config.Server == "" {
		config.Server = "localhost"
	}
	for _, rule := range config.Rules {
		if rule.In == "" {
			Fatal("CONFIG", "must specify 'in' port in rule")
		}
		if rule.Out == "" {
			rule.Out = rule.In
		}
		if _, err := strconv.Atoi(rule.In); err == nil {
			rule.In = ":" + rule.In
		}
	}
}

func (config *Config) apply() {
	config.parse()

	idMap := make(map[string]bool)
	for name, rule := range config.Rules {
		id := rule.getId(name)
		idMap[id] = true
	}

	for id, client := range config.Clients {
		if _, ok := idMap[id]; !ok {
			client.Close()
			config.RemoveClient(id)
		}
	}

	for name, rule := range config.Rules {
		id := rule.getId(name)
		if _, ok := config.Clients[id]; !ok {
			client := Client{
				id:     id,
				server: config.Server,
				name:   name,
				in:     rule.In,
				out:    rule.Out,
				onClose: func(id string) {
					config.RemoveClient(id)
				},
				onEvent: func(id string, event string, data string) {
					Log("EVENT", id+","+event+","+data)
				},
			}
			status := <-client.Start()
			if status {
				config.AddClient(&client)
			}
		}
	}
}

func (config *Config) AddClient(client *Client) {
	config.Clients[client.id] = client
}

func (config *Config) RemoveClient(id string) {
	if _, ok := config.Clients[id]; ok {
		delete(config.Clients, id)
	}
}

func (config *Config) Watch() {
	config.Clients = make(map[string]*Client)
	config.apply()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		Error("CONFIG", err)
		return
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					config.apply()
					Log("CONFIG", "applyed new configuration")
				}
			case err := <-watcher.Errors:
				Error("CONFIG", err)
			}
		}
	}()

	err = watcher.Add("./config.yaml")
	if err != nil {
		Fatal("CONFIG", err)
	}

	push := Push{}
	push.Start()
	<-done
}
