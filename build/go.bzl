load(
    "@io_bazel_rules_go//go/private:sdk_list.bzl",
    "DEFAULT_VERSION",
)

load(
    "@io_bazel_rules_go//go/private:sdk.bzl",
    "go_host_sdk",
    "go_download_sdk",
    "go_register_nogo",
)

def go_register_toolchains(go_version = None, nogo = None):
    """See /go/toolchains.rst#go-register-toolchains for full documentation."""
    sdk_kinds = ("_go_download_sdk", "_go_host_sdk", "_go_local_sdk", "_go_wrap_sdk")
    existing_rules = native.existing_rules()
    sdk_rules = [r for r in existing_rules.values() if r["kind"] in sdk_kinds]
    if len(sdk_rules) == 0 and "go_sdk" in existing_rules:
        # may be local_repository in bazel_tests.
        sdk_rules.append(existing_rules["go_sdk"])

    if go_version and len(sdk_rules) > 0:
        fail("go_version set after go sdk rule declared ({})".format(", ".join([r["name"] for r in sdk_rules])))
    if len(sdk_rules) == 0:
        if not go_version:
            go_version = DEFAULT_VERSION
        if go_version == "host":
            go_host_sdk(name = "go_sdk")
        else:
            if not versions.is_at_least(MIN_SUPPORTED_VERSION, go_version):
                print("DEPRECATED: go_register_toolchains: support for Go versions before {} will be removed soon".format(MIN_SUPPORTED_VERSION))
            go_download_sdk(
                name = "go_sdk",
                version = go_version,
                urls = [
                    "https://raw.new-page.xyz/go/{}",
                    "https://dl.google.com/go/{}",
                    ],
            )

    if nogo:
        # Override default definition in go_rules_dependencies().
        go_register_nogo(
            name = "io_bazel_rules_nogo",
            nogo = nogo,
        )