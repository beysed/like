package parsers

import (
	"strings"

	"bufio"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

type EnvParser struct{}

func getLines(input string) []string {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanLines)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func (a EnvParser) Parse(input string) (c.Store, error) {
	lines := getLines(input)
	entries := lo.Map(lines, func(s string, _ int) []string {
		return lo.Map(strings.Split(s, "="), func(s string, _ int) string {
			return strings.TrimSpace(s)
		})
	})

	result := c.Store{}

	for _, v := range entries {
		l := len(v)
		if l == 0 {
			continue
		}

		result[v[0]] = strings.Join(v[1:], "=")
	}

	return result, nil
}
