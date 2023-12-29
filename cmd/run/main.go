package main

import (
	"net"
	"time"

	"github.com/amupxm/tcp-loadbalancer/internal/client"
	"github.com/amupxm/tcp-loadbalancer/internal/server"
	"github.com/amupxm/tcp-loadbalancer/pkg/logger"
)

type (
	Config struct {
		Servers []servers
		Clients []clients
	}
	servers struct {
		Host string
		Port int
	}
	clients struct {
		Host            string
		Port            int
		MessageInterval time.Duration
	}
)

var runConfig = Config{
	Servers: []servers{
		{
			Host: "localhost",
			Port: 8080,
		},
	},
	Clients: []clients{
		{
			Host:            "localhost",
			Port:            8080,
			MessageInterval: 1 * time.Second,
		},
		{
			Host:            "localhost",
			Port:            8080,
			MessageInterval: 1 * time.Second,
		},
		{
			Host:            "localhost",
			Port:            8080,
			MessageInterval: 1 * time.Second,
		},
	},
}

func main() {
	for _, srv := range runConfig.Servers {
		go func() {
			instance := server.NewServer(srv.Host, srv.Port, logger.NewLogger("run/server"), TCPHandler)
			instance.Listen()
		}()
	}
	time.Sleep(1 * time.Second)
	for _, cl := range runConfig.Clients {
		go func() {
			instance := client.NewClient(logger.NewLogger("run/client"), cl.Host, cl.Port)
			instance.Connect()
			time.Sleep(1 * time.Second)

			_, err := instance.SendMessage([]byte("hello\n"))
			if err != nil {
				logger.NewLogger("run/client").Error(err)
			}
		}()
	}
	select {}
	// TODO : implement graceful shutdown
}

func TCPHandler(n *net.TCPConn, message string) bool {
	// TODO : implement handler for accepting more messages
	log := logger.NewLogger("run/server/tcpHandler")
	log.Infof("message received 2", message)
	server.ResponseString(n, message)

	return false
}
