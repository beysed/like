bazel run @go_sdk//:bin/go -- mod init
bazel run @go_sdk//:bin/go -- mod tidy
bazel run @io_bazel_rules_go//go -- mod tidy -v

bazel run @io_bazel_rules_go//go
bazel run //:gazelle

bazel run //:gazelle -- update-repos -from_file go.mod

The latest versions are always listed on https://registry.bazel.build/.



