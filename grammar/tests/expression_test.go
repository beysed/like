package tests

import (
	g "like/grammar"

	. "like/grammar/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grammar", func() {
	DescribeTable("Parses: list expressions with reference", func(input string) {
		var actual = ParseInupt(input, "write", false)

		result, ok := actual.(g.Write)
		Expect(ok).To(BeTrue())
		Log(result.String())

		//Expect(result.String()).To(Equal(expect))
	},
		Entry("a", "` aaa$b"),
		Entry("b", "` aaa$(b)ccc"))
})
