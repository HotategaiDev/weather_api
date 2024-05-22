package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/urimeba/codeChallenge-jha/models"
)

var (
	units = map[string]string{
		"default":  "Kelvin",
		"metric":   "Celsius",
		"imperial": "Fahrenheit",
	}
)

type OpenWeatherRepository struct {
	Configuration *models.Configuration
	HTTPClient    *http.Client
}

func NewOpenWeatherRepository(
	configuration *models.Configuration,
) *OpenWeatherRepository {
	return &OpenWeatherRepository{
		Configuration: configuration,
		HTTPClient:    configureHttpClient(configuration),
	}
}

func configureHttpClient(configuration *models.Configuration) *http.Client {
	timeout := time.Duration(configuration.OpenWeather.Timeout)

	return &http.Client{
		Timeout: timeout * time.Second,
	}
}

func (repository *OpenWeatherRepository) GetCurrentWeather(lat string, long string, responseChan chan models.WrappedResponse) {
	response := new(models.Response)

	currentWeatherEndpoint := repository.Configuration.OpenWeather.Host + repository.Configuration.OpenWeather.CurrentWeatherEndpoint
	endpoint := fmt.Sprintf(currentWeatherEndpoint, lat, long, repository.Configuration.OpenWeather.ApiKey, repository.Configuration.OpenWeather.Units)
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		responseChan <- models.WrappedResponse{
			Response: nil,
			Error:    err,
		}
		return
	}

	request.Header.Add("Authorization", "Bearer "+repository.Configuration.OpenWeather.ApiKey)
	httpResponse, err := repository.HTTPClient.Do(request)
	if err != nil {
		responseChan <- models.WrappedResponse{
			Response: nil,
			Error:    err,
		}
		return
	}

	if httpResponse.StatusCode != 200 {
		responseChan <- models.WrappedResponse{
			Response: nil,
			Error:    errors.New("error retrieving current weather information with status code " + httpResponse.Status),
		}
		return
	}

	defer httpResponse.Body.Close()

	openWeatherResponse := new(models.OpenWeatherResponse)
	err = json.NewDecoder(httpResponse.Body).Decode(openWeatherResponse)
	if err != nil {
		responseChan <- models.WrappedResponse{
			Response: nil,
			Error:    err,
		}
		return
	}

	response.StatusCode = httpResponse.StatusCode
	response.Condition = openWeatherResponse.Weather[0].Main
	response.TemperatureDescription = validateTemperature(openWeatherResponse.Main.Temp)
	response.Temperature = fmt.Sprintf("%.2f", openWeatherResponse.Main.Temp)
	response.TemperatureUnit = validatUnit(repository.Configuration.OpenWeather.Units)
	response.Detail = openWeatherResponse
	responseChan <- models.WrappedResponse{
		Response: response,
		Error:    nil,
	}
}

func validateTemperature(temperature float64) string {
	if temperature >= 20 {
		return "Hot"
	}

	if temperature < 20 && temperature > 10 {
		return "Moderate"
	}

	if temperature <= 10 {
		return "Cold"
	}

	return "Cold"
}

func validatUnit(unit string) string {
	value, exists := units[unit]
	if exists {
		return value
	}

	return units["default"]
}
