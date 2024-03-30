package tests

import (
	g "like/grammar"
	. "like/grammar/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {
	DescribeTable("T", func(f string, e string) {
		var c = Read(f)

		context, result := MakeContext()

		err := g.Execute(&context, c)
		Expect(err).To(BeNil())
		Expect(result.String()).To(Equal(e))
	},
		Entry("samples/sample.like", "samples/sample.like", "b\n"),
		Entry("samples/simple_template.like", "samples/simple_template.like", "aaacbbb\n"),
		Entry("samples/lambda_invoke", "samples/lambda_invoke.like", "_1_2_\n"),
		Entry("samples/lambda_invoke_space", "samples/lambda_invoke_space.like", "_ 1 2 _\n"),
		Entry("samples/func", "samples/func.like", "one\na oops\na oops\n"))
})
