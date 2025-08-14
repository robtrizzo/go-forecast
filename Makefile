# in a real project, sensitive information would be passed in via env files
# include .envrc

## run/api: run the cmd/api application for development
.PHONY: run/api
run/api:
	go run ./cmd/api -port=4000 -cors-trusted-origins="http://localhost:4000" -weather-url="https://api.weather.gov"

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api