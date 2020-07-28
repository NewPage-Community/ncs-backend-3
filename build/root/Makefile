.PHONY: all
all: dep build test lint

.PHONY: dep
dep:
	go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle -- update -exclude .cache

.PHONY: push_docker
push_docker:
	build/push-images.sh

.PHONY: push_docker_changed
push_docker_changed:
	build/push-images.sh changed

.PHONY: build
build:
	bazel build -k //app/... //pkg/...

.PHONY: test
test:
	bazel test -k //app/... //pkg/...

.PHONY: clean
clean:
	bazel clean

%.pb: %.proto
	protoc -I. --gofast_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. $<
	mockgen -source=$@.go -destination=$@.mock.go -package grpc

.PHONY: deploy
deploy:
	helm upgrade --install --namespace ncs ncs ./chart

.PHONY: lint
lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run -v
	helm lint ./chart
