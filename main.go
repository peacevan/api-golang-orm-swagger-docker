package main

import (
	"github.com/peacevan/api-golang-orm-swagger-docker/config"
	"github.com/peacevan/api-golang-orm-swagger-docker/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
