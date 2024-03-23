package tests

import (
	g "like/grammar"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Assigns", func() {

	DescribeTable("Identifier", func(input string, indentifier string) {
		var actual = ParseInupt(input, "assign", false)
		Expect(actual).To(BeAssignableToTypeOf(g.Assign{}))
		var assign = actual.(g.Assign)
		Expect(assign.Identifier).To(Equal(indentifier))
		Expect(assign.Value).NotTo(BeNil())
	},
		Entry("literal", "a = a", "a"),
		Entry("quoted string", "b = 'a'", "b"))

	DescribeTable("Evaluate assigns", func(input string, indentifier string, value string) {
		var actual = ParseInupt(input, "assign", false)

		assign, ok := actual.(g.Expression)
		Expect(ok).To(BeTrue())

		var locals = g.Context{}
		var globals = g.Context{}
		var system = TestSystem{}

		var result = assign.Evaluate(system, globals, locals)

		Expect(locals[indentifier]).Should(Equal(value))
		Expect(locals[indentifier]).Should(Equal(result))
	},
		Entry("literal", "a = b", "a", "b"),
		Entry("quoted string", "b = 'a'", "b", "'a'"),
		Entry("a int", "a = 10", "a", "10"),
		Entry("array int", "a = a[0]", "a", "a[0]"),
		Entry("array string index", "a = a['asd']", "a", "a['asd']"),
		Entry("array string string index", "a = 'a'['asd']", "a", "'a'['asd']"),
		Entry("array int int", "a = a[0][1]", "a", "a[0][1]"),
		Entry("array string string", "a = 'asd'['def']", "a", "'asd'['def']"))
})
