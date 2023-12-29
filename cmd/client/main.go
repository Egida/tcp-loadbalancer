package main

import (
	"time"

	config "github.com/amupxm/tcp-loadbalancer/cfg"
	"github.com/amupxm/tcp-loadbalancer/internal/client"
	"github.com/amupxm/tcp-loadbalancer/pkg/logger"
)

// Client will just send a mssg in a loop every 1 second. Its just to flood the server via dummy messages.
// Client is build only for test purposes.
func main() {
	log := logger.NewLogger("client")
	cl := client.NewClient(log, config.GetFlags().Host, config.GetFlags().Port)
	cl.Connect()

	for {
		cl.SendMessage([]byte("hello"))
		time.Sleep(1 * time.Second)
	}
}
