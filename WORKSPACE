workspace(name = "like")

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

# ga1zelle:repo bazel_gazelle
# gazelle:repository_macro deps.bzl%go_dependencies
load("//:deps.bzl", "go_dependencies")
go_dependencies()
go_rules_dependencies()

go_register_toolchains(version = "host")
gazelle_dependencies()
