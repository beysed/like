bazel run @go_sdk//:bin/go -- mod init
bazel run @go_sdk//:bin/go -- mod tidy
bazel run @io_bazel_rules_go//go -- mod tidy -v

bazel run @io_bazel_rules_go//go
bazel run //:gazelle

bazel run //:gazelle -- update-repos -from_file go.mod

https://registry.bazel.build/.

<p align="right"><img src="https://sonarcloud.io/images/project_badges/sonarcloud-white.svg" alt="SonarCloud"/></p>