package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError map[string][]string

func FormatValidationError(err error) map[string][]string {
	var ve validator.ValidationErrors

	out := make(map[string][]string)

	if errors.As(err, &ve) {
		for _, fe := range ve {
			field := strings.ToLower(fe.Field())

			var msg string

			switch fe.Tag() {
			case "required":
				msg = field + " is required"
			case "numeric":
				msg = field + " must be a number"
			case "min":
				msg = field + " must be at least " + fe.Param()
			case "max":
				msg = field + " must be at most " + fe.Param()
			case "email":
				msg = field + " must be a valid email address"
			default:
				msg = field + " is not valid"
			}

			out[field] = append(out[field], msg)
		}
	}
	return out
}
