package main

import (
	"log"

	"github.com/astroparam/ceres/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":4000")
	log.Fatal(srv.ListenAndServe())
}
