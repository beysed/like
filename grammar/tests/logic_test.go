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
		Entry("not", "~ ! ''", "T"),
		Entry("not not", "~ ! ! ''", ""),
		Entry("not not paren", "~ !(!'')", ""),
		Entry("not equal", "~ a != b", "T"),
		Entry("equal", "~ a == a", "T"))
})
