package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/rand"
	"net"
	"regexp"
	"strings"
	"sync/atomic"
	"time"
)

const (
	RANDOM_STRING_LEN = 5
)

type Counter struct {
	count int64
}

func SendPacket(conn net.Conn, data []byte) error {
	err := binary.Write(conn, binary.BigEndian, uint16(len(data)))
	err = binary.Write(conn, binary.BigEndian, data)
	log.Println("Send", string(data), len(data))
	return err
}

func RecvPacket(conn net.Conn) []byte {
	lenData := make([]byte, 2)
	n, err := conn.Read(lenData)
	if err == nil && n == 2 {
		len := binary.BigEndian.Uint16(lenData)
		data := make([]byte, len)
		n, err = conn.Read(data)
		if err == nil && n > 0 {
			log.Println("Receive", n)
			return data
		}
	}
	return nil
}

func (counter *Counter) Up() {
	atomic.AddInt64(&((*counter).count), 1)
	// log.Printf("Connections: + %d\n", counter.count)
}

func (counter *Counter) Down() {
	atomic.AddInt64(&((*counter).count), -1)
	// log.Printf("Connections: - %d\n", counter.count)
}

func ParseDomain(data []byte) (domain string) {
	pos := bytes.Index(data, []byte("\r\n\r\n"))
	header := strings.ToLower(string(data[:pos]))
	log.Println(header)
	regex := regexp.MustCompile("\r\nhost:(.*)\r\n")
	rets := regex.FindStringSubmatch(header)
	if len(rets) >= 2 {
		return strings.TrimSpace(rets[1])
	}
	return
}

func RandomString() string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, RANDOM_STRING_LEN)
	for i := 0; i < RANDOM_STRING_LEN; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
