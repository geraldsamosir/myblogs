package webserver

import (
	"net/http"
	"strconv"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/geraldsamosir/myblogs/helper"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ArticleHandler struct {
	ArticleUsecase domain.ArticleUsecase
	validation     helper.ValidationRequest
}

func NewArticleHandler(e *echo.Group, ArtUseCase domain.ArticleUsecase, valreq helper.ValidationRequest) {
	handler := &ArticleHandler{
		ArticleUsecase: ArtUseCase,
		validation:     valreq,
	}
	e.GET("/Articles", handler.FindAll)
	e.GET("/Articles/:id", handler.GetByID)
	e.POST("/Articles", handler.Create)
	e.PUT("/Articles/:id", handler.Update)
	e.DELETE("/Articles/:id", handler.DeleteByID)

}

func (Ah *ArticleHandler) FindAll(c echo.Context) error {
	numS := c.QueryParam("page")
	num, _ := strconv.Atoi(numS)
	limmits := c.QueryParam("limit")
	limmit, _ := strconv.Atoi(limmits)
	ctx := c.Request().Context()
	// filter allow
	ID64, _ := strconv.ParseUint(c.QueryParam("id"), 0, 32)
	ID := uint(ID64)
	CategoryID64, _ := strconv.ParseUint(c.QueryParam("categoryId"), 0, 32)
	CategoryID := uint(CategoryID64)
	CreatorID64, _ := strconv.ParseUint(c.QueryParam("creatorId"), 0, 32)
	CreatorID := uint(CreatorID64)

	art := domain.Article{
		ID:         ID,
		Title:      c.QueryParam("title"),
		CategoryID: CategoryID,
		CreatorID:  CreatorID,
	}

	listAr, err := Ah.ArticleUsecase.FindAll(ctx, int64(num), int64(limmit), art)
	if err != nil {
		logrus.Error(err)
		return helper.ResponseList(GetStatusCode(err), nil, err.Error(), 0, 0, c)
	}

	countAr, err := Ah.ArticleUsecase.CountPage(ctx, int64(num), int64(limmit), art)
	if err != nil {
		logrus.Error(err)
		return helper.ResponseList(GetStatusCode(err), nil, err.Error(), 0, 0, c)
	}

	if num <= 0 {
		num = 1
	}

	return helper.ResponseList(http.StatusOK, listAr, nil, num, (countAr), c)
}

func (Ah *ArticleHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	article, err := Ah.ArticleUsecase.GetByID(ctx, int64(id))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}
	//handle notfound
	if article.ID == 0 {
		return helper.Response(http.StatusNotFound, nil, nil, c)
	}
	return helper.Response(GetStatusCode(err), article, nil, c)
}

func (Ah *ArticleHandler) Create(c echo.Context) error {
	var article domain.Article
	ctx := c.Request().Context()
	err := c.Bind(&article)
	if err != nil {
		return helper.Response(http.StatusUnprocessableEntity, nil, nil, c)
	}

	if newErr := Ah.validation.ValidateHandling(article); newErr != nil {
		return helper.Response(http.StatusBadRequest, nil, newErr, c)
	}
	err = Ah.ArticleUsecase.Create(ctx, article)
	if err != nil {
		return helper.Response(http.StatusBadRequest, nil, err, c)
	}

	artc, err := Ah.ArticleUsecase.GetByID(ctx, int64(article.ID))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}

	return helper.Response(http.StatusCreated, artc, nil, c)

}

func (Ah *ArticleHandler) Update(c echo.Context) error {
	var article domain.Article
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	err := c.Bind(&article)
	if err != nil {
		return helper.Response(http.StatusUnprocessableEntity, nil, nil, c)
	}

	err = Ah.ArticleUsecase.Update(ctx, int64(id), article)
	if err != nil {
		return helper.Response(http.StatusBadRequest, nil, err, c)
	}

	artc, err := Ah.ArticleUsecase.GetByID(ctx, int64(article.ID))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}

	return helper.Response(http.StatusOK, artc, nil, c)

}

func (Ah *ArticleHandler) DeleteByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	message, err := Ah.ArticleUsecase.DeleteByID(ctx, int64(id))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}
	return helper.Response(GetStatusCode(err), message, nil, c)
}
