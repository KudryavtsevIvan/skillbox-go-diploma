package main

import (
	"graduatework/internal/app"
	"graduatework/internal/config"
)

func main() {
	config.GlobalConfig = config.NewConfig("config/config.yaml")

	app.RunServer()

}