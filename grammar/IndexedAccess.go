package grammar

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Value struct {
	value string
}

func (v Value) String() string {
	return v.value
}

type IndexedAccess struct {
	value   Value
	indexes []Value
}

func (a IndexedAccess) String() string {
	index := strings.Join(
		lo.Map(a.indexes, func(s Value, _ int) string {
			return fmt.Sprintf("[%s]", s.String())
		}), "")

	return fmt.Sprintf("%s%s", a.value.String(), index)
}
