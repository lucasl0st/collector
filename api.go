package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"strings"
	"time"
)

type Api struct {
	storage Storage
	apiKeys []string
}

func NewApi(storage Storage, apiKeys []string) *Api {
	return &Api{storage: storage, apiKeys: apiKeys}
}

func (a *Api) RegisterEndpoints(e *echo.Echo) {
	e.GET("/health", a.Health)
	e.POST("/store/:entity", a.Store, middleware.Logger(), authenticator(a.apiKeys))
}

func (a *Api) Health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (a *Api) Store(c echo.Context) error {
	entity := c.Param("entity")
	if len(entity) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "empty entity")
	}

	err := ValidateEntity(entity)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("failed to validate entity: %v", err))
	}

	labels := strings.Split(c.QueryParam("labels"), ",")
	if len(labels) == 1 && labels[0] == "" {
		labels = []string{}
	}

	data, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to read body")
	}

	if len(data) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "empty body")
	}

	id := fmt.Sprintf("%d_%s", time.Now().Unix(), uuid.New().String())
	err = a.storage.Store(entity, id, labels, data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to store data")
	}

	return c.String(http.StatusAccepted, "accepted data")
}
