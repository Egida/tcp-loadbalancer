package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/amupxm/tcp-loadbalancer/pkg/logger"
)

type (
	Server interface {
		// Listen starts the server on the host and port specified
		Listen()
	}
	server struct {
		fn                TCPHandlerFunction
		log               logger.Logger
		activeConnections int
		host              string
		port              int
	}
	// you can response to the client with a string message. Returning false will close the connection
	TCPHandlerFunction func(n *net.TCPConn, message string) bool
)

func NewServer(host string, port int, log logger.Logger, handlerFn TCPHandlerFunction) Server {
	return &server{
		log:               log,
		host:              host,
		port:              port,
		activeConnections: 0,
		fn:                handlerFn,
	}
}

func (s *server) Listen() {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		s.log.Errorf("address is invalid", err)
	}

	ls, err := net.ListenTCP("tcp", addr)
	if err != nil {
		s.log.Errorf("error while listening", err)
	}
	defer ls.Close()

	s.log.Infof("Start listening on [%s:%s]", s.host, s.port)
	for {
		conn, err := ls.AcceptTCP()
		s.log.Infof("New connection from %s", conn.RemoteAddr().String())
		if err != nil {
			s.log.Errorf("error while creating connection", err)
		}
		go s.DefaultHandler(conn)
	}
}

func (s *server) DefaultHandler(c *net.TCPConn) error {
	s.log.Infof("Serving %s\n", c.RemoteAddr().String())
	s.activeConnections = s.activeConnections + 1
	for {
		s.log.Infof("message received %s", 11)

		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return err
		}

		message := strings.TrimSpace(string(netData))
		s.log.Infof("message received %s", message)
		if ok := s.HandleMessage(c, message); !ok {
			break
		}
	}
	c.Close()
	s.activeConnections = s.activeConnections - 1
	return nil
}

func (s *server) HandleMessage(c *net.TCPConn, msg string) bool {
	var operation = msg
	msgArr := strings.Split(msg, "::")
	if len(msgArr) > 1 {
		operation = msgArr[0]
	}

	switch strings.ToUpper(operation) {
	case "STOP":
		return false
	case "ACTIVE_CONNECTIONS":
		ResponseString(c, fmt.Sprintf("%d", s.activeConnections))
		return true
	default:
		return s.fn(c, msg)
	}
}

func ResponseString(c *net.TCPConn, str string) error {
	return ResponseByte(c, []byte(str))
}

func ResponseInt(c *net.TCPConn, i int) error {
	return ResponseString(c, fmt.Sprintf("%d", i))
}

func ResponseByte(c *net.TCPConn, b []byte) error {
	_, err := c.Write(b)
	return err
}
