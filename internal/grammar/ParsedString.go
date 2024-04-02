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

func (a ParsedString) Evaluate(context *c.Context) (any, error) {
	// todo: parse references
	return a.Unquote(), nil
}

func MakeParsedString(quote string, body string) ParsedString {
	return ParsedString{Quote: quote, RawBody: body}
}
