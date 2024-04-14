package common

import "fmt"

type Store map[string]any
type List []any

type KeyValue[K any, V any] struct {
	Key   K
	Value V
}

type NamedValue KeyValue[string, any]

func (a NamedValue) String() string {
	if len(a.Key) == 0 {
		return Stringify(a.Value)
	}
	return fmt.Sprintf("%s: %s", a.Key, a.Value)
}

func CopyStore(m map[string]interface{}) Store {
	cp := Store{}
	for k, v := range m {
		vm, ok := v.(Store)
		if ok {
			cp[k] = CopyStore(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}
