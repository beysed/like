package main

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strings"

	g "github.com/beysed/like/internal/grammar"
	c "github.com/beysed/like/internal/grammar/common"
	p "github.com/beysed/like/internal/grammar/parsers"
)

func ExitPrintUsage() {
	d, _ := debug.ReadBuildInfo()
	var version string
	if d != nil {
		version = d.Main.Version
	}

	fmt.Printf("Like, Template Scripting Language %s\n", version)
	fmt.Println("  using file: like input.like [args...]")
	fmt.Println("  using data from stdin: like [args...]")
	os.Exit(1)
}

func isInPipeMode() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		os.Stderr.WriteString("can't stat stdin")
		os.Exit(128)
	}

	return (stat.Mode()&os.ModeCharDevice) == 0 && false
}

func main() {
	var err error
	var input []byte

	if isInPipeMode() {
		input, err = io.ReadAll(os.Stdin)

		if err != nil {
			os.Stderr.WriteString("error reading stdin")
			os.Exit(128)
		}
	}

	args := []string{}
	var fileName string

	if len(input) == 0 {
		if len(os.Args) > 1 {
			fileName = os.Args[1]
		} else {
			ExitPrintUsage()
		}

		if len(os.Args) > 2 {
			args = os.Args[2:]
		}
	} else {
		if len(os.Args) > 1 {
			args = os.Args[1:]
		}
	}

	if len(fileName) == 0 {
		fileName = "stdin"
	} else {
		input, err = os.ReadFile(fileName)
		if err != nil {
			os.Stderr.WriteString("error reading input")
			os.Exit(1)
		}
	}

	globals := c.Store{}
	system := MakeSystemContext()
	cleanup := func(code int) int {
		system.Out.Flush()
		system.Err.Flush()

		return code
	}

	context := g.MakeContext(c.MakeLocals(globals), g.MakeDefaultBuiltIn(), system)
	globals["args"] = make([]any, len(args))
	for i, v := range args {
		globals["args"].([]any)[i] = v
	}

	parser := p.EnvParser{}

	senv := strings.Join(os.Environ(), "\n")
	if _, err = os.Stat(".env"); err == nil {
		local, err := os.ReadFile(".env")
		if err == nil {
			senv = strings.Join([]string{senv, string(local)}, "\n")
		} else {
			system.OutputError("unable to read .env file\n")
		}
	}

	env, err := parser.Parse(senv)
	if err != nil {
		system.OutputError("unable to parse environment\n")
		os.Exit(cleanup(128))
	} else {
		globals["env"] = env
	}

	_, err = g.Execute(fileName, context, input)
	if err != nil {
		system.OutputError(fmt.Sprintf("error:\t%s\n", err.Error()))
		os.Exit(cleanup(128))
	}

	os.Exit(cleanup(0))
}
