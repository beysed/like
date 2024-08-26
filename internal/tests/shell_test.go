package tests

import (
	"os"
	"strings"

	p "github.com/beysed/like/internal/grammar/parsers"
)

func GetShell() any {
	var environ = os.Environ()
	var environment, _ = p.GetParser("env").Parse(strings.Join(environ, "\n"))
	return environment["LIKE_SH"]
}
