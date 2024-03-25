package grammar

import . "like/expressions"

type Call struct {
	Indentifier string
	Arguments   []Expression
}
