package main

import (
	"net/http"
	"net/http/httputil"
)

type Server interface {
	Address() string
	// this may be the simpliest way to get started
	// just have this as an easy way of makeing a request
	Serve(rw http.ResponseWriter, req *http.Request)
}

type BasicServer struct {
	addr string
	// TODO: make a reverse proxy
	proxy *httputil.ReverseProxy // ah this is the magic
}

func (s *BasicServer) Address() string {
	return s.addr
}

func (s *BasicServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}
