package helpers

import (
	"github.com/kataras/iris/v12"
)

type Response struct {
	iris iris.Context
}

func NewResponse(c iris.Context) *Response {
	return &Response{
		iris: c,
	}
}

func (r *Response) Error(message string) {
	r.iris.JSON(map[string]string{
		"message": message,
		"status":  "error",
	})
}

func (r *Response) ErrorWithData(message string, data interface{}) {
	r.iris.JSON(map[string]interface{}{
		"message": message,
		"status":  "error",
		"data":    data,
	})
}

func (r *Response) Success(data interface{}) {
	r.iris.JSON(map[string]interface{}{
		"message": "",
		"status":  "ok",
		"data":    data,
	})
}
