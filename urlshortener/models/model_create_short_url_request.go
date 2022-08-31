package models

type CreateShortUrlRequest struct {

	// The long URL that will be shortened
	Url string `json:"url"`
}
