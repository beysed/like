package tests

import (
	g "like/grammar"

	. "like/grammar/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grammar", func() {
	DescribeTable("Parses: list expressions with reference", func(input string) {
		var actual = ParseInupt(input, "write")

		_, ok := actual.(g.Write)
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
})
