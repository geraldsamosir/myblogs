package helper

import (
	echo "github.com/labstack/echo"
)

type ResList struct {
	Status    int         `json:"status"`
	Data      interface{} `json:"data"`
	Error     interface{} `json:"error"`
	Page      int         `json:"page"`
	TotalPage int64       `json:"totalPage"`
}

type Res struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

type DefaultMessage struct {
	Message string
}

func ResponseList(httpStatus int, data interface{}, error interface{}, page int, totalPage int64, ctx echo.Context) error {
	payload := ResList{Status: httpStatus, Data: data, Error: error, Page: page, TotalPage: totalPage}
	return ctx.JSON(httpStatus, payload)

}

func Response(httpStatus int, data interface{}, error interface{}, ctx echo.Context) error {
	payload := Res{Status: httpStatus, Data: data, Error: error}
	return ctx.JSON(httpStatus, payload)
}
