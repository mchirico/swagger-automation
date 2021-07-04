package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// HeadRequest represents body of Head request.
type HeadRequest struct {
	Msg string `json:"msg"`
}

// HeadResponse represents body of Head response.
type HeadResponse struct {
	Baz struct {
		Prop string `json:"prop"`
	} `json:"baz"`
}

// FooBarHandler handles incoming foobar requests
func HeadHandler(ctx echo.Context) error {
	req := HeadRequest{}
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp := actionRequest(req)
	return ctx.JSON(http.StatusOK, resp)
}

func actionRequest(req HeadRequest) FooBarResponse {
	return FooBarResponse{}
}
