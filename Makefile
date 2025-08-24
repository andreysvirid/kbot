REGISTRY=ghcr.io
REPO=den-andreysvirid/kbot
OS=linux
ARCH=amd64
TAG=$(shell git describe --tags --always)-$(OS)-$(ARCH)

build:
	docker buildx build --platform $(OS)/$(ARCH) -t $(REGISTRY)/$(REPO):$(TAG) .

push:
	docker push $(REGISTRY)/$(REPO):$(TAG)

helm-update:
	@sed -i "s|tag:.*|tag: \"$(TAG)\"|g" helm/values.yaml
