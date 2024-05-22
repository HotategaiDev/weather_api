package services

import "github.com/urimeba/codeChallenge-jha/models"

type IOpenWeatherService interface {
	GetCurrentWeather(lat string, long string) (*models.Response, error)
}
