package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/rand"
	"net"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	RANDOM_STRING_LEN = 5
)

const (
	CLR_W = ""
	CLR_R = "\x1b[31;1m"
	CLR_G = "\x1b[32;1m"
	CLR_B = "\x1b[34;1m"
	CLR_Y = "\x1b[33;1m"
)

// Print log
func Log(name string, info interface{}) {
	log.Printf("%s\t\t%s\n", name, info)
}

// Print error log
func Error(name string, info interface{}) {
	if runtime.GOOS == "windows" {
		log.Printf("ERROR: %s %s\n", name, info)
	} else {
		log.Printf("%s%s\t\t%s%s", CLR_R, name, info, "\x1b[0m")
	}
}

// Print error log and exit
func Fatal(name string, info interface{}) {
	Error(name, info)
	os.Exit(1)
}

type Counter struct {
	count int64
}

type Packet struct {
	event string
	data  string
}

func SendPacket(conn net.Conn, event string, data string) error {
	data = event + ":" + data
	dataByte := []byte(data)
	err := binary.Write(conn, binary.BigEndian, uint16(len(dataByte)))
	err = binary.Write(conn, binary.BigEndian, dataByte)
	Log("SEND", data)
	return err
}

func RecvPacket(conn net.Conn) *Packet {
	lenData := make([]byte, 2)
	n, err := conn.Read(lenData)
	if err == nil && n == 2 {
		len := binary.BigEndian.Uint16(lenData)
		data := make([]byte, len)
		n, err = conn.Read(data)
		if err == nil && n > 0 {
			Log("RECEIVE", string(data))
			dataAry := strings.Split(string(data), ":")
			packet := Packet{event: dataAry[0], data: dataAry[1]}
			return &packet
		}
	}
	return nil
}

func TryClose(conn net.Conn) {
	if conn != nil {
		conn.Close()
	}
}

func (counter *Counter) Up() {
	atomic.AddInt64(&((*counter).count), 1)
	Log("CONNECTION", strconv.Itoa(int(counter.count)))
}

func (counter *Counter) Down() {
	atomic.AddInt64(&((*counter).count), -1)
	Log("CONNECTION", strconv.Itoa(int(counter.count)))
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
