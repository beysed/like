package grammar

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func ShowError(t *testing.T, err error) {
	if err == nil {
		return
	}

	if errList, ok := err.(*errList); ok {
		t.Log(errList.String())
	} else {
		t.Log(err)
	}
}

func text(s string) []byte {
	return []byte(s)
}

func (lst *errList) String() string {
	if lst == nil {
		return ""
	}

	var builder strings.Builder
	for _, e := range *lst {
		var perr *parserError
		indent := 0
		errors.As(e, &perr)
		for perr != nil {
			builder.WriteString(fmt.Sprintf("%s | %s\n", strings.Repeat(" ", indent), perr.Error()))
			if perr.Inner != nil {
				if errors.As(perr.Inner, &perr) {
					indent += 2
				} else {
					builder.WriteString(perr.Inner.Error())
					builder.WriteString("\n")
					perr = nil
				}
			}
		}
	}

	return builder.String()
}

func TestGrammar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grammar")
}

var _ = Describe("Grammar", func() {
	It("Parses: empty ", func() {
		_, e := Parse("a.like", text(""))
		Expect(e).To(BeNil())
	})

	var parse = func(input string, entrypoint string, debug bool) any {
		result, err := Parse("a.like", text(input), Entrypoint(entrypoint), Debug(debug))
		Expect(err).To(BeNil())
		return result
	}

	DescribeTable("Parses: correct indentifiers", func(input string, expected string) {
		actual := parse(input, "identifier", false)

		Expect(actual).To(BeAssignableToTypeOf(""))
		result := actual.(string)

		Expect(result).To(Equal(expected))
	},
		Entry("single", "a", "a"),
		Entry("multuple", "Aa", "Aa"))

	DescribeTable("Parses: include directive", func(input string, fn string) {
		var actual = parse(input, "directive", false)

		Expect(actual).To(BeAssignableToTypeOf(Include{}))
		result := actual.(Include)

		Expect(result.fileName).To(Equal(fn))
	},
		Entry("unquoted", "// include asd", "asd"),
		Entry("unquoted comment", "// include asd#comment", "asd"),
		Entry("unquoted comment space", "// include asd #comment", "asd"),
		Entry("single quoted", "// include 'asd'", "'asd'"),
		Entry("double quoted", "// include \"asd\"", "\"asd\""))

	DescribeTable("Parses: value / memeber_list", func(input string) {
		var actual = parse(input, "memeber_list", false)

		Expect(actual).To(BeAssignableToTypeOf(IndexedAccess{}))
		result := actual.(IndexedAccess)
		Expect(result.String()).To(Equal(input))
	},
		Entry("simple no index", "a"),
		Entry("double index", "a[0][1]"),
		Entry("single index", "a[0]"))

	DescribeTable("Parses: value", func(input string) {
		var actual = parse(input, "value", false)

		Expect(actual).To(BeAssignableToTypeOf(Value{}))
		result := actual.(Value)
		Expect(result.String()).To(Equal(input))
	},
		Entry("simple value", "a"),
		Entry("prefixed value", "$a"))

	DescribeTable("Parses: comment", func(input string, expected string) {
		var actual = parse(input, "comment", false)

		Expect(actual).To(BeAssignableToTypeOf(""))
		result := actual.(string)
		Expect(result).To(Equal(expected))
	},
		Entry("empty", "#", "#"),
		Entry("comment", "#asd", "#asd"))

	DescribeTable("Parses: assign", func(input string, indentifier string, source string) {
		var actual = parse(input, "assign", false)

		Expect(actual).To(BeAssignableToTypeOf(Assign{}))
		result := actual.(Assign)
		Expect(result.identifier).To(Equal(indentifier))
		Expect(concat(result.source)).To(Equal(source))
	},
		Entry("simple a", "a = a", "a", "a"),
		Entry("simple b", "a = b", "a", "b"),
		Entry("simple b int", "a = 10", "a", "10"),
		//Entry("dd", "exec=docker(exec -t $s.container)",
		Entry("array int", "a = a[0]", "a", "a[0]"),
		Entry("array string index", "a = a['asd']", "a", "a['asd']"),
		Entry("array string string index", "a = 'a'['asd']", "a", "'a'['asd']"),
		Entry("array int int", "a = a[0][1]", "a", "a[0][1]"),
		Entry("array string string", "a = 'asd'['def']", "a", "'asd'['def']"))
})
