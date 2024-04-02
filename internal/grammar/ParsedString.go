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

	exps := lo.Map(e.([]any), func(a any, _ int) Expression {
		return a.(Expression)
	})

	return Expressions([]Expression(exps)).Evaluate(context)
}

func MakeParsedString(quote string, body string) ParsedString {
	return ParsedString{Quote: quote, RawBody: body}
}
