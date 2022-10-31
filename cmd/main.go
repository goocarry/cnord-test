package main

import (
	"log"
	"os"

	"github.com/goocarry/cnord-test/internal/config"
	"github.com/goocarry/cnord-test/internal/server"
	"github.com/goocarry/cnord-test/internal/store"
)

func main() {
	log := log.New(os.Stdout, "api ", log.LstdFlags)
	
	log.Print("info-34a630a8: gather config")
	cfg := config.GetConfig()

	log.Print("info-4e7e501c: creating store")
	store := store.New(cfg, log)

	log.Println("info-65f340a6: creating new server")
	server.Run(cfg, log, store)

}