package tests

import (
	g "github.com/beysed/like/internal/grammar"
	. "github.com/beysed/like/internal/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Literals", func() {

	DescribeTable("Incorrect literal", func(input string) {
		var _, err = ParseInuptIncorrect(input, "literal")
		Expect(err).NotTo(BeNil())
	},
		Entry("#a", "#a"),
		Entry("{a", "{a"),
		Entry("(a", "(a"))

	DescribeTable("Parses: correct escaped literal", func(input string, expected string) {
		var actual = ParseInupt(input, "literal")

		result, ok := actual.(g.Literal)
		Expect(ok).To(BeTrue())

		Expect(result.String()).To(Equal(expected))
	},
		Entry("dbl bslsh", "\\", "\\"),
		Entry("bslsh b", "\\b", "\\b"),
		Entry("bslsh", "\\", "\\"),
		Entry("bslsh @", "\\@", "@"))

	DescribeTable("Parses: correct literal", func(input string) {
		var actual = ParseInupt(input, "literal")

		result, ok := actual.(g.Literal)
		Expect(ok).To(BeTrue())

		Expect(result.String()).To(Equal(input))
	},
		Entry("a", "a"),
		Entry("-b", "-b"),
		Entry("_b", "-b"),
		Entry("/b", "/b"),

		Entry("1b", "1b"),
		Entry("'asd'", "asd"),
		Entry("\"asd\"", "asd"),
		Entry("10", "10"))
})
