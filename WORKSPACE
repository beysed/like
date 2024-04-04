workspace(name = "like")

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

# ga1zelle:repo bazel_gazelle
# gazelle:repository_macro deps.bzl%go_dependencies
load("//:deps.bzl", "go_dependencies")

go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.22.2")

gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
git_repository(
    name = "bazel_sonarqube",
    commit = "d6109e6627a00ad84c24c38df5d2d17159203f47",
    remote = "https://github.com/Zetten/bazel-sonarqube.git",
)

load("@bazel_sonarqube//:repositories.bzl", "bazel_sonarqube_repositories")

bazel_sonarqube_repositories()
