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
		_, e := Parse("a.like", []byte(""))
		Expect(e).To(BeNil())
	})

	DescribeTable("Parses_Comments",
		func(a TestEntry) {
			actual, err := Parse("a.like", text(a.input), Entrypoint("comment"))
			Expect(err).To(BeNil())
			Expect(actual.(string)).To(Equal(a.expected))
		},

		Entry("empty comment", NewTestEntry("#", "#")),
		Entry("ordinal comment", NewTestEntry("#asd", "#asd")),
		Entry("comment after directive", NewTestEntry("// include file #asd", "#asd")))
})

type TestEntry struct {
	input    string
	expected string
}

func NewTestEntry(input, expected string) TestEntry {
	return TestEntry{input, expected}
}
