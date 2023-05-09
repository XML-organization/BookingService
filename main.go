package main

import (
	"booking-service/startup"
	cfg "booking-service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
