package tcpserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func DefaultHandler(c *net.TCPConn) error {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	activeConnections = activeConnections + 1
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return err
		}

		message := strings.TrimSpace(string(netData))
		if ok := handleMessage(c, message); !ok {
			break
		}
	}
	c.Close()
	activeConnections = activeConnections - 1
	return nil
}

func handleMessage(c *net.TCPConn, msg string) bool {
	var operation = msg
	msgArr := strings.Split(msg, "::")
	if len(msgArr) > 1 {
		operation = msgArr[0]
	}

	switch strings.ToUpper(operation) {
	case "STOP":
		return false
	case "ACTIVE_CONNECTIONS":
		ResponseString(c, fmt.Sprintf("%d", activeConnections))
		return true
	}
	return true
}
