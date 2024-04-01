package grammar

type LikeError struct {
	ErrorText string
}

func (a *LikeError) Error() string {
	return a.ErrorText
}

func MakeError(err string) error {
	return &LikeError{ErrorText: err}
}
