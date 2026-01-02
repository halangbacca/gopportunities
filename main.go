package main

import (
	"github.com/halangbacca/gopportunities/config"
	"github.com/halangbacca/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("init config err: %v", err)
		return
	}

	router.Initialize()
}
