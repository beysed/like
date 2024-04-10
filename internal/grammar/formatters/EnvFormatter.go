package formatters

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
)

type EnvFormatter struct{}

func (a EnvFormatter) Format(input any) (string, error) {

	store, ok := input.(c.Store)
	if !ok {
		return "", c.MakeError("input should be an object", nil)
	}

	result := strings.Builder{}
	for k, v := range store {
		result.WriteString(fmt.Sprintf("%s=%s\n", k, v))
	}

	return result.String(), nil
}
