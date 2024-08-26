package tests

import (
	g "github.com/beysed/like/internal/grammar"
	c "github.com/beysed/like/internal/grammar/common"
	. "github.com/beysed/like/internal/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Precedence", func() {
	DescribeTable("Eval Pipes", func(input string, expected string) {
		sys, res, err := Evaluate(input)
		Expect(err).To(BeNil())
		Expect(sys.Stdout.String()).To(BeEmpty())
		Expect(res).To(Equal(expected))
	},
		Entry("3 in one", "tf=A\n(err = $tf | & fake fmt -) | $fmt", "faked(stdin:A; args:fmt, -)"),
		Entry("bash 3 in one", "tf='echo A'\n(err = $tf | & $_shell -) | $fmt", "A\n"))

	DescribeTable("Debug Pipes", func(input string, expected string) {
		var expr = ParseInupt(input, "file")
		Expect(g.Expressions(expr.([]g.Expression)).Debug()).To(Equal(expected))
	},
		Entry("output pipeout", "~ a > b", ">(~(a) b))"), // need to fix to
		Entry("3 in one", "tf=A\n(err = $tf | & fake fmt -) | $fmt", "=(tf A)|(=(err |($tf &(fake fmt -))) $fmt)"),
		Entry("simple", "& grep | & sort", "|(&(grep) &(sort))"),
		Entry("call simple", "$grep() | $sort()", "|($grep() $sort())"),
		Entry("call with pipe", "& grep ($a | & some) | & sort", "|(&(grep |($a &(some))) &(sort))"))
	DescribeTable("debug function", func(input string, expected string, stdout string) {
		sys, result, err := Evaluate(input)
		Expect(err).To(BeNil())
		Expect(c.Stringify(result)).To(Equal(expected))
		Expect(c.Stringify(sys.Stdout.String())).To(Equal(stdout))
	},
		Entry("call with pipe", "$debug('& grep | & sort')", "|(&(grep) &(sort))", ""),
		Entry("call str", "$debug('\\$a')", "$a", ""),
		Entry("output", "'a' > 'a.tf'", "a", ""))
})
