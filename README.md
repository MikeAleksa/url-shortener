# URL Shortener

A URL shortener using Go and PostgreSQL

## Requirements

- Docker (tested with version 20.10.16)
- Go (for development, tested with version 1.16.2)
  ...

## Quick-Start

Start services:

```sh
docker-compose up -d
```

Swagger will be available at localhost:8080
URL Shortener API will be available at localhost:8000

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

...

- Redis - caching most recent, rate limiting
- IaC - terraform for deployment on AWS or K8s
- CI/CD - pipeline for redeploying on updates, minified dockerfile with prebuilt binary (and without Docker-in-Docker) saved as artifact and deployed
- Postman Integration Tests
- Remove .env (for demonstration purposes)

- Unit Test
- Swagger
- OpenAPI generator
