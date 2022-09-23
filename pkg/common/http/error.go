package http

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	AppCode        int64  `json:"code,omitempty"`
	Errortext      string `json:"error,omitempty"`
}

func (e *ErrorResponse) Render(e http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)

	return nil
}

func ErrInternal(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		Errortext:      err.Error(),
	}
}

func ErrBadRequest(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Errortext:      err.Error(),
	}
}
