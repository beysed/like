package tests

import (
	. "github.com/beysed/like/internal/tests/common"

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
		Entry("empty", "# ", "# "),
		Entry("comment", "# asd", "# asd"),
		Entry("comment wo spc", "#asd", "#asd"),
		Entry("comment next line", "# asd\n", "# asd"))
})
