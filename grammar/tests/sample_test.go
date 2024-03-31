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
		Entry("samples/sample", "samples/sample.like", "b\n"),
		Entry("samples/simple_template", "samples/simple_template.like", "aaacbbb\n"),
		Entry("samples/lambda_invoke", "samples/lambda_invoke.like", "_1_2_\n"),
		Entry("samples/lambda_invoke_space", "samples/lambda_invoke_space.like", "_ 1 2 _\n"),
		Entry("samples/func", "samples/func.like", "one\na oops\na oops\n"),
		Entry("samples/execute", "samples/execute.like", "AAA\n\n"),
		Entry("samples/assigns", "samples/assigns.like", "1\n"),
		Entry("samples/template", "samples/template.like", "AfredB\n___\n"))

	DescribeTable("Valid file", func(f string, e string) {
		context, result := MakeContext()

		err := g.Execute(&context, []byte(f))
		Expect(err).To(BeNil())
		Expect(result.String()).To(Equal(e))
	},
		Entry("empty", "", ""),
		Entry("line w/o line break", "` a", "a\n"),
		Entry("line with line break", "` a\n", "a\n"))
})
