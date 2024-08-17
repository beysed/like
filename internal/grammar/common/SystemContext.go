package common

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/beysed/shell/execute"
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

func (a CliSystem) ResolvePath(context *Context, filePath string) (string, error) {
	isAbs := filepath.IsAbs(filePath)

	var p string
	if isAbs {
		p = strings.ReplaceAll(filePath, "\\", "//")
	} else {
		if strings.HasPrefix(filePath, "./") {
			f, l := context.PathStack.Peek()
			if !f {
				return filePath, MakeError("empty stack", nil)
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

func (a CliSystem) Invoke(executable string, args []string, stdin string) (string, string, string, error) {
	command := execute.MakeCommand(executable, args...)
	execution, err := execute.Execute(command)

	if err != nil {
		return "", "", "", err
	}

	if stdin != "" {
		execution.Stdin <- []byte(stdin)
	}

	close(execution.Stdin)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var stdmixed bytes.Buffer

	run := true
	var exitError error
	for run {
		select {
		case out := <-execution.Stderr:
			stderr.Write(out)
			stdmixed.Write(out)
		case out := <-execution.Stdout:
			stdout.Write(out)
			stdmixed.Write(out)
		case exitError = <-execution.Exit:
			run = false
		case <-time.After(time.Second * 10):
			execution.Kill()
			run = false
		}
	}

	return stdout.String(), stderr.String(), stdmixed.String(), exitError
}
