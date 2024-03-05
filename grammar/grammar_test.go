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
	RunSpecs(t, "Grammar Suite")
}

var _ = Describe("Grammar", func() {
	It("Parses_Empty ", func() {
		_, e := Parse("a.like", text(""))
		Expect(e).To(BeNil())
	})

	var parseInput = func(input string, rule string, expected string) {
		actual, err := Parse("a.like", text(input), Entrypoint(rule))
		Expect(err).To(BeNil())
		Expect(actual.(string)).To(Equal(expected))
	}

	var parse = func(rule string) func(TestEntry) {
		return func(a TestEntry) {
			parseInput(a.input, rule, a.expected)
		}
	}

	var same = func(s string) TableEntry {
		return Entry(s, NewTestEntry(s, s))
	}

	It("Parses q ", func() {
		parseInput("asd.my", "unquotedString", "asd.my")
	})

	It("Parses include q ", func() {
		parseInput("asd.my", "include", "asd.my")
	})

	DescribeTable("Parses comment", parse("comment"),
		same("#"),
		same("#asd"))

	DescribeTable("Parses unquotedString", parse("unquotedString"),
		same("asd.my"))

	DescribeTable("Parses string param", parse("stringParam"),
		same("asd.my"),
		same("'asd.my'"),
		same("\"asd.my\""))

	DescribeTable("Parses_Line",
		parse("line"),
		Entry("include: file param", NewTestEntry("include a", "include a")),
		Entry("include: quoted string", NewTestEntry("include 'a'", "include 'a'")),
		Entry("include: file param comment", NewTestEntry("include a", "#")),
		Entry("include: quoted string comment", NewTestEntry("include a", "#")),
		Entry("include: quoted string space comment", NewTestEntry("include a", "#")))
})

type TestEntry struct {
	input    string
	expected string
}

func NewTestEntry(input, expected string) TestEntry {
	return TestEntry{input, expected}
}
