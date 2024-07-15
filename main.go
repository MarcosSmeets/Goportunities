package main

import (
	"github.com/MarcosSmeets/Goportunities/config"
	"github.com/MarcosSmeets/Goportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err.Error())
		return
	}

	// Initialize Router
	router.Initialize()
}
