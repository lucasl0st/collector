package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"strings"
)

func main() {
	c, err := ParseConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to load config: %v", err))
	}

	apiKeys := strings.Split(c.ApiKeys, ",")

	storage, err := NewFileStorage(c.StoragePath)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create storage: %v", err))
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	a := NewApi(storage, apiKeys)
	a.RegisterEndpoints(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", c.Port)))
}
