package tests

import (
	g "like/grammar"
	. "like/grammar/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {
	DescribeTable("T", func(f string, e string) {
		var c = Read(f)
		system := TestSystem{}
		store := g.Store{}

		context := g.Context{
			System:  &system,
			Locals:  store,
			Globals: store,
		}

		err := g.Execute(&context, c)
		Expect(err).To(BeNil())
		Expect(system.Result.String()).To(Equal(e))

	},
		Entry("samples/sample.like", "samples/sample.like", "b"),
		Entry("samples/simple_template.like", "samples/simple_template.like", "aaacbbb"))
})
