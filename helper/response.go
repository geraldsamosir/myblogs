package helper

import (
	echo "github.com/labstack/echo"
)

type resList struct {
	Status    int         `json:"status"`
	Data      interface{} `json:"data"`
	Error     interface{} `json:"error"`
	Page      int         `json:"page"`
	TotalPage int64       `json:"totalPage"`
}

type res struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

type DefaultMessage struct {
	Message string
}

func ResponseList(httpStatus int, data interface{}, error interface{}, page int, totalPage int64, ctx echo.Context) error {
	payload := resList{Status: httpStatus, Data: data, Error: error, Page: page, TotalPage: totalPage}
	return ctx.JSON(httpStatus, payload)

}

func Response(httpStatus int, data interface{}, error interface{}, ctx echo.Context) error {
	payload := res{Status: httpStatus, Data: data, Error: error}
	return ctx.JSON(httpStatus, payload)
}
