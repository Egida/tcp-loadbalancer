package client

import (
	"fmt"
	"net"

	"github.com/amupxm/tcp-loadbalancer/pkg/logger"
)

type (
	Client interface {
		Connect()
		SendMessage(msg []byte) ([]byte, error)
	}
	client struct {
		log  logger.Logger
		host string
		port int
		conn *net.TCPConn
	}
	// GetMessage()
)

func NewClient(log logger.Logger, host string, port int) Client {
	return &client{
		log:  log,
		host: host,
		port: port,
	}
}

func (c *client) Connect() {
	c.log.Info("listening on", c.host, c.port)

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", c.host, c.port))
	if err != nil {
		c.log.Error(err)
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		c.log.Errorf("can not connect to server ", err)
		return
	}
	c.conn = conn
}

func (c *client) SendMessage(msg []byte) ([]byte, error) {
	if c.conn == nil {
		c.log.Warn("connection is nil")
		c.Connect()
	}
	c.log.Warn("Writing")

	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		c.log.Errorf("failed to send TCP msg ", err)
		return nil, err
	}
	c.log.Info("sending", msg, "to", c.host)
	reply := make([]byte, 1024)

	_, err = c.conn.Read(reply)
	if err != nil {
		c.log.Errorf("failed to get reply TCP msg ", err)
		return nil, err
	}
	return reply, nil
}
