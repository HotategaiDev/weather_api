package container

import (
	"github.com/rs/zerolog/log"
	"github.com/urimeba/codeChallenge-jha/config"
	"github.com/urimeba/codeChallenge-jha/controllers"
	"github.com/urimeba/codeChallenge-jha/repositories"
	"github.com/urimeba/codeChallenge-jha/services"
	"go.uber.org/dig"
)

func CreateContainer() *dig.Container {
	container := dig.New()

	err1 := container.Provide(config.NewConfiguration)
	err2 := container.Provide(controllers.NewOpenWeatherController)
	err3 := container.Provide(services.NewOpenWeatherService)
	err4 := container.Provide(repositories.NewOpenWeatherRepository)

	processErrors(err1, err2, err3, err4)
	return container
}

func processErrors(errors ...error) {
	for _, err := range errors {
		if err != nil {
			log.Fatal().Msgf("[%s][%s] %s", "Container", "CreateContainer", err.Error())
		}
	}
}
