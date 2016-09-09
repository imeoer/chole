package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Rule struct {
	Out string `yaml:"out"`
	In  string `yaml:"in"`
}

type Config struct {
	Rules map[string]Rule `yaml:"rules"`
}

func (config *Config) Parse() {
	content, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
	}
}
