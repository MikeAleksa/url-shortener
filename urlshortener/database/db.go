package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/MikeAleksa/url-shortener/models"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var db_context = context.Background()

// ConnectDB opens a connection to the database
func ConnectDB() {
	// Get environment variables for database connection
	db_port := os.Getenv("POSTGRES_PORT")
	db_host := os.Getenv("POSTGRES_HOST")
	db_user := os.Getenv("POSTGRES_USER")
	db_pass := os.Getenv("POSTGRES_PASSWORD")
	db_name := os.Getenv("POSTGRES_NAME")

	// Establish connection to PostgreSQL database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db_host, db_port, db_user, db_pass, db_name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db
}

func CloseDB() {
	DB.Close()
}

func CreateUrl(urlData models.UrlData) error {
	var err error

	// check if short URL already exists
	_, err = GetUrl(urlData.ShortUrl)
	if err == nil {
		// short URL already exists in database - do nothing
		return nil
	}

	queryStatement := `
	INSERT INTO url (shorturl, longurl, redirects) VALUES ($1, $2, $3);
	`
	_, err = DB.ExecContext(db_context, queryStatement, urlData.ShortUrl, urlData.LongUrl, 0)
	if err != nil {
		return err
	}
	return nil
}

func GetUrl(shortUrl string) (models.UrlData, error) {
	var err error
	var urlData models.UrlData
	queryStatement := `
	SELECT shorturl, longurl, redirects FROM url WHERE shorturl = $1;
	`
	results := DB.QueryRowContext(db_context, queryStatement, shortUrl)
	err = results.Scan(&urlData.ShortUrl, &urlData.LongUrl, &urlData.Redirects)
	if err != nil {
		return models.UrlData{}, err
	}
	return urlData, nil
}

func IncrementRedirectCount(shortUrl string) error {
	var err error
	queryStatement := `
	UPDATE url SET redirects = redirects + 1 WHERE shorturl = $1;
	`
	_, err = DB.ExecContext(db_context, queryStatement, shortUrl)
	if err != nil {
		return err
	}
	return nil
}
