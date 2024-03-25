package grammar

import . "like/expressions"

type Lambda struct {
	Arguments []string
	Body      []Expression
}
