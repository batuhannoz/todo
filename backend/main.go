package main

import (
	TodoApp "github.com/batuhannoz/todo/backend/app"
	"github.com/batuhannoz/todo/backend/config"
)

func main() {
	appConfig := config.GetConfig()
	app := TodoApp.App{}
	err := app.Initialize(appConfig)
	if err != nil {
		panic(err)
	}
	if err := app.Run(appConfig.Server.Port); err != nil {
		panic(err)
	}
}
