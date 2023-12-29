package config

import (
	"flag"
	"fmt"
)

type (
	InstanceType int
	Config       struct {
		Host         string
		Port         int
		InstanceType InstanceType
	}
)

const (
	InstanceTypeClient InstanceType = iota
	InstanceTypeServer
	InstanceTypeLoadBalancer
)

var flg = Config{}

func init() {
	var strInstanceType string
	flag.StringVar(&flg.Host, "host", "localhost", "host to listen on")
	flag.IntVar(&flg.Port, "port", 8080, "port to listen on")
	flag.StringVar(&strInstanceType, "instance-type", "client", "instance type [client, server, loadbalancer]")

	flg.InstanceType.Set(strInstanceType)

	flag.Parse()
}

func (i InstanceType) String() string {
	switch i {
	case InstanceTypeClient:
		return "client"
	case InstanceTypeServer:
		return "server"
	case InstanceTypeLoadBalancer:
		return "loadbalancer"
	default:
		return "unknown"
	}
}

func (i *InstanceType) Set(s string) error {
	switch s {
	case "client":
		*i = InstanceTypeClient
	case "server":
		*i = InstanceTypeServer
	case "loadbalancer":
		*i = InstanceTypeLoadBalancer
	default:
		return fmt.Errorf("unknown instance type: %s", s)
	}

	return nil
}

func GetFlags() Config {
	return flg
}

func GetAddress() string {
	return flg.Host + ":" + string(flg.Port)
}
