load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix like
gazelle(name = "gazelle")

go_library(
    name = "like_lib",
    srcs = ["main.go"],
    importpath = "like",
    visibility = ["//visibility:private"],
    deps = ["//grammar"],
)

go_binary(
    name = "like",
    embed = [":like_lib"],
    visibility = ["//visibility:public"],
)
