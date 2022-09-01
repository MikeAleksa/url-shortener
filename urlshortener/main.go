package main

import (
	"time"

	"github.com/MikeAleksa/url-shortener/database"
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

	// sleep for a while to allow the database to be created
	time.Sleep(time.Second * 5)

	// Connect to database
	database.ConnectDB()
	defer database.CloseDB()
	e.Logger.Debug("Successfully connected to database!")

	// CreateShortUrl - Create a new shortened URL
	e.POST("/api/v1/create", c.CreateShortUrl)

	// GetAnalytics - Get analytics
	e.GET("/api/v1/analytics", c.GetAnalytics)

	// GetLongUrl - Get a long URL
	e.GET("/api/v1/retrieve", c.GetLongUrl)

	// RedirectToUrl - redirect a short URL to the long URL
	e.GET("/:url", c.RedirectToUrl)

	// Start server
	e.Logger.Debug("Starting listening on port 5000!")
	e.Logger.Fatal(e.Start(":5000"))
}
