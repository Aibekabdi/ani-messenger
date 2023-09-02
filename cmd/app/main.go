package main

import (
	"os"

	"github.com/Aibekabdi/ani-messenger/config"
	"github.com/Aibekabdi/ani-messenger/internal/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config loading error: %s", err)
	}

	// Setup logger
	if err := setupLogger(cfg.Log); err != nil {
		log.Fatalf("Error when setting up the logger: %s", err)
	}

	//Run
	app.Run(cfg)
}

func setupLogger(cfg config.Log) error {
	level, err := log.ParseLevel(cfg.Level)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	switch cfg.Format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}

	switch cfg.Output {
	case "stderr":
		log.SetOutput(os.Stderr)
	case "stdout":
		log.SetOutput(os.Stdout)
	default:
		file, err := os.OpenFile(cfg.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Infof("Failed to log to file, using default stderr: %v", err)
		}
	}
	return nil
}

//
