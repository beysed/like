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
		_, locals := context.Locals.Peek()
		locals.Store["_shell"] = getShell()

		_, err := g.Execute("a.like", context, ([]byte)(input))
		Expect(err).To(BeNil())
		Expect(builder.String()).To(Equal(expected))
	},
		Entry("test echo", "~ & $_shell -c 'echo 1 2 3'", "1 2 3\n"),
		Entry("assign list", "a=[1 2 3]\n~ $a", "[1 2 3]"),
		Entry("exec lists", "a=[1 2 3]\nb=[4 5 6]\n~ & $_shell -c 'echo $a $b'", "123 456\n"),
		Entry("exec glue lists", "a=[1 2 3]\nb=[4 5 6]\n~ & $_shell -c 'echo $(a)$(b)'", "123456\n"),
		Entry("exec assigned list", "a=[sss echo AAA]\n~ & $_shell -c 'echo $a'", "sssechoAAA\n"),
		Entry("exec empty list", "a=[]\n~ & $_shell -c 'echo $a'", "\n"),
		Entry("exec direct list", "~ & $_shell [ -c 'echo AAA' ]", "AAA\n"))
})
