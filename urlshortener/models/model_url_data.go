package models

type UrlData struct {

	// The long URL that will be shortened
	LongUrl string `json:"longUrl,omitempty"`

	// The shortened URL
	ShortUrl string `json:"shortUrl,omitempty"`

	// Click analytics for shortened URL
	Redirects int32 `json:"redirects,omitempty"`
}
