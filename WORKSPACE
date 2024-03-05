workspace(name = "like")

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

# gazelle:repo bazel_gazelle

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:fQnZVsXk8uxXIStYb0N4bGk7jeyTalG/wsZjQ25dO0g=",
    version = "v1.17.2",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
    version = "v4.20.0+incompatible",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_smarty_assertions",
    importpath = "github.com/smarty/assertions",
    sum = "h1:cR//PqUBUiQRakZWqBiFFQ9wb8emQGDb0HeGdqGByCY=",
    version = "v1.15.0",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    importpath = "github.com/smartystreets/goconvey",
    sum = "h1:qGjIddxOk4grTu9JPOU31tVfq3cNdBlNa5sSznIX1xY=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:hDPOHmpOpP40lSULcqw7IrRb/u7w6RpDC9399XyoNd0=",
    version = "v1.6.1",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
    version = "v1.4.13",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
    version = "v0.0.0-20161208181325-20d25e280405",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
    version = "v3.0.1",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:SernR4v+D55NyBH2QiEQrlBAnj1ECL6AGrA5+dPaMY8=",
    version = "v0.15.0",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:AQyQV4dYCvJ7vGmJyKki9+PBdyvhkSd8EIx/qb0AYv4=",
    version = "v0.21.0",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:5BMeUDZ7vkXGfEr1x9B4bRcTH4lpkTkpdh0T/J+qjbQ=",
    version = "v0.6.0",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:DBdB3niSjOA/O0blCZBqDefyWNYveAYMNF1Wum0DYQ4=",
    version = "v0.18.0",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:k8NLag8AGHnn+PHbl7g43CtqZAwG60vZkLqgyZgIHgQ=",
    version = "v0.18.0",
)

go_repository(
    name = "com_github_samber_lo",
    importpath = "github.com/samber/lo",
    sum = "h1:4gTz1wUhNYLhFSKl6O+8peW0v2F4BCY034GRpU9WnuA=",
    version = "v1.39.0",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:3MTrJm4PyNL9NBqvYDSj3DHl46qQakyfqfWo4jgfaEM=",
    version = "v0.0.0-20220303212507-bbda1eaf7a17",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:8JEhPFa5W2WU7YfeZzPNqzMP6Lwt7L2715Ggo0nosvA=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    sum = "h1:pKouT5E8xu9zeFC39JXRDukb6JFQPXM5p5I91188VAQ=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_go_task_slim_sprig",
    importpath = "github.com/go-task/slim-sprig",
    sum = "h1:tfuBGBXKqDEevZMzYi5KSi8KkcZtzBcTgAUUtapy0OI=",
    version = "v0.0.0-20230315185526-52ccab3ef572",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
    version = "v1.5.3",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_google_pprof",
    importpath = "github.com/google/pprof",
    sum = "h1:yAJXTCF9TqKcTiHJAE8dj7HMvPfh66eeA2JYW7eFpSE=",
    version = "v0.0.0-20210407192527-94a9f03dee38",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_nxadm_tail",
    importpath = "github.com/nxadm/tail",
    sum = "h1:8feyoE3OzPrcshW5/MJ4sGESc5cqmGkGCWlco4l0bqY=",
    version = "v1.4.11",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:8xi0RTUf59SOSfEtZMvwTvXYMzG4gV23XVHOZiXNtnE=",
    version = "v1.16.5",
)

go_repository(
    name = "com_github_onsi_ginkgo_v2",
    importpath = "github.com/onsi/ginkgo/v2",
    sum = "h1:7q1w9frJDzninhXxjZd+Y/x54XNjG/UlRLIYPZafsPM=",
    version = "v2.16.0",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:KYppCUK+bUgAZwHOu7EXVBKyQA6ILvOESHkn/tgoqvo=",
    version = "v1.31.1",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
    version = "v1.4.7",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:clyUAQHOM3G0M3f5vQj7LuJrETvjVot3Z5el9nffUtU=",
    version = "v2.3.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:w43yiav+6bVFTBQFZX0r7ipe9JQ1QsbMgHwbBziscLw=",
    version = "v1.28.0",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:ENy+Az/9Y1vSrlrvBSyna3PITt4tiZLf7sgCjZBX7Wo=",
    version = "v0.19.0",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:mkTF7LCd6WGJNL3K1Ad7kwxNfYAW6a8a8QqtMblp/4U=",
    version = "v0.17.0",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=",
    version = "v0.14.0",
)

go_repository(
    name = "com_github_chzyer_logex",
    importpath = "github.com/chzyer/logex",
    sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_readline",
    importpath = "github.com/chzyer/readline",
    sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name = "com_github_chzyer_test",
    importpath = "github.com/chzyer/test",
    sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version = "v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name = "com_github_ianlancetaylor_demangle",
    importpath = "github.com/ianlancetaylor/demangle",
    sum = "h1:mV02weKRL81bEnm8A0HT1/CAelMQDBuQIfLw8n+d6xI=",
    version = "v0.0.0-20200824232613-28f6c0f3b639",
)

go_repository(
    name = "org_golang_x_telemetry",
    importpath = "golang.org/x/telemetry",
    sum = "h1:+Kc94D8UVEVxJnLXp/+FMfqQARZtWHfVrcRtcG8aT3g=",
    version = "v0.0.0-20240208230135-b75ee8823808",
)

go_rules_dependencies()

go_register_toolchains(version = "1.22.0")

gazelle_dependencies()
