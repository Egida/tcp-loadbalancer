package loadbalancer

import "net"

type (
	client struct {
		UUID       string
		IP         string
		Connection net.Conn
	}
)

var WhiteList []string
var BlackList []string
var ActiveClients []client
