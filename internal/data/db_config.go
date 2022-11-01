package data

import (
	"fmt"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	fileName := fmt.Sprintf("./config/config.json")
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
