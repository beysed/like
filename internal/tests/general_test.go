package tests

import (
	g "github.com/beysed/like/internal/grammar"
	. "github.com/beysed/like/internal/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grammar", func() {
	It("Parses: empty ", func() {
		_, e := g.Parse("a.like", Text(""))
		Expect(e).To(BeNil())
	})

	DescribeTable("Parses: correct indentifiers", func(input string, expected string) {
		actual := ParseInupt(input, "identifier")

		Expect(actual).To(BeAssignableToTypeOf(""))
		result := actual.(string)

		Expect(result).To(Equal(expected))
	},
		Entry("single", "a", "a"),
		Entry("multuple", "Aa", "Aa"))

	DescribeTable("Parses: include directive", func(input string, fn string) {
		var actual = ParseInupt(input, "directive")

		Expect(actual).To(BeAssignableToTypeOf(g.Include{}))
		result := actual.(g.Include)

		a := result.FileName.String()
		Expect(a).To(Equal(fn))
	},
		Entry("unquoted", "#include asd", "asd"),
		Entry("unquoted comment", "#include asd#comment", "asd"),
		Entry("unquoted comment space", "#include asd #comment", "asd"),
		Entry("single quoted", "#include 'asd'", "'asd'"),
		Entry("double quoted", "#include \"asd\"", "\"asd\""))

	DescribeTable("Parses: expression / member", func(input string) {
		var actual = ParseInupt(input, "expression")

		result, ok := actual.(g.Expression)
		Expect(ok).To(BeTrue())

		Expect(result.String()).To(Equal(input))
	},
		Entry("simple no index", "a"),
	)

	DescribeTable("Parses: reference", func(input string, expect string) {
		var actual = ParseInupt(input, "reference")

		result, ok := actual.(g.Reference)
		Expect(ok).To(BeTrue())

		Expect(result.String()).To(Equal(expect))
	},
		Entry("a", "$a", "$a"),
		Entry("simple index", "$a[0]", "$a[0]"),
		Entry("a[0][1]", "$a[0][1]", "$a[0][1]"),
		Entry("a[0][1][2]", "$a[0][1][2]", "$a[0][1][2]"))
})
