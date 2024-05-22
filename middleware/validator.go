package middleware

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/urimeba/codeChallenge-jha/models"
)

var (
	validate = validator.New()
)

func GetCurrentWeatherValidatorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(context echo.Context) error {
			request := new(models.Request)
			var validationErrors []models.ValidationError

			request.Lat = context.QueryParam("lat")
			request.Long = context.QueryParam("long")

			// Validate Request
			err := validate.Struct(request)
			if err != nil {
				validationError := new(models.ValidationError)
				for _, e := range err.(validator.ValidationErrors) {
					validationError.Key = e.Field()
					validationError.Error = e.Error()
					validationErrors = append(validationErrors, *validationError)
				}
			}

			if len(validationErrors) != 0 {
				return echo.NewHTTPError(http.StatusBadRequest, validationErrors)
			}

			context.Set("queryParams", request)
			return next(context)
		}

	}
}
