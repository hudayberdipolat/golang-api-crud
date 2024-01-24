package main

import (
	"fmt"
	"github.com/hudayberdipolat/golang-api-crud/internal/app"
	"github.com/hudayberdipolat/golang-api-crud/internal/setup/constructor"
	"log"
)

func main() {
	appDependencies := app.GetAppDependencies()
	log.Println("starting program...")
	constructor.Build(appDependencies)
	newApp := app.NewApp(&appDependencies)
	runServer := fmt.Sprintf("%s:%s", appDependencies.AppConfig.ServerConfig.ServerHost,
		appDependencies.AppConfig.ServerConfig.ServerPort)
	if err := newApp.Listen(runServer); err != nil {
		log.Fatal("server run error", err.Error())
		return
	}
}
