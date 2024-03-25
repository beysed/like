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

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
git_repository(
    name = "bazel_sonarqube",
    branch = "master",
    remote = "https://github.com/Zetten/bazel-sonarqube.git",
)

load("@bazel_sonarqube//:repositories.bzl", "bazel_sonarqube_repositories")

bazel_sonarqube_repositories()
