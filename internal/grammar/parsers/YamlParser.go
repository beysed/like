package parsers

import (
	c "github.com/beysed/like/internal/grammar/common"
	y "gopkg.in/yaml.v3"
)

type YamlParser struct{}

func (a YamlParser) Parse(input string) (c.Store, error) {
	result := c.Store{}
	err := y.Unmarshal([]byte(input), result)
	if err != nil {
		return nil, c.MakeError("can't parse YAML", err)
	}
	return result, nil
}
