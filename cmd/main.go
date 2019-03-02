package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/toothless-discovery/server"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 1234, "Discovery server port")
	flag.Parse()
}

func startServer(handler http.Handler, addr string) {
	h := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	if err := h.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %s", err)
	}
}

func main() {
	s := server.New()

	addr := fmt.Sprintf(":%d", port)
	go startServer(s.Handler(), addr)

	log.Printf("Discovery server started on port %d.", port)

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
