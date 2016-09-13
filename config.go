package main

import (
	"io/ioutil"
	"log"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Rule struct {
	Out string `yaml:"out"`
	In  string `yaml:"in"`
}

type Config struct {
	Server string           `yaml:"server"`
	Rules  map[string]*Rule `yaml:"rules"`
}

func (config *Config) Parse() {
	content, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(content, &config); err != nil {
		log.Fatal(err)
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
