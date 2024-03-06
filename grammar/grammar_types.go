package grammar

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func stringify[T fmt.Stringer](a []T) string {
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

func valueArray(a any) []Value {
	if a == nil {
		return []Value{}
	}

	return lo.Map(a.([]interface{}), func(v interface{}, _ int) Value {
		return v.(Value)
	})
}

type Value struct {
	prefix     string
	identifier string
}

func (t *Value) String() string {
	return fmt.Sprintf("%s%s", t.prefix, t.identifier)
}

type Include struct {
	fileName string
}

type ArrayAccess struct {
	value   Value
	indexes []Value
}

type Assign struct {
	identifier string
	source     []ArrayAccess
}

func (a ArrayAccess) String() string {
	index := strings.Join(
		lo.Map(a.indexes, func(s Value, _ int) string {
			return fmt.Sprintf("[%s]", s.String())
		}), "")

	return fmt.Sprintf("%s%s", a.value.String(), index)
}

// func assignSource(_prefix any, _member any, _access any) (Assign, any) {
// 	// assign_source
// 	prefix, ok := _prefix.(string)
// 	if !ok {
// 		return Assign{}, _prefix
// 	}

// 	var member, ok = _member.(ArrayAccess)
// 	if !ok {
// 		return Assign{}, member
// 	}

// 	p, ok := _access.(Source)
// 	if !ok
// 		return Assign{}, p

// 	return append(Source{ArrayAccessFrom(_p.(string), _m)}, p...), nil
// }
