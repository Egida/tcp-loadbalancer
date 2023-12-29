package loadbalancer

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/amupxm/tcp-loadbalancer/pkg/logger"
)

// LoadBalancer acts as a TCP proxy. You can register couple of server instances and automatically distributes requests

type (
	LoadBalancer interface {
		StartAndListen()
		RegisterServer(host string, port int)
		RemoveServer(host string, port int)
	}
	loadBalancer struct {
		host string
		port int
		log  logger.Logger
	}
)

func NewLoadBalancer(remoteHost string, remotePort int, host string, port int, log logger.Logger) LoadBalancer {
	return &loadBalancer{
		host: host,
		port: port,
		log:  log,
	}
}

// Server[ip:port]activeConnections
var Servers = map[string]int{}

func (l *loadBalancer) StartAndListen() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", l.host, l.port))
	if err != nil {
		l.log.Errorf("Error listening:", err)
		return
	}
	defer listener.Close()

	l.log.Infof("TCP Proxy listening on %s " + fmt.Sprintf("%s:%d", l.host, l.port))

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			l.log.Errorf("Error accepting connection:", err)
			continue
		}

		go handleConnection(clientConn, l.findPreferredServer())
	}
}

func (l *loadBalancer) findPreferredServer() string {
	targetServer := ""
	targetServerActiveConnections := -1
	for server, activeConnections := range Servers {
		if targetServerActiveConnections < activeConnections {
			targetServer = server
			targetServerActiveConnections = activeConnections
		}
	}
	return targetServer
}

func (l *loadBalancer) RegisterServer(host string, port int) {
	Servers[fmt.Sprintf("%s:%d", host, port)] = 0
}

func (l *loadBalancer) RemoveServer(host string, port int) {
	if Servers[fmt.Sprintf("%s:%d", host, port)] != 0 {
		l.log.Errorf("Server is still active. Please wait until all connections are closed", errors.New("server is still active. Please wait until all connections are closed"))
		return
	}

	delete(Servers, fmt.Sprintf("%s:%d", host, port))
}

func handleConnection(clientConn net.Conn, remoteAddr string) {
	Servers[remoteAddr] = Servers[remoteAddr] + 1
	// Connect to the remote server
	remoteConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		fmt.Println("Error connecting to remote server:", err)
		clientConn.Close()
		return
	}
	defer remoteConn.Close()

	// Copy data from the client to the remote server and vice versa
	go func() {
		_, err := io.Copy(remoteConn, clientConn)
		if err != nil {
			if err.Error() != "EOF" {
				Servers[remoteAddr] = Servers[remoteAddr] - 1
				return
			}

		}
	}()

	_, err = io.Copy(clientConn, remoteConn)
	if err != nil {
		fmt.Println("Error copying from remote to client:", err)
	}

	clientConn.Close()
}
