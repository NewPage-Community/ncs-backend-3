.PHONY: dep
dep:
	go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle

.PHONY: push_docker
push_docker:
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //app/user/account/cmd:account

.PHONY: build
build:
	bazel build //...

.PHONY: test
test:
	bazel test //...

.PHONY: clean
clean:
	bazel clean
