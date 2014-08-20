package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jijeshmohan/goseed/model"
)

var (
	host = flag.String("h", "", "Host name")
	port = flag.String("p", "8080", "port number")
)

func main() {
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *host, *port)

	if err := model.InitDb(); err != nil {
		fmt.Printf("Error while initializing db %s", err)
		os.Exit(1)
	}

	server := NewHttpServer("./ui")
	server.StartServer(address)
}
