package helper

import (
	echo "github.com/labstack/echo"
)

type res struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

func Response(httpStatus int, data interface{}, error interface{}, ctx echo.Context) error {
	payload := res{Status: httpStatus, Data: data, Error: error}
	return ctx.JSON(httpStatus, payload)
}
