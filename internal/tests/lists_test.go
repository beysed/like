package tests

import (
	g "github.com/beysed/like/internal/grammar"
	. "github.com/beysed/like/internal/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("List tests", func() {
	DescribeTable("Execs", func(input string, expected string) {
		context, builder := MakeContext()

		_, err := g.Execute("a.like", context, ([]byte)(input))
		Expect(err).To(BeNil())
		Expect(builder.String()).To(Equal(expected))
	},
		Entry("test echo", "~ & echo 1 2 3", "1 2 3\n"),
		Entry("assign list", "a=[1 2 3]\n` $a", "[1 2 3]\n"),
		Entry("exec lists", "a=[1 2 3]\nb=[4 5 6]\n~ & echo $a $b", "1 2 3 4 5 6\n"),
		Entry("exec glue lists", "a=[1 2 3]\nb=[4 5 6]\n~ & echo $(a)$(b)", "1 2 3 4 5 6\n"),
		Entry("exec assigned list", "a=[sss echo AAA]\n~ & echo $a", "sss echo AAA\n"),
		Entry("exec list", "~ & bash [ -c 'echo AAA' ]", "AAA\n"))
})
