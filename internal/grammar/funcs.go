package grammar

import (
	"fmt"
	"strings"

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

func unquote(s any) string {
	return s.(ParsedString).Unquote()
}

func listFrom[T any](f any, rest any) []T {
	if rest == nil {
		return []T{f.(T)}
	}

	return append([]T{f.(T)}, arrayify[T](rest)...)
}

func convert(s any) string {
	return strings.Join(arrayify[string](s), "")
}

func trim(s string) string {
	return strings.TrimSpace(s)
}
