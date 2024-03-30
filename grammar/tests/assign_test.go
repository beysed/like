package tests

import (
	g "like/grammar"
	. "like/grammar/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Assigns", func() {

	DescribeTable("Identifier", func(input string, indentifier string, expr string) {
		var actual = ParseInupt(input, "assign", false)
		Expect(actual).To(BeAssignableToTypeOf(g.Assign{}))
		var assign = actual.(g.Assign)
		Expect(assign.Identifier.String()).To(Equal(indentifier))
		Expect(assign.Value.String()).To(Equal(expr))
	},
		Entry("literal", "a = b", "a", "b"),
		Entry("quoted string", "b = 'a'", "b", "'a'"),
		Entry("array int", "a = $a[0]", "a", "$a[0]"),
		Entry("array string index", "a = $a['asd']", "a", "$a['asd']"),
		Entry("array int int", "a = $a[0][1]", "a", "$a[0][1]"),
		Entry("array string string", "a = $'asd'['def']", "a", "$'asd'['def']"))

	DescribeTable("Evaluate assigns", func(input string, indentifier string, value string) {
		var actual = ParseInupt(input, "assign", false)

		assign, ok := actual.(g.Expression)
		Expect(ok).To(BeTrue())

		var system = TestSystem{}

		var globals = g.Store{}
		var context = g.Context{
			Locals:  globals,
			Globals: globals,
			System:  &system,
			//Builtin: g.Store{},
		}

		result, err := assign.Evaluate(&context)
		Expect(err).To(BeNil())
		Expect(context.Locals[indentifier].(g.Store)["value"]).Should(Equal(value))
		Expect(context.Locals[indentifier].(g.Store)["value"]).Should(Equal(result))
	},
		Entry("Evaluate: literal", "a = b", "a", "b"),
		Entry("Evaluate: quoted string", "b = 'a'", "b", "'a'"),
		Entry("Evaluate: a int", "a = 10", "a", "10"))
})
