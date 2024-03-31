package grammar

import (
	"fmt"
	"regexp"
	"strings"
)

// "strings"

// "github.com/samber/lo"

type StoreAccess struct {
	Reference Expression
	Index     Expression
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

func (a StoreAccess) Evaluate(context *Context) (any, error) {
	if literal, ok := a.Reference.(Literal); ok {
		var v = literal.String()
		if context.Locals[v] != nil {
			return context.Locals[v], nil
		} else if context.Globals[v] != nil {
			return context.Globals[v], nil
		} else {
			context.Locals[v] = Store{}
			return context.Locals[v], nil
		}
	}

	if store, ok := a.Reference.(StoreAccess); ok {
		v, e := store.Evaluate(context)
		if e != nil {
			return v, e
		}

		s := v.(Store)
		if a.Index == nil {
			return s, nil
		}

		local := &Context{
			Globals: s,
			Locals:  s,
			System:  context.System,
		}

		n := StoreAccess{
			Reference: a.Index}
		return n.Evaluate(local)
	}

	return nil, nil
}
