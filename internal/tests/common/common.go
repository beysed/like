package common

import (
	"os"
	"path/filepath"
	"strings"

	g "github.com/beysed/like/internal/grammar"

	"github.com/bazelbuild/rules_go/go/runfiles"
	k "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Text(input string) []byte {
	return []byte(input)
}

func ParseInupt(input string, entrypoint string) any {
	result, err := g.Parse("a.like", Text(input), g.Entrypoint(entrypoint))
	Expect(err).To(BeNil())
	return result
}

func ParseInuptIncorrect(input string, entrypoint string) (any, error) {
	return g.Parse("a.like", Text(input), g.Entrypoint(entrypoint))
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
	f := File(fileName)
	stat, err := os.Stat(f)
	Expect(err).To(BeNil())

	reader, err := os.Open(f)
	Expect(err).To(BeNil())

	var buf = make([]byte, stat.Size())
	reader.Read(buf)

	return buf
}

func Evaluate(code string) (string, error) {
	context, result := MakeContext()
	wd, _ := os.Getwd()
	_, err := g.Execute(wd, &context, []byte(code))
	return result.String(), err
}
