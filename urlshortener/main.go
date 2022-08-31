package main

import (
	"github.com/MikeAleksa/url-shortener/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CreateShortUrl - Create a new shortened URL
	e.POST("/api/v1/create", c.CreateShortUrl)

	// GetAnalytics - Get analytics
	e.GET("/api/v1/analytics", c.GetAnalytics)

	// GetLongUrl - Get a long URL
	e.GET("/api/v1/retrieve", c.GetLongUrl)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}