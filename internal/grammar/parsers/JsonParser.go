package parsers

import (
	c "github.com/beysed/like/internal/grammar/common"
	j "github.com/valyala/fastjson"
)

type JsonParser struct{}

func wrap(v *j.Value) any {
	switch v.Type() {
	case j.TypeObject:
		m := c.Store{}
		v.GetObject().Visit(func(k []byte, v *j.Value) {
			m[string(k)] = wrap(v)
		})
		return m
	case j.TypeNumber:
		return v.GetInt()
	case j.TypeArray:
		var a []any
		for _, v := range v.GetArray() {
			a = append(a, wrap(v))
		}
		return a
	case j.TypeString:
		return string(v.GetStringBytes())
	case j.TypeTrue:
		return "true"
	case j.TypeFalse:
		return "false"
	default:
		return nil
	}
}

func (a JsonParser) Parse(input string) (c.Store, error) {
	v, err := j.Parse(input)

	if err != nil {
		return nil, c.MakeError("can't parse JSON", err)
	}

	w, ok := wrap(v).(c.Store)
	if !ok {
		return nil, c.MakeError("JSON should be an object", nil)
	}

	return w, nil
}
