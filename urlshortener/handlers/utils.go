package handlers

import (
	"crypto/sha256"
	"encoding/base64"
)

// return base64 encoded sha256 hash of the input string
func getSHA256HashAsBase64(text string) string {
	hash := sha256.Sum256([]byte(text))
	return base64.StdEncoding.EncodeToString(hash[:])
}

// return the first 6 characters of the base64 encoded sha256 hash of the input url
// this allows for 64^6 (or ~68 billion) possible encoded url values
func ShortenURL(url string) string {
	hash := getSHA256HashAsBase64(url)
	return hash[:6]
}
