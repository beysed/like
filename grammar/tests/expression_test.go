package tests

import (
	g "like/grammar"

	. "like/grammar/tests/common"

	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samber/lo"
)

var _ = Describe("Grammar", func() {
	DescribeTable("Parses: list expressions with reference", func(input string) {
		var actual = ParseInupt(input, "writeln")

		_, ok := actual.(g.WriteLn)
		Expect(ok).To(BeTrue())
	},
		Entry("a", "` aaa$b"),
		Entry("b", "` aaa$(b)ccc"))

	DescribeTable("Parses: store access", func(input string, expected string) {
		var actual = ParseInupt(input, "store")

		res, ok := actual.(g.StoreAccess)
		Expect(ok).To(BeTrue())
		Expect(res.String()).To(Equal(expected))
	},
		Entry("a", "a", "a"),
		Entry("a.b", "a.b", "a.b"))

	DescribeTable("Parses: lambda assign", func(input string, expected string) {
		var actual = ParseInupt(input, "assign")

		res, ok := actual.(g.Assign)
		Expect(ok).To(BeTrue())
		Expect(res.String()).To(Equal(expected))
	},
		Entry("simple lambda", "a = (a) $a", "a = (a) $a"))

	DescribeTable("Parses: lambda", func(input string, expexted string) {
		var actual = ParseInupt(input, "lambda")

		res, ok := actual.(g.Lambda)
		Expect(ok).To(BeTrue())
		Expect(res.String()).To(Equal(expexted))
	},
		Entry("no arg lambda", "() $a", "() $a"),
		Entry("one arg lambda", "(a) $a", "(a) $a"),
		Entry("one arg list lambda", "(a) _ $a _", "(a) _ $a _"))

	DescribeTable("Parses: various files", func(input string, expexted string) {
		var actual = ParseInupt(input, "file")

		res, ok := actual.([]g.Expression)
		Expect(ok).To(BeTrue())
		r := strings.Join(lo.Map(res, func(e g.Expression, _ int) string { return e.String() }), "\n")
		Expect(r).To(Equal(expexted))
	},
		Entry("empty file", "", ""),
		Entry("empty line end", "\n", ""),
		Entry("empty line space end", " \n", ""),
		Entry("case1", "a=1", "a = 1"),
		Entry("code comment", "a=1#asd", "a = 1"),
		Entry("single comment", "#asd", ""),
		Entry("single line end comment", "#asd\n", ""),
		Entry("two comment lines", "#asd\n#def", "\n"),
		Entry("case2", "a=1\na=2", "a = 1\na = 2"),
		Entry("case3", "a=1\n\na=2", "a = 1\n\na = 2"),
		Entry("case4", "a=1\n \na=2", "a = 1\n\na = 2"),
		Entry("case5", "a=1\n#asd\na=2", "a = 1\n\na = 2"))

	DescribeTable("Parses: block lambda", func(input string, expexted string) {
		var actual = ParseInupt(input, "lambda")

		res, ok := actual.(g.Lambda)
		Expect(ok).To(BeTrue())
		Expect(res.String()).To(Equal(expexted))
	},
		Entry("no arg empty", "() {\n}", "() {\n\n}"),
		Entry("one arg call", "(a) {\nb()\n}", "(a) {\nb()\n}"),
		Entry("one arg assign", "(a) {\na=b\n}", "(a) {\na = b\n}"),
		Entry("one arg operator", "(a) {\n` a\n}", "(a) {\n` a\n}"),
		Entry("one arg space operator", "(a) {\n ` a\n}", "(a) {\n` a\n}"))

	DescribeTable("Parses: template", func(input string, expexted string) {
		var actual = ParseInupt(input, "template")

		res, ok := actual.(g.Assign)
		Expect(ok).To(BeTrue())
		Expect(res.String()).To(Equal(expexted))
	},
		Entry("one arg template empty", "`` t(a)\n\n``", "t = (a) "),
		Entry("one arg template reference", "`` t(a)\nA$(user)B\n___\n``", "t = (a) A$userB\n___"),
		Entry("one arg template literal", "`` t(a)\nABCD\n``", "t = (a) ABCD"),
		Entry("one arg multiline template literal", "`` t(a)\nAB\nCD\n``", "t = (a) AB\nCD"),
		Entry("one arg multiline reference template literal", "`` t(a)\n$(A)B$c\nCD\n``\n", "t = (a) $AB$c\nCD")) // todo: fix me (paren)

	DescribeTable("Parses: file template", func(input string, expexted string) {
		var actual = ParseInupt(input, "template")

		res, ok := actual.(g.Assign)
		Expect(ok).To(BeTrue())
		Expect(res.String()).To(Equal(expexted))
	},
		Entry("one arg template empty", "`` t(a)\nA$(b)A\n``", "t = (a) A$bA"))
})
