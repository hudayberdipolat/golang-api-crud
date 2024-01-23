package main

import (
	"fmt"
	"github.com/hudayberdipolat/golang-api-crud/internal/app"
	"log"
)

func main() {
	appDependencies := app.GetAppDependencies()
	log.Println("starting program...")
	newApp := app.NewApp(&appDependencies)
	runServer := fmt.Sprintf("%s:%s", appDependencies.AppConfig.ServerConfig.ServerHost,
		appDependencies.AppConfig.ServerConfig.ServerPort)
	if err := newApp.Listen(runServer); err != nil {
		log.Fatal("server run error", err.Error())
		return
	}
}
