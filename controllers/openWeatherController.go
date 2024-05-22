package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/urimeba/codeChallenge-jha/models"
	"github.com/urimeba/codeChallenge-jha/services"
)

type OpenWeatherController struct {
	Service services.IOpenWeatherService
}

func NewOpenWeatherController(service *services.OpenWeatherService) *OpenWeatherController {
	return &OpenWeatherController{
		Service: service,
	}
}

func (controller *OpenWeatherController) GetCurrentWeather(context echo.Context) error {
	queryParams := context.Get("queryParams").(*models.Request)
	response, err := controller.Service.GetCurrentWeather(queryParams.Lat, queryParams.Long)
	if err != nil {
		errorResponse := new(models.ResponseError)
		errorResponse.DetailError = err.Error()
		return context.JSON(http.StatusInternalServerError, errorResponse)
	}
	return context.JSON(response.StatusCode, response)
}
