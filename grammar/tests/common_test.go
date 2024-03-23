package tests

import (
	. "like/grammar"

	. "github.com/onsi/gomega"
)

func Text(input string) []byte {
	return []byte(input)
}

func ParseInupt(input string, entrypoint string, debug bool) any {
	result, err := Parse("a.like", Text(input), Entrypoint(entrypoint), Debug(debug))
	Expect(err).To(BeNil())
	return result
}

func ParseInuptIncorrect(input string, entrypoint string, debug bool) (any, error) {
	return Parse("a.like", Text(input), Entrypoint(entrypoint), Debug(debug))
}
