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
		locals.Store["_shell"] = getShell()

		_, err := g.Execute("a.like", context, c)
		Expect(err).To(BeNil())
		Expect(result.Stdout.String()).To(Equal(stdout))
		Expect(result.Stderr.String()).To(Equal(stderr))
	},

		Entry("samples/stderr_pipe2pipe", "samples/stderr_pipe2pipe.like", "faked(faked(:):)faked(fake-err:)", "fake-errfake-err"),
		Entry("samples/stderr", "samples/stderr.like", "stdout1stdout2\n", "stderr1stderr2\n"),
		Entry("samples/stderr_pipe", "samples/stderr_pipe.like", "faked(:)\nfake-err\n", ""),
	)
})
