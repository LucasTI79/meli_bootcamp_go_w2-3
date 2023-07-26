package web

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field   string
	Msg     string
	Message string
}

func (e *ApiError) CustomError(err error) []ApiError {
	var ve validator.ValidationErrors
	out := make([]ApiError, 0)
	if errors.As(err, &ve) {
		for _, fe := range ve {
			out = append(out, ApiError{
				Field:   fe.Field(),
				Msg:     fe.Tag(),
				Message: fmt.Sprintf("invalid %s field", fe.Field()),
			})
		}
	}
	return out

}
