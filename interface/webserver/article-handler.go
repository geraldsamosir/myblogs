package webserver

import (
	"log"
	"net/http"
	"strconv"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/geraldsamosir/myblogs/helper"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

)

type ArticleHandler struct {
	ArticleUsecase domain.ArticleUsecase
}

func NewArticleHandler (e *echo.Group, ArtUseCase domain.ArticleUsecase){
	handler := &ArticleHandler{
		ArticleUsecase: ArtUseCase,
	}
	e.GET("/articles", handler.FindAll)

}


func (Ah * ArticleHandler) FindAll(c echo.Context) error {
	numS := c.QueryParam("page")
	num, _ := strconv.Atoi(numS)
	limmits := c.QueryParam("limit")
	limmit, _ := strconv.Atoi(limmits)
	ctx := c.Request().Context()

	if num <= 0 {
		num = 0 
	} else {
		num  = num * limmit -1
	}


	listAr, err := Ah.ArticleUsecase.FindAll(ctx,  int64(num), int64(limmit))
	if err != nil {
		log.Println(err)
		return helper.Response(getStatusCode(err), nil, err.Error(), c)
	}
	return helper.Response(http.StatusOK, listAr, nil, c)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
