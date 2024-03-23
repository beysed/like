package tests

import (
	g "like/grammar"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grammar", func() {
	It("Parses: empty ", func() {
		_, e := g.Parse("a.like", Text(""))
		Expect(e).To(BeNil())
	})

	DescribeTable("Parses: correct indentifiers", func(input string, expected string) {
		actual := ParseInupt(input, "identifier", false)

		Expect(actual).To(BeAssignableToTypeOf(""))
		result := actual.(string)

		Expect(result).To(Equal(expected))
	},
		Entry("single", "a", "a"),
		Entry("multuple", "Aa", "Aa"))

	DescribeTable("Parses: include directive", func(input string, fn string) {
		var actual = ParseInupt(input, "directive", false)

		Expect(actual).To(BeAssignableToTypeOf(g.Include{}))
		result := actual.(g.Include)

		Expect(result.FileName).To(Equal(fn))
	},
		Entry("unquoted", "// include asd", "asd"),
		Entry("unquoted comment", "// include asd#comment", "asd"),
		Entry("unquoted comment space", "// include asd #comment", "asd"),
		Entry("single quoted", "// include 'asd'", "'asd'"),
		Entry("double quoted", "// include \"asd\"", "\"asd\""))

	DescribeTable("Parses: value / memeber_list", func(input string) {
		var actual = ParseInupt(input, "member", false)

		Expect(actual).To(BeAssignableToTypeOf(g.Member{}))
		result := actual.(g.Member)
		Expect(result.String()).To(Equal(input))
	},
		Entry("simple no index", "a"),
		Entry("double index", "a[0][1]"),
		Entry("tripple index", "a[0][1][2]"),
		Entry("single index", "a[0]"))

	DescribeTable("Parses: value", func(input string) {
		var actual = ParseInupt(input, "value", false)

		Expect(actual).To(BeAssignableToTypeOf(g.Value{}))
		result := actual.(g.Value)
		Expect(result.String()).To(Equal(input))
	},
		Entry("simple value", "a"),
		Entry("prefixed value", "$a"))
})
