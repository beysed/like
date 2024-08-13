package grammar

import (
	"fmt"
	"regexp"
	"strconv"
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
	var ref string
	var v any
	var err error

	if s, ok := a.Reference.(Literal); ok {
		ref = s.String()
	} else if s, ok := a.Reference.(Reference); ok {
		v, err = s.Evaluate(context)
		if err != nil {
			return s, err
		}

		if ref, ok = v.(string); ok {
		}
	} else if s, ok := a.Reference.(ParsedString); ok {
		r, err := s.Evaluate(context)
		if err != nil {
			return ref, err
		}
		ref = r.(string)
	} else {
		v, err = a.Reference.Evaluate(context)
		if err != nil {
			return a.Reference, err
		}
	}

	if ref != "" {
		store := findStore(context, ref)
		if store != nil {
			return &StoreReference{
				Store:     store,
				Reference: ref}, nil
		}

		_, locals := context.Locals.Peek()
		store = locals.Store
		return &StoreReference{
			Store:     store,
			Reference: ref}, nil
	}

	s, ok := v.(*StoreReference)
	if !ok {
		if a.Index == nil {
			return v, nil
		}

		if s, ok := v.(c.Store); ok {
			if i, ok := a.Index.(StoreAccess); ok {
				local := MakeContext(c.MakeLocals(s), context.BuiltIn, context.System)
				n := StoreAccess{Reference: i}

				return n.Evaluate(local)
			}
		}

		return v, c.MakeError("don't know, fix me", nil)
	}

	if a.Index == nil {
		return s, nil
	}

	val := s.Get()
	store, ok := val.(c.Store)
	if !ok {
		if val != nil {
			if arr, ok := val.(c.List); ok {
				if a.Index == nil {
					return arr, nil
				}
				idx, err := a.Index.Evaluate(context)
				if err != nil {
					return a.Index, err
				}
				idxN, err := strconv.Atoi(c.Stringify(idx))
				if err != nil {
					return a.Index, c.MakeError("wrong index value", nil)
				}
				return arr[idxN], nil
			}
		}

		store = c.Store{}
		s.Set(store)
	}

	if i, ok := a.Index.(StoreAccess); ok {
		local := MakeContext(c.MakeLocals(store), context.BuiltIn, context.System)
		n := StoreAccess{Reference: i}

		return n.Evaluate(local)
	}

	res, err := a.Index.Evaluate(context)
	if err != nil {
		return a.Index, err
	}

	if expr, ok := res.(Expression); ok {
		res, err = expr.Evaluate(context)
		if err != nil {
			return expr, err
		}

	}

	idx := c.Stringify(unwrap(res))
	return store[idx], nil
}

func unwrap(t any) any {
	if a, ok := t.([]any); ok {
		l := len(a)
		if l == 0 {
			return nil
		} else if l == 1 {
			return unwrap(a[0])
		}
	}

	return t
}
