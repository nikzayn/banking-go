package main

import (
	"github.com/nikzayn/banking/app"
	"github.com/nikzayn/banking/logger"
)

func main() {
	//Application start
	logger.Info("Starting the application..")
	app.Start()
}
