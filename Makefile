VERSION?=$(shell git describe --tags --match "v*")
IMAGE_REGISTRY=ghcr.io/ebiiim
IMAGE_NAME=toy-scheduler:$(VERSION)

.PHONY: build-image
build-image:
	docker build -f ./Dockerfile -t $(IMAGE_REGISTRY)/$(IMAGE_NAME) .

.PHONY: build-bin
build-bin:
	go build -ldflags "-X k8s.io/component-base/version.gitVersion=$(VERSION) -w" -o bin/kube-scheduler cmd/scheduler/main.go
