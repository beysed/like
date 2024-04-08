package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

type ParsedString struct {
	Quote   string
	RawBody string
}

func (a ParsedString) String() string {
	return fmt.Sprintf("%s%s%s", a.Quote, a.RawBody, a.Quote)
}

func (a ParsedString) Unquote() string {
	return strings.ReplaceAll(a.RawBody, "\\"+a.Quote, a.Quote)
}

func (a ParsedString) Evaluate(context *c.Context) (any, error) {
	s := a.Unquote()
	e, err := Parse("a.like", []byte(s), Entrypoint("string_content"))
	if err != nil {
		return s, c.MakeError("unable to parse string", err)
	}

	result := strings.Builder{}
	for _, a := range e.([]any) {
		v, ok := a.(Expression)
		if !ok {
			return a, c.MakeError("not an expression", nil)
		}

		var r any
		r = v
		for {
			r, err = r.(Expression).Evaluate(context)
			if err != nil {
				return v, err
			}
			if _, ok := r.(Expression); ok {
				continue
			}

			if v, ok := r.([]any); ok {
				result.WriteString(
					strings.Join(
						lo.Map(v,
							func(v any, _ int) string {
								return fmt.Sprint(v)
							}), " "))

			} else {
				result.WriteString(fmt.Sprint(r))
			}
			break
		}
	}

	return result.String(), nil
}

func MakeParsedString(quote string, body string) ParsedString {
	return ParsedString{Quote: quote, RawBody: body}
}
