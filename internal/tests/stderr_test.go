package tests

import (
	g "github.com/beysed/like/internal/grammar"
	. "github.com/beysed/like/internal/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stderr", func() {
	DescribeTable("Outputs", func(f string, stdout string, stderr string) {
		var c = Read(f)

		context, result := MakeContext()
		_, locals := context.Locals.Peek()
		locals.Store["_shell"] = GetShell()

		_, err := g.Execute("a.like", context, c)
		Expect(err).To(BeNil())
		Expect(result.Stdout.String()).To(Equal(stdout))
		Expect(result.Stderr.String()).To(Equal(stderr))
	},
		Entry("samples/lambda_pipe", "samples/lambda_pipe.like", "out: A\n", "err: <nil>\n"),
		Entry("samples/stderr_lose", "samples/stderr_lose.like", "", "fake-err"),
		Entry("samples/stderr_pipe2pipe2", "samples/stderr_pipe2pipe2.like", "faked(stdin:ascd; args:)\nfake-err\n", ""),
		Entry("samples/stderr_pipe2pipe", "samples/stderr_pipe2pipe.like", "faked(stdin:faked(stdin:; args:); args:)faked(stdin:fake-err; args:)", "fake-errfake-err"),
		Entry("samples/stderr", "samples/stderr.like", "stdout1stdout2\n", "stderr1stderr2\n"),
		Entry("samples/stderr_pipe", "samples/stderr_pipe.like", "faked(stdin:; args:)\nfake-err\n", ""),
	)
})
