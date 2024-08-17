package tests

import (
	. "github.com/beysed/like/internal/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grammar Logic", func() {

	It("Error", func() {
		_, err := Evaluate("$error(oops)")

		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("[while evaluating: { $error(oops) }] because of:\n\t[[oops]]"))
	})

	DescribeTable("Not", func(input string, expected string) {
		result, err := Evaluate(input)

		Expect(err).To(BeNil())
		Expect(result.Stdout.String()).To(Equal(expected))
	},
		Entry("eval", "$eval('~ a')", "a"),
		Entry("simple loop", "@ [a b c] ~ -$_k$_v", "-0a-1b-2c"),
		Entry("block loop", "@ [1 2 3] {\n~ a\n~$_v\n}", "a1a2a3"),
		Entry("if yes", "~ T ? yes", "yes"),
		Entry("if no", "~ '' ? yes", ""),
		Entry("ternary1", "~ T ? yes", "yes"),
		Entry("ternary2", "~ T ? yes % no", "yes"),
		Entry("ternary3", "~ '' ? yes % no", "no"),
		Entry("not", "~ ! ''", "T"),
		Entry("not not", "~ ! ! ''", ""),
		Entry("not not paren", "~ !(!'')", ""),
		Entry("not equal", "~ a != b", "T"),
		Entry("equal", "~ a == a", "T"))
})
