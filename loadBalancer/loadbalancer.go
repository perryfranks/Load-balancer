package main

import (
	"fmt"
	"net/http"
)

type LoadBalancer interface {
	LBNext() Server
	Balance(rw http.Request, req *http.Request)
	NewLoadBalancer() *LoadBalancer
}

type RoundRobinBalancer struct {
	robinIndex int
	port       string
	servers    []Server
}

func (lb *RoundRobinBalancer) LBNext() Server {

	server := lb.servers[lb.robinIndex]
	lb.robinIndex++
	if lb.robinIndex >= len(lb.servers) {
		lb.robinIndex = 0
	}

	return server
}

func (lb *RoundRobinBalancer) Balance(rw http.ResponseWriter, req *http.Request) {
	server := lb.LBNext()
	fmt.Printf("Sending request to port :%d", server.Address())

	server.Serve(rw, req)
}

func NewLoadBalancer(port string, servers []Server) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		robinIndex: 0,
		port:       port,
		servers:    servers,
	}
}
