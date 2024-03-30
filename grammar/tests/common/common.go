package common

import (
	. "like/grammar"
	"os"
	"path/filepath"
	"strings"

	"github.com/bazelbuild/rules_go/go/runfiles"
	k "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Text(input string) []byte {
	return []byte(input)
}

func ParseInupt(input string, entrypoint string) any {
	result, err := Parse("a.like", Text(input), Entrypoint(entrypoint))
	Expect(err).To(BeNil())
	return result
}

func ParseInuptIncorrect(input string, entrypoint string) (any, error) {
	return Parse("a.like", Text(input), Entrypoint(entrypoint))
}

func Log(format string, args ...any) {
	k.GinkgoWriter.Printf(format+"\n", args...)
}

func File(fileName string) string {
	var current = k.CurrentSpecReport().LeafNodeLocation.FileName
	var result = filepath.Join(filepath.Dir(current), fileName)

	var f, e = runfiles.Rlocation(
		strings.ReplaceAll(
			filepath.Join("_main", result), "\\", "/"))

	if e != nil {
		return result
	} else {
		return f
	}
}

func Read(fileName string) []byte {
	stat, err := os.Stat(File(fileName))
	Expect(err).To(BeNil())

	reader, err := os.Open(File(fileName))
	Expect(err).To(BeNil())

	var buf = make([]byte, stat.Size())
	reader.Read(buf)

	return buf
}
