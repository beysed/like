package main

import (
	"fmt"
	"os"
	"strings"

	"path"
	"path/filepath"

	c "github.com/beysed/like/internal/grammar/common"
)

type CliSystem struct {
	Cwd string
}

func MakeSystemContext() c.System {
	w, e := os.Getwd()
	if e != nil {
		os.Stderr.WriteString("unable to get working directory\n")
		os.Exit(128)
	}

	return CliSystem{
		Cwd: w,
	}
}

func (a CliSystem) ResolvePath(context *c.Context, filePath string) (string, error) {
	isAbs := filepath.IsAbs(filePath)

	var p string
	if isAbs {
		p = filePath
	} else {
		if strings.HasPrefix(filePath, "./") {
			f, l := context.PathStack.Peek()
			if !f {
				return filePath, c.MakeError("empty stack", nil)
			}

			p = path.Join(l, filePath)
		} else {
			p = path.Join(a.Cwd, filePath)
		}
	}

	_, err := os.Stat(p)
	return p, err
}

func (c CliSystem) OutputText(text string) {
	// odd: cut last line, need to investigate
	fmt.Fprintf(os.Stdout, "%s", text)
}

func (c CliSystem) OutputError(text string) {
	fmt.Fprintf(os.Stderr, "%s", text)
}
