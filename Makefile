SHELL := /usr/bin/sh
PROJECT_NAME = nft-service
BUILD_FOLDER = ./dist
VERSION=0.0.1
BINARY_PATH = $(BUILD_FOLDER)/$(PROJECT_NAME)-v${VERSION}
HOST=10.122.118.229
SERVICE_NAME=nft-service

.PHONY: dev-setup
dev-setup:
	echo "Setting up development environment..."
	go mod tidy
	go mod vendor
	cp .env.example .env || true
	docker compose up -d 

.PHONY: build
build:
	echo "Building $(PROJECT_NAME)..."
	mkdir -p $(BUILD_FOLDER)
	go mod tidy
	go mod vendor
	export GOPRIVATE=github.com/blcvn/*
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY_PATH}
	ssh ${HOST} "sudo systemctl stop ${SERVICE_NAME}.service"
	scp $(BINARY_PATH) ${HOST}:/opt/${PROJECT_NAME}
	ssh ${HOST} "sudo systemctl start ${SERVICE_NAME}.service"
	# ssh ${HOST} "cd /opt/${PROJECT_NAME}/${SERVICE_NAME} && docker restart ${SERVICE_NAME}"


build-image:
	docker save -o $(BUILD_FOLDER)/images.tar postgres:latest ipfs/kubo:latest
	scp $(BUILD_FOLDER)/images.tar ${HOST}:/opt/${PROJECT_NAME}

load-image:
	ssh ${HOST} "docker load -i /opt/${PROJECT_NAME}/images.tar"

tunel:
	ssh -L 5432:localhost:5432 -L 5001:localhost:5001 ${HOST}