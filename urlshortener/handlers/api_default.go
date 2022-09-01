package handlers

import (
	"net/http"

	"github.com/MikeAleksa/url-shortener/database"
	"github.com/MikeAleksa/url-shortener/models"
	"github.com/labstack/echo/v4"
)

// CreateShortUrl - Create a new shortened URL
func (c *Container) CreateShortUrl(ctx echo.Context) error {

	// unmarshall the request body into a CreateShortUrlRequest struct
	var request models.CreateShortUrlRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, "Something went wrong")
	}

	shortUrl := ShortenURL(request.Url)

	// create a new short URL record
	UrlRecord := models.UrlData{
		ShortUrl:  shortUrl,
		LongUrl:   request.Url,
		Redirects: 0,
	}

	// insert the new short URL record into the database
	err := database.CreateUrl(UrlRecord)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, "Error creating short URL")
	}

	return ctx.JSON(
		http.StatusOK,
		models.ShortUrlData{
			Url: shortUrl,
		})
}

// GetAnalytics - Get analytics
func (c *Container) GetAnalytics(ctx echo.Context) error {
	shortUrl := ctx.QueryParam("url")

	urlData, err := database.GetUrl(shortUrl)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusNotFound, "URL not found")
	}

	return ctx.JSON(http.StatusOK, models.AnalyticsData{
		Redirects: urlData.Redirects,
	})
}

// GetLongUrl - Get a long URL
func (c *Container) GetLongUrl(ctx echo.Context) error {
	shortUrl := ctx.QueryParam("url")

	urlData, err := database.GetUrl(shortUrl)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusNotFound, "URL not found")
	}

	return ctx.JSON(http.StatusOK, models.LongUrlData{
		Url: urlData.LongUrl,
	})
}
