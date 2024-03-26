package grammar

type Environment map[string]string

func Execute(context Context, code []byte) (string, error) {
	_, err := Parse("a.like", code)

	if err != nil {
		return "", err
	}
	return "", nil
	// expressions = arrayify[Expression](result)

	// for e := range expressions {
	// }
}
