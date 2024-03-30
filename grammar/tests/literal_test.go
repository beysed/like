package tests

import (
	g "like/grammar"
	. "like/grammar/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Literals", func() {

	DescribeTable("Incorrect literal", func(input string) {
		var _, err = ParseInuptIncorrect(input, "literal", false)
		Expect(err).NotTo(BeNil())
	},
		Entry("#a", "#a"),
		Entry("{a", "{a"),
		Entry("(a", "(a"))

	DescribeTable("Parses: correct literal", func(input string) {
		var actual = ParseInupt(input, "literal", false)

		result, ok := actual.(g.Literal)
		Expect(ok).To(BeTrue())

		Expect(result.String()).To(Equal(input))
	},
		Entry("a", "a"),
		Entry("-b", "-b"),
		Entry("_b", "-b"),
		Entry("/b", "/b"),
		Entry("\\b", "\\b"),
		Entry("1b", "1b"),
		Entry("'asd'", "'asd'"),
		Entry("\"asd\"", "\"asd\""),
		Entry("10", "10"))

	DescribeTable("Parses: literal list", func(input string, expected []string) {
		var actual = ParseInupt(input, "literal_list", false)

		result, ok := actual.(g.LiteralList)
		Expect(ok).To(BeTrue())
		Expect(result).To(Equal(g.MakeLiteralList(expected)))
	},
		Entry("single", "a", []string{"a"}),
		Entry("multi", "a b", []string{"a", "b"}),
		Entry("multi extra space", "_\ta	0", []string{"_", "a", "0"}))
})
