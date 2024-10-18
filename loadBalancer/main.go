package main

import (
	"flag"
	"fmt"
	"net/http"
)

var lb LoadBalancer

func handleRequest(w http.ResponseWriter, r *http.Request) {

}

func getServers(configLocation string) []Server {
	// read the file and get the addresses
	// massage them into a server structure
	// servers may need to change

	return []Server{}

}

func main() {
	port := flag.String("port", "4000", "Port for incomming requests\n")
	configLocation := flag.String("config", "./config.yml", "location of config file with server endpoints")
	flag.Parse()

	fmt.Printf("Loadbalancer listening on port:%v\n", *port)
	//

	servers := getServers(*configLocation)

	lb = NewLoadBalancer(*port, servers)

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":"+*port, nil)

}
