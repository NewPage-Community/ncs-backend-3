load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_download_sdk",
    "go_register_toolchains",
    "go_rules_dependencies"
)

def init_rules_go():
    go_download_sdk(
        name = "go_sdk",
        version = "1.17",
        urls = [
            "http://nexus3.nexus3:8081/repository/google/go/{}",
            "https://nexus.new-page.xyz/repository/google/go/{}",
            "https://dl.google.com/go/{}",
        ],
    )
    go_rules_dependencies()
    go_register_toolchains()
