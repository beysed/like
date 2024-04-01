package tests

import (
	. "like/grammar/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grammar Logic", func() {
	DescribeTable("Not", func(input string, expected string) {
		result, err := Evaluate(input)

		Expect(err).To(BeNil())
		Expect(result).To(Equal(expected))
	},
		Entry("simple loop", "@ [1 2 3] ~ a$_", "a1a2a3"),
		Entry("block loop", "@ [1 2 3] {\n~ a\n~$_\n}", "a1a2a3"),
		Entry("if yes", "~ ? T yes", "yes"),
		Entry("not", "~ ! ''", "T"),
		Entry("not not", "~ ! ! ''", ""),
		Entry("not not paren", "~ !(!'')", ""),
		Entry("not equal", "~ a != b", "T"),
		Entry("equal", "~ a == a", "T"))
})
