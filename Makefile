VERSION?=$(shell git describe --tags --match "v*")
IMAGE_REGISTRY=ghcr.io/ebiiim
IMAGE_NAME=toy-scheduler:$(VERSION)

.PHONY: build-image
build-image:
	docker build -f ./Dockerfile --build-arg VERSION="$(VERSION)" -t $(IMAGE_REGISTRY)/$(IMAGE_NAME) .
