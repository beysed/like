package tests

import (
	. "like/grammar/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Comments tests", func() {
	DescribeTable("Comments", func(input string, expected string) {
		var actual = ParseInupt(input, "comment")

		Expect(actual).To(BeAssignableToTypeOf(""))
		result := actual.(string)
		Expect(result).To(Equal(expected))
	},
		Entry("empty", "#", "#"),
		Entry("comment", "#asd", "#asd"),
		Entry("comment", "#asd\n", "#asd"))
})
