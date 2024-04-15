package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
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

func prepareString(s []byte) (Expressions, error) {
	e, err := Parse("a.like", []byte(s), Entrypoint("string_content"))
	if err != nil {
		return nil, c.MakeError("unable to parse string", err)
	}

	return e.(Expressions), nil
}

func evaluateStringContent(context *c.Context, s []byte) (string, error) {
	e, err := prepareString(s)
	if err != nil {
		return "", err
	}

	result := strings.Builder{}
	for _, v := range e {
		var r any
		r = v
		for {
			r, err = r.(Expression).Evaluate(context)
			if err != nil {
				return "", err
			}
			if _, ok := r.(Expression); ok {
				continue
			}

			for _, v := range flat(r) {
				result.WriteString(fmt.Sprint(v))
			}
			break
		}
	}

	return result.String(), nil
}

func (a ParsedString) Evaluate(context *c.Context) (any, error) {
	s := a.Unquote()
	return evaluateStringContent(context, []byte(s))
}

func MakeParsedString(quote string, body string) ParsedString {
	return ParsedString{Quote: quote, RawBody: body}
}
