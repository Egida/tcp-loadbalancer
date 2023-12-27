package main

import (
	config "github.com/amupxm/tcp-loadbalancer/cfg"
	tcpserver "github.com/amupxm/tcp-loadbalancer/pkg/tcpServer"
)

func main() {
	tcpSrv := tcpserver.StartNewServer(config.GetAddress(), tcpserver.DefaultHandler)
	tcpSrv.ListenAndServe()
}
