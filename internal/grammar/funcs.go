package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

func concat[T fmt.Stringer](a []T) string {
	return strings.Join(lo.Map(a, func(s T, _ int) string { return s.String() }), "")
}

func castify[T any](a any) T {
	if a == nil {
		var t T
		return t
	}
	return a.(T)
}

func arrayify[T any](a any) []T {
	if a == nil {
		return []T{}
	}

	return lo.Map(a.([]interface{}), func(e interface{}, _ int) T { return e.(T) })
}

func listFrom[T any](f any, rest any) []T {
	if rest == nil {
		return []T{f.(T)}
	}

	return append([]T{f.(T)}, arrayify[T](rest)...)
}

func unquote(s any) string {
	return s.(ParsedString).Unquote()
}

func convert(s any) string {
	return strings.Join(arrayify[string](s), "")
}

func quote(t string) string {
	if !strings.Contains(t, " ") {
		return t
	}

	return fmt.Sprintf("'%s'", strings.ReplaceAll(t, "'", "\\'"))
}

func qstringify(v any) string {
	switch t := v.(type) {
	case string:
		return quote(t)
	}

	return stringify(v)
}

func stringify(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case c.Store:
		s := []string{}
		for k, v := range t {
			s = append(s, fmt.Sprintf("%s: %s", k, qstringify(v)))
		}

		return fmt.Sprintf("{%s}", strings.Join(s, ", "))
	case List:
		return fmt.Sprintf("[%s]",
			strings.Join(
				lo.Map(t,
					func(q any, _ int) string {
						return qstringify(q)
					}), " "))
	case nil:
		return ""
	}

	// default
	return fmt.Sprint(v)
}

func trim(s string) string {
	return strings.TrimSpace(s)
}
