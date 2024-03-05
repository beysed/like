load("@bazel_skylib//lib:shell.bzl", "shell")

def _ginkgo_test_impl(ctx):
    wrapper = ctx.actions.declare_file(ctx.label.name)
    ctx.actions.write(
        content = """#!/usr/bin/env bash
set -e
exec {ginkgo} {ginkgo_args} {go_test} -- "$@"
""".format(
            ginkgo = shell.quote(ctx.executable._ginkgo.short_path),
            ginkgo_args = " ".join([shell.quote(arg) for arg in ctx.attr.ginkgo_args]),
            # Ginkgo requires the precompiled binary end with ".test".
            go_test = shell.quote(ctx.executable.go_test.short_path + ".test"),
        ),
        is_executable = True,
        output = wrapper,
    )

    return [DefaultInfo(
        executable = wrapper,
        runfiles = ctx.runfiles(
            files = ctx.files.data,
            symlinks = {ctx.executable.go_test.short_path + ".test": ctx.executable.go_test},
            transitive_files = depset([], transitive = [ctx.attr._ginkgo.default_runfiles.files, ctx.attr.go_test.default_runfiles.files]),
        ),
    )]

ginkgo_test = rule(
    attrs = {
        "data": attr.label_list(allow_files = True),
        "go_test": attr.label(
            cfg = "target",
            executable = True,
        ),
        "ginkgo_args": attr.string_list(),
        "_ginkgo": attr.label(
            cfg = "target",
            #default = "//vendor/github.com/onsi/ginkgo/ginkgo",
            default = "ginkgo",
            executable = True,
        ),
    },
    executable = True,
    test = True,
    implementation = _ginkgo_test_impl,
)
