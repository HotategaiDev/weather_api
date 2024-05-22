package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urimeba/codeChallenge-jha/models"
	"gopkg.in/yaml.v2"
)

var (
	environment = "dev"
)

func NewConfiguration() *models.Configuration {
	return readFile()
}

func readFile() *models.Configuration {
	configuration := models.Configuration{}

	basePath, err := os.Getwd()
	if err != nil {
		printError(err)
	}

	path := basePath + "/config/config." + environment + ".yml"
	file, err := os.Open(path)
	if err != nil {
		printError(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		printError(err)
	}

	return &configuration
}

func printError(err error) {
	log.Fatal().Msgf("[%s][%s] %s", "Config", "printError", err.Error())
}
