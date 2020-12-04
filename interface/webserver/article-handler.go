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
	// filter allow
	ID64, _ := strconv.ParseUint(c.QueryParam("id"), 0, 32)
	ID := uint(ID64)
	CategoryID64, _:=strconv.ParseUint(c.QueryParam("categoryId"), 0, 32)
	CategoryID := uint(CategoryID64)
	CreatorID64, _:=strconv.ParseUint(c.QueryParam("creatorId"), 0, 32)
	CreatorID := uint(CreatorID64)


	art := domain.Article{
		ID: ID,
		Title: c.QueryParam("title"),
		CategoryID: CategoryID,
		CreatorID: CreatorID,
	}


	listAr, err := Ah.ArticleUsecase.FindAll(ctx,  int64(num), int64(limmit), art)
	if err != nil {
		log.Println(err)
		return helper.ResponseList(getStatusCode(err), nil, err.Error(), 0, 0, c)
	}

	countAr , err := Ah.ArticleUsecase.CountAll(ctx,  int64(num), int64(limmit), art)
	if err != nil {
		log.Println(err)
		return helper.ResponseList(getStatusCode(err), nil, err.Error(), 0, 0, c)
	}

	if num <= 0 {
		num = 1
	}

	return helper.ResponseList(http.StatusOK, listAr, nil, num, countAr, c)
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
