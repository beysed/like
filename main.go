package main

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strings"

	g "github.com/beysed/like/internal/grammar"
	p "github.com/beysed/like/internal/grammar/parsers"
)

func ExitPrintUsage() {
	d, _ := debug.ReadBuildInfo()

	fmt.Printf("Like, Template Scripting Language, %s\n", d.Main.Version)
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

	return (stat.Mode() & os.ModeCharDevice) == 0
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

	context := g.MakeDefaultContextFor(MakeSystemContext())
	context.Globals["args"] = args
	parser := p.EnvParser{}

	senv := strings.Join(os.Environ(), "\n")
	if _, err = os.Stat(".env"); err == nil {
		local, err := os.ReadFile(".env")
		if err == nil {
			senv = strings.Join([]string{senv, string(local)}, "\n")
		} else {
			os.Stderr.WriteString("unable to read .env file\n")
		}
	}

	env, err := parser.Parse(senv)
	if err != nil {
		os.Stderr.WriteString("unable to parse environment\n")
	} else {
		context.Globals["env"] = env
	}

	_, err = g.Execute(fileName, context, input)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error: %s\n", err.Error()))
		os.Exit(128)
	}
}
