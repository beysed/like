package main

import (
	"bufio"
	"os"
	"strings"

	"path"
	"path/filepath"

	c "github.com/beysed/like/internal/grammar/common"
)

type CliSystem struct {
	Cwd string
	Out *bufio.Writer
	Err *bufio.Writer
}

func MakeSystemContext() CliSystem {
	w, e := os.Getwd()
	if e != nil {
		os.Stderr.WriteString("unable to get working directory\n")
		os.Exit(128)
	}

	w = strings.ReplaceAll(w, "\\", "/")
	return CliSystem{
		Cwd: w,
		Out: bufio.NewWriter(os.Stdout),
		Err: bufio.NewWriter(os.Stderr)}
}

func (a CliSystem) ResolvePath(context *c.Context, filePath string) (string, error) {
	isAbs := filepath.IsAbs(filePath)

	var p string
	if isAbs {
		p = strings.ReplaceAll(filePath, "\\", "//")
	} else {
		if strings.HasPrefix(filePath, "./") {
			f, l := context.PathStack.Peek()
			if !f {
				return filePath, c.MakeError("empty stack", nil)
			}

			dir := path.Dir(strings.ReplaceAll(l, "\\", "//"))
			p = path.Join(dir, filePath)
		} else {
			p = path.Join(a.Cwd, filePath)
		}
	}

	_, err := os.Stat(p)
	return p, err
}

func (a CliSystem) OutputText(text string) {
	a.Out.WriteString(text)
}

func (a CliSystem) OutputError(text string) {
	a.Err.WriteString(text)
}
