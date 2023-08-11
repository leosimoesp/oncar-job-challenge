package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leosimoesp/oncar-job-challenge/api/route"
	"github.com/leosimoesp/oncar-job-challenge/config"
	"github.com/leosimoesp/oncar-job-challenge/db"
)

//go:embed web/public/*
var public embed.FS

func main() {
	cfg := config.Load()
	config.ConnectDB(cfg)

	dbPool := config.GetDBPool()
	defer dbPool.Close()

	args := []string{}
	_ = db.RunMigrate("up", cfg.GetDBString(), "db/migrations", args...)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes api
	route.SetupApiRoutes(cfg, dbPool, e)

	//Routes webApp
	setupWebRoutes(e)

	// Start server
	adressAppServer := fmt.Sprint(":", cfg.AppServerPort)
	e.Logger.Fatal(e.Start(adressAppServer))
}

func setupWebRoutes(e *echo.Echo) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "web/public",
		Index:      "index.html",
		HTML5:      true,
		Filesystem: http.FS(public),
	}))
}
