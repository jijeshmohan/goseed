package main

import (
	"flag"
	"fmt"
)

var (
	host = flag.String("h", "", "Host name")
	port = flag.String("p", "8080", "port number")
)

func main() {
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *host, *port)

	server := NewHttpServer("./ui")
	server.StartServer(address)
}
