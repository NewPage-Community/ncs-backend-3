load(
    "@io_bazel_rules_docker//go:image.bzl",
    "go_image",
    go_image_repos = "repositories",
)
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_image",
    "container_push",
    "container_pull",
)
load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repos = "repositories",
)
load("@io_bazel_rules_docker//go:go.bzl", BASE_DIGESTS = "DIGESTS")
load("@io_bazel_rules_docker//go:static.bzl", STATIC_DIGESTS = "DIGESTS")
load("@bazel_gazelle//:deps.bzl", "go_repository")

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
        registry = "registry-in.new-page.xyz",
        repository = "newpage/ncs/" + app_name,
        tag = "{STABLE_DOCKER_TAG}",
    )

def init_rules_docker():
    container_pull(
        name = "go_image_base",
        digest = BASE_DIGESTS["latest"],
        registry = "gcr.io",
        repository = "distroless/base",
    )
    container_pull(
        name = "go_debug_image_base",
        digest = BASE_DIGESTS["debug"],
        registry = "gcr.io",
        repository = "distroless/base",
    )
    container_pull(
        name = "go_image_static",
        digest = STATIC_DIGESTS["latest"],
        registry = "gcr.io",
        repository = "distroless/static",
    )
    container_pull(
        name = "go_debug_image_static",
        digest = STATIC_DIGESTS["debug"],
        registry = "gcr.io",
        repository = "distroless/static",
    )
    go_repository(
        name = "com_github_google_go_containerregistry",
        build_directives = [
            # Silence Go module warnings about unused modules.
            "gazelle:exclude pkg/authn/k8schain",
        ],
        importpath = "github.com/google/go-containerregistry",
        sha256 = "c3e28d8820056e7cc870dbb5f18b4f7f7cbd4e1b14633a6317cef895fdb35203",
        strip_prefix = "go-containerregistry-0.5.1",
        urls = [
            "http://nexus3.nexus3:8081/repository/github/google/go-containerregistry/archive/v0.5.1.tar.gz",
            "https://nexus.new-page.xyz/repository/github/google/go-containerregistry/archive/v0.5.1.tar.gz",
        ],
    )
    container_repos()
    go_image_repos()