package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type (
	Client interface {
	}
	client struct {
		address  string
		interval time.Duration
	}
)

// This is a test client and will send a message in a loop to the server
func NewClient(address string, interval time.Duration) Client {
	return &client{
		address:  address,
		interval: interval,
	}
}

func (c *client) ListenAndServe() {
	addr, err := net.ResolveTCPAddr("tcp", c.address)
	if err != nil {
		log.Println("address is invalid", c.address)
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("can not connect to server :", c.address)
		return
	}

	for {
		msg := ""
		fmt.Println()
		log.Println("sending", msg, "to", c.address)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println("failed to send TCP msg ", err)
			return
		}
		reply := make([]byte, 1024)

		_, err = conn.Read(reply)
		if err != nil {
			log.Println("failed to get reply TCP msg ", err)
			return
		}

		parseReply(reply)
		time.Sleep(c.interval)
	}

}

func parseReply(reply []byte) {
	log.Println("reply is", string(reply))
}
