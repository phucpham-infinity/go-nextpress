package common_response

import "github.com/gofiber/fiber/v2"

type successResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Padding interface{} `json:"padding,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
	Ctx     *fiber.Ctx  `json:"-"`
}

func NewSuccessResponse(c *fiber.Ctx) *successResponse {
	return &successResponse{
		Ctx:     c,
		Data:    nil,
		Padding: nil,
		Filters: nil,
		Status:  200,
		Message: "ok",
	}
}

func (r *successResponse) WithData(data interface{}) *successResponse {
	r.Data = data
	return r
}

func (r *successResponse) WithPadding(padding interface{}) *successResponse {
	r.Padding = padding
	return r
}

func (r *successResponse) WithFilters(filters interface{}) *successResponse {
	r.Filters = filters
	return r
}

func (r *successResponse) SendJSON() error {
	return r.Ctx.Status(200).JSON(r)
}
