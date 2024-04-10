package tests

import (
	"strings"

	g "github.com/beysed/like/internal/grammar"
	. "github.com/beysed/like/internal/tests/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Samples", func() {
	DescribeTable("Sample Incorrect", func(f string, e string) {
		var c = Read(f)

		context, _ := MakeContext()

		_, err := g.Execute("a.like", context, c)
		Expect(err).NotTo(BeNil())
		Expect(strings.HasPrefix(err.Error(), e)).To(BeTrue())
	},
		Entry("samples/len-error", "samples/len-error.like", "['len' accept only single argument]"),
		Entry("samples/error", "samples/error.like", "[syntax error] because of"))

	DescribeTable("Sample Correct", func(f string, e string) {
		var c = Read(f)

		context, result := MakeContext()

		_, err := g.Execute("a.like", context, c)
		Expect(err).To(BeNil())
		Expect(result.String()).To(Equal(e))
	},
		Entry("samples/stringify_array", "samples/stringify_array.like", "[1 2 'a apple' bus]\n"),
		Entry("samples/stringify", "samples/stringify.like", "\n{}\n\n[]\nsss\n{a: 1, b: 2}\n[1 2 'a apple' bus]\n"),
		Entry("samples/exec_arg", "samples/exec_arg.like", "a\n"),
		Entry("samples/conditions", "samples/conditions.like", "1\n\n3\n6\n3\n\n5\n8\nb1f1\nf2\nb\nf3\n"),
		Entry("samples/conditions2", "samples/conditions2.like", "nnnnyyyy"),
		Entry("samples/templ_resource", "samples/templ_resource.like", "resource \"name\" {\r\n}a"),
		Entry("samples/writes", "samples/writes.like", "a\nbbc\nc\n"),
		Entry("samples/loop_write", "samples/loop_write.like", "abc"),
		Entry("samples/len", "samples/len.like", "1133"),
		Entry("samples/global_lambda", "samples/global_lambda.like", "lambda"),
		Entry("samples/global_lambda2", "samples/global_lambda2.like", "docker exec -it flagmap-db a"),
		Entry("samples/docker3", "samples/docker3.like", "/a/docker/exec/-it/b////\n"),
		Entry("samples/docker2", "samples/docker2.like", "a b c\n"),
		Entry("samples/bom", "samples/bom.like", "a"),
		Entry("samples/docker", "samples/docker.like", "docker exec -it abc a\n"),
		Entry("samples/indexes", "samples/indexes.like", "dcccc"),
		Entry("samples/space", "samples/space.like", "Hello World"),
		Entry("samples/interpolation", "samples/interpolation.like", "-asdf-"),
		Entry("samples/yaml_parse", "samples/yaml_parse.like", "a stringtrue3f"),
		Entry("samples/json_parse", "samples/json_parse.like", "1btruef"),
		Entry("samples/env_parse", "samples/env_parse.like", "b1"),
		Entry("samples/yaml_format", "samples/yaml_format.like", "b:\n    d: e\n"),
		Entry("samples/json_format", "samples/json_format.like", "{\"b\":\"a\"}"),
		Entry("samples/env_format", "samples/env_format.like", "b=a\n"),
		Entry("samples/include_test", "samples/include_test.like", "Z"),
		Entry("samples/loop", "samples/loop.like", "a1"),
		Entry("samples/quotes", "samples/quotes.like", "\"a'"),
		Entry("samples/sample", "samples/sample.like", "b\n"),
		Entry("samples/simple_template", "samples/simple_template.like", "aaacbbb\n"),
		Entry("samples/lambda_invoke", "samples/lambda_invoke.like", "_1_2_\n"),
		Entry("samples/lambda_invoke_space", "samples/lambda_invoke_space.like", "_ 1 2 _\n"),
		Entry("samples/func", "samples/func.like", "one\na oops\na oops\n"),
		Entry("samples/execute", "samples/execute.like", "AAA\n"),
		Entry("samples/assigns", "samples/assigns.like", "a1"),
		Entry("samples/template", "samples/template.like", "AfredB\n___\n"))

	DescribeTable("Valid file", func(f string, e string) {
		context, result := MakeContext()

		_, err := g.Execute("a.like", context, []byte(f))
		Expect(err).To(BeNil())
		Expect(result.String()).To(Equal(e))
	},
		Entry("empty", "", ""),
		Entry("line w/o line break", "` a", "a\n"),
		Entry("line with line break", "` a\n", "a\n"))
})
