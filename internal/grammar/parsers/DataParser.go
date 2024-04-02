package parsers

import (
	c "github.com/beysed/like/internal/grammar/common"
)

type DataPraser interface {
	Parse(input string) (c.Store, error)
}
