# Tools version
PROTOC_VERSION  = 3.11.4
GO_VERSION      = 1.14.4
BAZEL_VERSION   = 3.2.0
# Docker image
HELM_IMG        = alpine/helm
BAZEL_IMG       = registry.new-page.xyz/cloud-marketplace-containers/google/bazel:$(BAZEL_VERSION)
GO_IMG          = golang:$(GO_VERSION)-alpine
GOLANGCI_IMG    = golangci/golangci-lint:latest-alpine
# Golang
ifdef GOPATH
	GO_PATH_SET = --env GOPATH=/gopath -v $(GOPATH):/gopath
endif
# Docker command
PROXY_PASS      = --network host --env HTTP_PROXY=$(HTTP_PROXY) --env HTTPS_PROXY=$(HTTPS_PROXY) --env http_proxy=$(http_proxy) --env https_proxy=$(https_proxy) --env NO_PROXY=goproxy.cn,new-page.xyz
DOCKER_CMD      = docker run --rm -v $(shell pwd):/apps -w /apps --entrypoint "" $(PROXY_PASS)
GO_ENV			= --env GOPROXY=https://goproxy.cn,direct $(GO_PATH_SET)
# Tools run on docker
HELM_CMD        = $(DOCKER_CMD) $(HELM_IMG)
BAZEL_CMD       = $(DOCKER_CMD) $(BAZEL_IMG)
GO_CMD          = $(DOCKER_CMD) $(GO_ENV) $(GO_IMG)
GOLANGCI_CMD    = $(DOCKER_CMD) $(GO_ENV) $(GOLANGCI_IMG)

# Rules
.PHONY: all
all: img proto dao-mock dep build test lint

.PHONY: dep
dep:
	$(GO_CMD) go mod tidy
	$(BAZEL_CMD) \
		bazel version && \
		bazel run //:gazelle -- update-repos -from_file=go.mod && \
		bazel run //:gazelle -- update -exclude .cache

.PHONY: push_docker
push_docker:
	$(BAZEL_CMD) \
		bazel version && \
		./build/push-images.sh

.PHONY: push_docker_changed
push_docker_changed:
	$(BAZEL_CMD) \
		bazel version && \
		./build/push-images.sh changed

.PHONY: build
build:
	$(BAZEL_CMD) \
		bazel version && \
		bazel build -k --noshow_progress //app/... //pkg/...

.PHONY: test
test:
	$(BAZEL_CMD) \
		bazel version && \
		bazel test -k --noshow_progress //app/... //pkg/...

.PHONY: clean
clean:
	$(BAZEL_CMD) \
		bazel version && \
		bazel clean

.PHONY: proto
proto:
	$(GO_CMD) \
		go mod download && \
		build/build-proto.sh app

.PHONY: deploy
deploy:
	$(HELM_CMD) helm upgrade --install --namespace ncs ncs ./chart

.PHONY: lint
lint:
	$(GOLANGCI_CMD) golangci-lint run -v --timeout 2m
	$(HELM_CMD) helm lint ./chart

.PHONY: img
img:
	docker pull $(HELM_IMG)
	docker pull $(BAZEL_IMG)
	docker pull $(GO_IMG)

.PHONY: dao-mock
dao-mock:
	$(GO_CMD) \
		go mod download && \
		build/build-dao-mock.sh app
