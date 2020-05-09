load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_push")

def push_docker(name, binary):
    go_image(
        name = name,
        binary = binary,
    )
    container_push(
        name = "push_" + name,
        image = ":" + name,
        format = "Docker",
        registry = "harbor.new-page.xyz",
        repository = "newpage/ncs/" + name,
    )