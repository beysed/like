package common

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func quote(t string) string {
	if !strings.Contains(t, " ") {
		return t
	}

	return fmt.Sprintf("'%s'", strings.ReplaceAll(t, "'", "\\'"))
}

func Qstringify(v any) string {
	switch t := v.(type) {
	case string:
		return quote(t)
	}

	return Stringify(v)
}

func Stringify(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case Store:
		s := []string{}
		for k, v := range t {
			s = append(s, fmt.Sprintf("%s: %s", k, Qstringify(v)))
		}

		return fmt.Sprintf("{%s}", strings.Join(s, ", "))
	case List:
		return fmt.Sprintf("[%s]",
			strings.Join(
				lo.Map(t,
					func(q any, _ int) string {
						return Qstringify(q)
					}), " "))
	case nil:
		return ""
	}

	return fmt.Sprint(v)
}
