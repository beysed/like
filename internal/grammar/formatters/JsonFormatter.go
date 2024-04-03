package formatters

import (
	j "encoding/json"
)

type JsonFormatter struct{}

func (a JsonFormatter) Format(input any) (string, error) {
	r, err := j.Marshal(input)
	return string(r), err
}
