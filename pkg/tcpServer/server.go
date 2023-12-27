package tcpserver

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/amupxm/tcp-loadbalancer/internal/logger"
)

type Server interface {
	ListenAndServe()
}

type (
	server struct {
		address           string
		activeConnections int
		handler           handlerFn
	}
	handlerFn func(n *net.TCPConn) error
)

var activeConnections = 0

// StartNewServer creates a new simple tcp server.
func StartNewServer(addr string, handlerFunction handlerFn) Server {
	s := server{
		address:           addr,
		activeConnections: 0,
		handler:           handlerFunction,
	}
	return &s
}

func (s *server) ListenAndServe() {
	log := logger.NewLogger()
	log.Infof("Starting server on %s \n", s.address)
	go s.listenAndServe(log)
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
}

func (s *server) listenAndServe(log logger.Logger) {
	addr, err := net.ResolveTCPAddr("tcp", s.address)
	if err != nil {
		log.Errorf("address is invalid", err)
	}

	ls, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Errorf("error while listening", err)
	}
	defer ls.Close()

	log.Infof("Start listening on :", s.address)

	for {
		conn, err := ls.AcceptTCP()
		if err != nil {
			log.Errorf("error while creating connection", err)
		}
		go s.handler(conn)
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
