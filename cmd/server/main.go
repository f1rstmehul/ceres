package main

import (
	"log"

	"github.com/astroparam/eros/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":4000")
	log.Fatal(srv.ListenAndServe())
}
