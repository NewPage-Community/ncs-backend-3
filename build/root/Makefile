.PHONY: all
all: dep build test

.PHONY: dep
dep:
	go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle

.PHONY: push_docker
push_docker:
	build/push-images.sh

.PHONY: push_docker_changed
push_docker_changed:
	build/push-images.sh changed

.PHONY: build
build:
	bazel build //app/... //pkg/...

.PHONY: test
test:
	bazel test //app/... //pkg/...

.PHONY: clean
clean:
	bazel clean

%.pb: %.proto
	protoc -I. --gofast_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. $<
	mockgen -source=$@.go -destination=$@.mock.go -package grpc
