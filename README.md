# URL Shortener

A URL shortener using Go, PostgreSQL, OpenAPI, and Swagger.

## Requirements

- Docker (tested with version 20.10.16)
- Go (for development, tested with version 1.16)
  ...

## Quick-Start

Start services:

```sh
git clone git@github.com:MikeAleksa/url-shortener.git
cd url-shortener
docker compose build && docker-compose up -d
```

URL Shortener API will be available at localhost:8000

Swagger API documentation will be available at localhost:8080

## Development

Run from the root of the repository

### Generate Go code from OpenAPI specifications

```sh
docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/openapi/v1.json \
  -g go-echo-server \
  -o /local/urlshortener \
  --additional-properties=packageName=api,serverPort=5000
```

## Further Development

If time allowed, I would also implement the following:

- Fix Swagger error that doesn't allow web UI to make API calls due to OPTIONS method (cURL and Postman both work for making API calls)

- Validation

  - check that a long url is a valid url format before processing

- Testing

  - unit tests of API handlers with mocked database calls
  - integration tests using Postman

- Redis

  - caching most recent queries (https://developer.redis.com/howtos/redisjson/using-go/)
  - rate limiting (https://redis.com/redis-best-practices/basic-rate-limiting/)

- Add IaC and CI/CD pipelines

  - use a multistage build to reduce the size of the application docker container and save it as an artifact in the build pipeline
  - automated unit and integration testings
  - automated deployments

- Security hardening before deployment
  - remove .env file that was included for demonstration purposes, and inject environment variables in deployment pipeline
