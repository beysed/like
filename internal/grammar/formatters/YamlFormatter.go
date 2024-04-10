package formatters

import (
	y "gopkg.in/yaml.v3"
)

type YamlFormatter struct{}

func (a YamlFormatter) Format(input any) (string, error) {
	v, err := y.Marshal(input)
	return string(v), err
}
