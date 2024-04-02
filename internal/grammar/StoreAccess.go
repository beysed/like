package grammar

import (
	"fmt"
	"regexp"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
)

type StoreAccess struct {
	Reference Expression
	Index     Expression
}

type StoreReference struct {
	Store     c.Store
	Reference string
}

type Value interface {
	Get() any
	Set(v any) any
}

func (a StoreReference) Get() any {
	return a.Store[a.Reference]
}

func (a StoreReference) Set(v any) any {
	a.Store[a.Reference] = v
	return v
}

func (a StoreAccess) String() string {
	ref := a.Reference.String()
	if a.Index == nil {
		return ref
	}

	ind := a.Index.String()

	var f string

	if strings.Contains(ind, " ") {
		f = "%s['%s']"
	} else if ok, _ := regexp.Match("^[0-9'\"]", []byte(ind)); ok {
		f = "%s[%s]"
	} else {
		f = "%s.%s"
	}

	return fmt.Sprintf(f, ref, ind)
}

func (a StoreAccess) Evaluate(context *c.Context) (any, error) {
	if literal, ok := a.Reference.(Literal); ok {
		var v = literal.String()
		if context.Locals[v] != nil {
			return &StoreReference{
				Store:     context.Locals,
				Reference: v}, nil
		} else if context.Globals[v] != nil {
			return &StoreReference{
				Store:     context.Globals,
				Reference: v}, nil
		} else {
			context.Locals[v] = c.Store{}
			return &StoreReference{
				Store:     context.Locals,
				Reference: v}, nil
		}
	}

	if st, ok := a.Reference.(StoreAccess); ok {
		v, e := st.Evaluate(context)
		if e != nil {
			return v, e
		}

		s := v.(*StoreReference)
		if a.Index == nil {
			return s, nil
		}

		store := s.Get().(c.Store)
		local := MakeContext(store, store, context.BuiltIn, context.System)

		n := StoreAccess{
			Reference: a.Index}

		return n.Evaluate(&local)
	}

	return nil, nil
}
