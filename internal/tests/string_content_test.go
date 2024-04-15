package tests

import (
	g "github.com/beysed/like/internal/grammar"
	. "github.com/beysed/like/internal/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("String Content", func() {

	DescribeTable("Parses: string content", func(input string, expected string) {
		var actual = ParseInupt(input, "string_content")

		result, ok := actual.(g.Expressions)
		Expect(ok).To(BeTrue())

		Expect(result.String()).To(Equal(expected))
	},
		Entry("case1", "aaa$a", "aaa$a"),
		Entry("case2", "aaa", "aaa"),
		Entry("case3", "\\$a", "$a"),
		Entry("case4", "\\$", "$"),
	)
})
