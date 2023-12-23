package main

import (
	"net"

	tcpserver "github.com/amupxm/tcp-loadbalancer/pkg/tcpServer"
)

type (
	Server interface {
		ListenAndServe()
	}
	server struct {
		Address           string
		activeConnections int
		handler           handlerFn
	}
	handlerFn func(n *net.TCPConn) error
)

func main() {
	tcpSrv := tcpserver.StartNewServer("localhost:8080", tcpserver.DefaultHandler)
	tcpSrv.ListenAndServe()
}
