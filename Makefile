.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go && cp -r ./configs/ ./.bin && cp ./.env ./.bin/

run: build
	docker-compose up --remove-orphans app

rebuild:
	docker-compose up -d --no-deps --build

run-stade: build
	./.bin/app