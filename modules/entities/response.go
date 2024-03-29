package entities

import "github.com/gofiber/fiber/v2"

func NewResponse(c *fiber.Ctx) IResponse {
	return &Response{
		Context: c,
	}
}

type IResponse interface {
	Success(code int, data any) IResponse
	Error(code int, tractId, msg string) IResponse
	Res() error
}

type Response struct {
	StatusCode int
	Data       any
	ErrorRes   *ErrorResponse
	Context    *fiber.Ctx
	IsError    bool
}

func (r *Response) Success(code int, data any) IResponse {
	r.StatusCode = code
	r.Data = data
	return r
}
func (r *Response) Error(code int, tractId, msg string) IResponse {
	r.StatusCode = code
	r.ErrorRes = &ErrorResponse{
		TraceId: tractId,
		Msg:     msg,
	}
	r.IsError = true
	return r
}
func (r *Response) Res() error {
	return r.Context.Status(r.StatusCode).JSON(func() any {
		if r.IsError {
			return &r.ErrorRes
		}
		return &r.Data
	}())
}

type ErrorResponse struct {
	TraceId string `json:"trace_id"`
	Msg     string `json:"message"`
}
