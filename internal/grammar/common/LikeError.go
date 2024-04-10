package common

import "fmt"

type LikeError struct {
	ErrorText  string
	InnerError error
}

func (a *LikeError) Error() string {
	if a.InnerError == nil {
		return fmt.Sprintf("[%s]", a.ErrorText)
	}

	return fmt.Sprintf("[%s] because of:\n\t%s", a.ErrorText, a.InnerError.Error())
}

func MakeError(err string, inner error) error {
	return &LikeError{ErrorText: err, InnerError: inner}
}
