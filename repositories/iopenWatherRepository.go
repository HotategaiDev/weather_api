package repositories

import "github.com/urimeba/codeChallenge-jha/models"

type IClientRepository interface {
	GetCurrentWeather(lat string, long string, responseChan chan models.WrappedResponse)
}
