package grammar

type Ref[T any] interface {
	Get() *T
	Set(v *T) *T
}

type PointerValue[T any] struct {
	value *T
}

func (a *PointerValue[T]) Get() *T {
	return a.value
}

func (a *PointerValue[T]) Set(v *T) *T {
	*(a.value) = *v
	return a.value
}

func MakeRef[T any](v *T) *PointerValue[T] {
	return &PointerValue[T]{
		value: v}
}
