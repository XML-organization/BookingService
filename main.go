package main

import (
	"booking-service/startup"
	cfg "booking-service/startup/config"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)
	config := cfg.NewConfig()
	log.Println("Starting server Booking Service...")
	server := startup.NewServer(config)
	server.Start()
}
