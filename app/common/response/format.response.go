package common_response

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

func FormatValidationError(err error) string {
	var errorMessages []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			var errorMessage string
			switch fieldError.Tag() {
			case "required":
				errorMessage = fmt.Sprintf("%s is required", fieldError.Field())
			case "email":
				errorMessage = fmt.Sprintf("%s must be a valid email address", fieldError.Field())
			case "min":
				errorMessage = fmt.Sprintf("%s must be at least %s characters long", fieldError.Field(), fieldError.Param())
			case "max":
				errorMessage = fmt.Sprintf("%s must be at most %s characters long", fieldError.Field(), fieldError.Param())
			default:
				errorMessage = fmt.Sprintf("%s is invalid", fieldError.Field())
			}

			errorMessages = append(errorMessages, errorMessage)
		}
	}

	return strings.Join(errorMessages, ", ")
}

func FormatDatabaseError(err error) string {
	if errors.Is(err, sql.ErrNoRows) {
		return "No record found"
	}

	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation":
			return fmt.Sprintf("Duplicate value violates unique constraint: %s", pqErr.Constraint)
		case "foreign_key_violation":
			return fmt.Sprintf("Foreign key violation: %s", pqErr.Constraint)
		default:
			return fmt.Sprintf("Database error: %s", pqErr.Message)
		}
	}

	return "Database error"
}
