package tests

import (
	g "like/grammar"
	. "like/grammar/tests/common"

	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func makeContext() (g.Context, *strings.Builder) {
	system := TestSystem{}
	store := g.Store{}

	return g.Context{
		System:  &system,
		Locals:  store,
		Globals: store,
	}, &system.Result
}

var _ = Describe("Sample", func() {
	DescribeTable("T", func(f string, e string) {
		var c = Read(f)

		context, result := makeContext()

		err := g.Execute(&context, c)
		Expect(err).To(BeNil())
		Expect(result.String()).To(Equal(e))

	},
		Entry("samples/sample.like", "samples/sample.like", "b"),
		Entry("samples/simple_template.like", "samples/simple_template.like", "aaacbbb"))
})
