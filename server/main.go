package main

import (
	"flag"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got a request\n")
	w.WriteHeader(http.StatusOK)
}

func main() {

	port := flag.String("port", "4000", "Port to listen on\n")
	flag.Parse()

	http.HandleFunc("/", handler)

	fmt.Printf("Server listening on port:%v\n", *port)

	http.ListenAndServe(":"+*port, nil)

}
