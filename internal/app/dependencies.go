package app

import (
	"github.com/hudayberdipolat/golang-api-crud/pkg/config"
	"github.com/hudayberdipolat/golang-api-crud/pkg/database/dbConfig"
	customHttp "github.com/hudayberdipolat/golang-api-crud/pkg/http"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type AppDependencies struct {
	DB         *gorm.DB
	AppConfig  *config.Config
	HttpClient *http.Client
}

func GetAppDependencies() AppDependencies {
	// get config
	getConfig, err := config.GetAppConfig()
	if err != nil {
		log.Fatal("error config file ", err.Error())
	}
	// db Connection
	dbConnect, err := dbConfig.NewDbConfig(getConfig)
	if err != nil {
		log.Fatal("error database config ", err.Error())
	}

	httpClient := customHttp.NewHttpClient()

	return AppDependencies{
		DB:         dbConnect,
		AppConfig:  getConfig,
		HttpClient: httpClient,
	}
}
