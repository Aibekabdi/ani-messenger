package main

import (
	"log"

	"github.com/Aibekabdi/ani-messenger/config"
)

func main() {
	// Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config loading error: %s", err)
	}
	
	log.Println(cfg)
	//Run
}
