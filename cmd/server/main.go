package main

import (
	"net"

	config "github.com/amupxm/tcp-loadbalancer/cfg"
	"github.com/amupxm/tcp-loadbalancer/internal/server"
	"github.com/amupxm/tcp-loadbalancer/pkg/logger"
)

type TCPhandler func(n *net.TCPConn, message string) bool

func main() {
	log := logger.NewLogger("server")
	srv := server.NewServer(config.GetFlags().Host, config.GetFlags().Port, log, TCPHandler)
	srv.Listen()
}

func TCPHandler(n *net.TCPConn, message string) bool {
	log := logger.NewLogger("server/tcpHandler")
	log.Info("message received", message)
	// TODO : implement handler for accepting more messages
	return true
}
