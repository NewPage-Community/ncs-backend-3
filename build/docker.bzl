load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")

def push_docker(name, app_name, embed):
    go_image(
        name = "app",
        embed = embed,
    )
    container_image(
        name = "image",
        base = ":app",
        stamp = True,
    )
    container_push(
        name = name,
        image = ":image",
        format = "Docker",
        registry = "harbor.new-page.xyz",
        repository = "newpage/ncs/" + app_name,
    )