package main

import (
	"log"

	"github.com/zhainar/gopoker/internal/server"
)

func main() {
	config, err := server.NewConfig("configs/.env")
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.New(config)
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
