package tests

import (
	g "github.com/beysed/like/internal/grammar"
	c "github.com/beysed/like/internal/grammar/common"
	. "github.com/beysed/like/internal/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Precedence", func() {
	DescribeTable("Pipes", func(input string, expected string) {
		var expr = ParseInupt(input, "file")
		Expect(g.Expressions(expr.([]g.Expression)).Debug()).To(Equal(expected))
	},
		Entry("simple", "& grep | & sort", "|(&(grep) &(sort))"),
		Entry("call simple", "$grep() | $sort()", "|($grep() $sort())"),
		Entry("call with pipe", "& grep ($a | & some) | & sort", "|(&(grep |($a &(some))) &(sort))"))

	It("debug function", func() {
		_, result, err := Evaluate("$debug('& grep | & sort')")
		Expect(err).To(BeNil())
		Expect(c.Stringify(result)).To(Equal("|(&(grep) &(sort))"))
	})
})
