package models

import "github.com/labstack/echo/v4"

type Route struct {
	Name           string
	Method         string
	Pattern        string
	HandlerFunc    echo.HandlerFunc
	MiddlewareFunc echo.MiddlewareFunc
}

type Routes []Route
