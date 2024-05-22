package services

import (
	"github.com/urimeba/codeChallenge-jha/models"
	"github.com/urimeba/codeChallenge-jha/repositories"
)

type OpenWeatherService struct {
	Configuration    *models.Configuration
	ClientRepository repositories.IClientRepository
}

func NewOpenWeatherService(
	configuration *models.Configuration,
	clientRepository *repositories.OpenWeatherRepository,
) *OpenWeatherService {
	return &OpenWeatherService{
		Configuration:    configuration,
		ClientRepository: clientRepository,
	}
}

func (service *OpenWeatherService) GetCurrentWeather(lat string, long string) (*models.Response, error) {
	responseChan := make(chan models.WrappedResponse)

	go service.ClientRepository.GetCurrentWeather(lat, long, responseChan)
	result := <-responseChan
	if result.Error != nil {
		return nil, result.Error
	}

	return result.Response, nil
}
