package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"

	"go.uber.org/dig"

	"github.com/urimeba/codeChallenge-jha/controllers"
	"github.com/urimeba/codeChallenge-jha/middleware"
	"github.com/urimeba/codeChallenge-jha/models"
	container "github.com/urimeba/codeChallenge-jha/server/container"
)

var (
	server *echo.Echo
	routes models.Routes
)

func RunServer() {
	server = echo.New()
	server.HideBanner = true

	createContainer()
}

func createContainer() {
	container := container.CreateContainer()
	setupRoutes(container)
	err := container.Invoke(func(configuration *models.Configuration) {
		setupMiddleware(configuration)
		startServer(configuration.Server.Port)
	})

	if err != nil {
		server.Logger.Fatal(err)
	}
}

func setupRoutes(container *dig.Container) {
	registerRoutes(container)

	for _, route := range routes {
		if route.MiddlewareFunc != nil {
			server.Add(route.Method, route.Pattern, route.HandlerFunc, route.MiddlewareFunc).Name = route.Name
		} else {
			server.Add(route.Method, route.Pattern, route.HandlerFunc).Name = route.Name
		}
	}
}

func registerRoutes(container *dig.Container) {
	err := container.Invoke(func(controller *controllers.OpenWeatherController) {
		routes = append(routes,
			models.Route{
				Name:           "CheckCurrentWeather",
				Method:         "GET",
				Pattern:        "/",
				HandlerFunc:    controller.GetCurrentWeather,
				MiddlewareFunc: middleware.GetCurrentWeatherValidatorMiddleware(),
			})
	})

	if err != nil {
		server.Logger.Fatal(err.Error())
	}
}

func setupMiddleware(configuration *models.Configuration) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	server.Use(
		echoMiddleware.CORSWithConfig(
			echoMiddleware.CORSConfig{
				AllowOrigins: configuration.Server.Cors,
				AllowMethods: []string{http.MethodGet},
				AllowHeaders: []string{"X-Request-With", "Content-Type", "Authorization"},
			},
		),
	)

	server.Use(echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(context echo.Context, reqLogvalues echoMiddleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", reqLogvalues.URI).
				Str("method", context.Request().Method).
				Int("status", reqLogvalues.Status).
				Msg("REQS")
			return nil
		},
	}))
}

func startServer(port int) {
	err := server.Start(fmt.Sprintf(":%d", port))

	if err != nil {
		server.Logger.Fatal(err)
	}
}
