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
	build/build.sh

.PHONY: build_changed
build_changed:
	build/build.sh changed

.PHONY: test
test:
	bazel test //...

.PHONY: clean
clean:
	bazel clean
