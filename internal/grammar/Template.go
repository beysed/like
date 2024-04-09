package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Template struct {
	Store     StoreAccess
	Arguments IdentifierList
	Value     Expression
	Text      string
}

func (a Template) String() string {
	return fmt.Sprintf("`` %s(%s)\n%s\n``", a.Store.String(), a.Arguments.String(), a.Text)
}

func (a Template) Evaluate(context *c.Context) (any, error) {
	if a.Value == nil {
		// trim orphan \r
		arr := []byte(a.Text)
		if arr[len(arr)-1] < 32 {
			arr = arr[:len(arr)-1]
		}

		str, err := prepareString(arr)
		if err != nil {
			return nil, err
		}

		a.Value = Lambda{
			Arguments: a.Arguments,
			Body:      str,
		}
	}

	assign := Assign{
		Store: a.Store,
		Value: a.Value}

	return assign.Evaluate(context)
}

func MakeTemplate(store StoreAccess, args IdentifierList, body string) Template {
	return Template{
		Store:     store,
		Arguments: args,
		Text:      body,
	}
}
