package common

type Store map[string]any

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
