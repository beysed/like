﻿bazel run @go_sdk//:bin/go -- mod init
bazel run @go_sdk//:bin/go -- mod tidy
bazel run //:gazelle

bazel run //:gazelle -- update-repos -from_file go.mod

The latest versions are always listed on https://registry.bazel.build/.

