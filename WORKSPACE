# gazelle:repository_macro build/repos.bzl%go_repositories
workspace(name = "ncs_backend")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "http://nexus3.nexus3:8081/repository/github/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://nexus.new-page.xyz/repository/github/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "http://nexus3.nexus3:8081/repository/github/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://nexus.new-page.xyz/repository/github/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "59d5b42ac315e7eadffa944e86e90c2990110a1c8075f1cd145f487e999d22b3",
    strip_prefix = "rules_docker-0.17.0",
    urls = [
        "http://nexus3.nexus3:8081/repository/github/bazelbuild/rules_docker/releases/download/v0.17.0/rules_docker-v0.17.0.tar.gz",
        "https://nexus.new-page.xyz/repository/github/bazelbuild/rules_docker/releases/download/v0.17.0/rules_docker-v0.17.0.tar.gz",
    ],
)

http_archive(
    name = "bazel_skylib",
    sha256 = "1c531376ac7e5a180e0237938a2536de0c54d93f5c278634818e0efc952dd56c",
    urls = [
        "http://nexus3.nexus3:8081/repository/github/bazelbuild/bazel-skylib/releases/download/1.0.3/bazel-skylib-1.0.3.tar.gz",
        "https://nexus.new-page.xyz/repository/github/bazelbuild/bazel-skylib/releases/download/1.0.3/bazel-skylib-1.0.3.tar.gz",
    ],
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
    strip_prefix = "protobuf-3.14.0",
    urls = [
        "http://nexus3.nexus3:8081/repository/github/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        "https://nexus.new-page.xyz/repository/github/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
    ],
)

http_archive(
    name = "zlib",
    build_file = "@com_google_protobuf//:third_party/zlib.BUILD",
    sha256 = "629380c90a77b964d896ed37163f5c3a34f6e6d897311f1df2a7016355c45eff",
    strip_prefix = "zlib-1.2.11",
    urls = [
        "http://nexus3.nexus3:8081/repository/github/madler/zlib/archive/v1.2.11.tar.gz",
        "https://nexus.new-page.xyz/repository/github/madler/zlib/archive/v1.2.11.tar.gz",
    ],
)

http_archive(
    name = "rules_cc",
    sha256 = "b295cad8c5899e371dde175079c0a2cdc0151f5127acc92366a8c986beb95c76",
    strip_prefix = "rules_cc-daf6ace7cfeacd6a83e9ff2ed659f416537b6c74",
    urls = [
        "http://nexus3.nexus3:8081/repository/github/bazelbuild/rules_cc/archive/daf6ace7cfeacd6a83e9ff2ed659f416537b6c74.zip",
        "https://nexus.new-page.xyz/repository/github/bazelbuild/rules_cc/archive/daf6ace7cfeacd6a83e9ff2ed659f416537b6c74.zip",
    ],
)

http_archive(
    name = "go_googleapis",
    patch_args = [
        "-E",
        "-p1",
    ],
    patches = [
        # find . -name BUILD.bazel -delete
        "@io_bazel_rules_go//third_party:go_googleapis-deletebuild.patch",
        # set gazelle directives; change workspace name
        "@io_bazel_rules_go//third_party:go_googleapis-directives.patch",
        # gazelle args: -repo_root .
        "@io_bazel_rules_go//third_party:go_googleapis-gazelle.patch",
    ],
    sha256 = "fa52a03d9c49c28e475a33f9adfb0551344ca0dbd15168d6862ac32cc8354c1c",
    strip_prefix = "googleapis-1c5c56c18ab4e2353f87a9c2e14864f218c5502c",
    # master, as of 2021-06-30
    urls = [
        "http://nexus3.nexus3:8081/repository/github/googleapis/googleapis/archive/1c5c56c18ab4e2353f87a9c2e14864f218c5502c.zip",
        "https://nexus.new-page.xyz/repository/github/googleapis/googleapis/archive/1c5c56c18ab4e2353f87a9c2e14864f218c5502c.zip",
    ],
)

http_archive(
    name = "rules_python",
    sha256 = "e5470e92a18aa51830db99a4d9c492cc613761d5bdb7131c04bd92b9834380f6",
    strip_prefix = "rules_python-4b84ad270387a7c439ebdccfd530e2339601ef27",
    urls = [
        "http://nexus3.nexus3:8081/repository/github/bazelbuild/rules_python/archive/4b84ad270387a7c439ebdccfd530e2339601ef27.tar.gz",
        "https://nexus.new-page.xyz/repository/github/bazelbuild/rules_python/archive/4b84ad270387a7c439ebdccfd530e2339601ef27.tar.gz",
    ],
)

load("//build:go.bzl", "init_rules_go")
init_rules_go()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

load("//build:docker.bzl", "init_rules_docker")
init_rules_docker()

load("//build:repos.bzl", "go_repositories")
go_repositories()