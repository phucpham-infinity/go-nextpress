package common_response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Status  int        `json:"status"`
	Message string     `json:"message,omitempty"`
	Log     string     `json:"log,omitempty"`
	Key     string     `json:"error_key,omitempty"`
	Ctx     *fiber.Ctx `json:"-"`
	RootErr error      `json:"_"`
}

func NewAppError(c *fiber.Ctx) *AppError {
	return &AppError{
		Ctx:     c,
		Status:  http.StatusBadRequest,
		Message: "",
		Log:     "",
		Key:     "",
		RootErr: nil,
	}
}

func (r *AppError) WithKey(key string) *AppError {
	r.Key = key
	return r
}

func (r *AppError) WithLog(log string) *AppError {
	r.Log = log
	return r
}

func (r *AppError) WithStatus(status int) *AppError {
	r.Status = status
	return r
}

func (r *AppError) WithMessage(message string) *AppError {
	r.Message = message
	return r
}

func (r *AppError) WithRootError(root error) *AppError {
	r.RootErr = root
	return r
}

func (r *AppError) IsInternalServerError(root error) *AppError {
	r.Status = http.StatusInternalServerError
	r.RootErr = root
	r.Message = http.StatusText(http.StatusInternalServerError)
	return r
}

func (r *AppError) IsNotFound(root error) *AppError {
	r.Status = http.StatusNotFound
	r.RootErr = root
	r.Message = http.StatusText(http.StatusNotFound)
	return r
}

func (r *AppError) IsUnauthorized(root error) *AppError {
	r.Status = http.StatusUnauthorized
	r.RootErr = root
	r.Message = http.StatusText(http.StatusUnauthorized)
	return r
}

func (r *AppError) IsForbidden(root error) *AppError {
	r.Status = http.StatusForbidden
	r.RootErr = root
	r.Message = http.StatusText(http.StatusForbidden)
	return r
}

func (r *AppError) IsConflict(root error) *AppError {
	r.Status = http.StatusConflict
	r.RootErr = root
	r.Message = http.StatusText(http.StatusConflict)
	return r
}

func (r *AppError) IsBadRequest(root error) *AppError {
	r.Status = http.StatusBadRequest
	r.RootErr = root
	r.Message = http.StatusText(http.StatusBadRequest)
	return r
}

func (r *AppError) RootError() error {
	if err, ok := r.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return r.RootErr
}

func (r *AppError) Error() string {
	return r.RootErr.Error()
}

func (r *AppError) SendJSON() error {
	return r.Ctx.Status(r.Status).JSON(r)
}
