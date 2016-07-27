package main

import (
	"bytes"
	"regexp"
	"strings"
	"log"
	"sync/atomic"
)

type Counter struct {
	count int64
}

func (counter *Counter) Up() {
	atomic.AddInt64(&((*counter).count), 1)
	log.Printf("Connections: + %d\n", counter.count)
}

func (counter *Counter) Down() {
	atomic.AddInt64(&((*counter).count), -1)
	log.Printf("Connections: - %d\n", counter.count)
}

var counter *Counter

func ParseDomain(data []byte) (domain string) {
	pos := bytes.Index(data, []byte("\r\n\r\n"))
	header := strings.ToLower(string(data[:pos]))
	regex := regexp.MustCompile("\r\nhost:(.*)\r\n")
	rets := regex.FindStringSubmatch(header)
	if len(rets) >= 2 {
		return strings.TrimSpace(rets[1])
	}
	return
}
