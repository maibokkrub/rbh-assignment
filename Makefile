# Define the version and other build information
IMAGE_NAME = maibokkrub/rbh-assignment
PLATFORM = linux/amd64,linux/arm64
EXECUTABLE = main
BUILDER_NAME = builder
COMMIT_SHA = $(shell git rev-parse --short HEAD)
BUILD_DATE = $(shell date +%Y-%m-%d)

GO = go
GOFLAGS =

.PHONY: all build builder_init builder_rm docker push clean

all: build

# Build the Go executable
build:
	$(GO) build $(GOFLAGS) -o $(EXECUTABLE) -ldflags="-X main.version=$(COMMIT_SHA) -X main.buildDate=$(BUILD_DATE)"

builder_init: 
	docker buildx create --use --name $(BUILDER_NAME)

builder_rm:
	docker buildx stop $(BUILDER_NAME)
	docker buildx rm $(BUILDER_NAME)

# Build a Docker image for the Go program
docker:
	docker buildx build --platform $(PLATFORM) -f deployments/Dockerfile -t $(IMAGE_NAME):$(COMMIT_SHA) -t $(IMAGE_NAME):latest .

# Push the Docker image to the repository
push: 
	docker
	docker push $(DOCKER_REPOSITORY)/$(EXECUTABLE):$(COMMIT_SHA)

# Clean up build artifacts
clean:
	rm -f $(EXECUTABLE)
