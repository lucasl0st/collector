package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func authenticator(apiKeys []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, apiKey := range apiKeys {
				if apiKey == c.Request().Header.Get("Authorization") {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header")
		}
	}
}
