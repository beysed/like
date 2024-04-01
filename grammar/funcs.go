package grammar

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func concat[T fmt.Stringer](a []T) string {
	return strings.Join(lo.Map(a, func(s T, _ int) string { return s.String() }), "")
}

func castify[T any](a any, d T) T {
	if a == nil {
		return d
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

func stringify(s any) string {
	return fmt.Sprintf("%s", s)
}

func trim(s string) string {
	return strings.TrimSpace(s)
}
