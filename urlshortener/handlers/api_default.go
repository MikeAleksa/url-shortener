package handlers

import (
	"net/http"

	"github.com/MikeAleksa/url-shortener/models"
	"github.com/labstack/echo/v4"
)

// CreateShortUrl - Create a new shortened URL
func (c *Container) CreateShortUrl(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetAnalytics - Get analytics
func (c *Container) GetAnalytics(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetLongUrl - Get a long URL
func (c *Container) GetLongUrl(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
