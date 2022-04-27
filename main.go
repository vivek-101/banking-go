package main

import (
	"github.com/vivek-101/banking/app"
	"github.com/vivek-101/banking/logger"
)

func main() {
	logger.Info("Starting the Application")
	app.Start()
}
