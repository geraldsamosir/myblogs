package webserver

import (
	"net/http"
	"strconv"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/geraldsamosir/myblogs/helper"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type CategoryHandler struct {
	CategoryUsecase domain.CategoryUsecase
}

func NewCategoryHandler(e *echo.Group, ArtUseCase domain.CategoryUsecase) {
	handler := &CategoryHandler{
		CategoryUsecase: ArtUseCase,
	}
	e.GET("/Categories", handler.FindAll)
	e.GET("/Categories/:id", handler.GetByID)
	e.POST("/Categories", handler.Create)
	e.PUT("/Categories/:id", handler.Update)
	e.DELETE("/Categories/:id", handler.DeleteByID)

}

func (Ah *CategoryHandler) FindAll(c echo.Context) error {
	numS := c.QueryParam("page")
	num, _ := strconv.Atoi(numS)
	limmits := c.QueryParam("limit")
	limmit, _ := strconv.Atoi(limmits)
	ctx := c.Request().Context()
	// filter allow
	ID64, _ := strconv.ParseUint(c.QueryParam("id"), 0, 32)
	ID := uint(ID64)

	art := domain.Category{
		ID:           ID,
		CategoryName: c.QueryParam("categoryName"),
	}

	listAr, err := Ah.CategoryUsecase.FindAll(ctx, int64(num), int64(limmit), art)
	if err != nil {
		logrus.Error(err)
		return helper.ResponseList(GetStatusCode(err), nil, err.Error(), 0, 0, c)
	}

	countAr, err := Ah.CategoryUsecase.CountPage(ctx, int64(num), int64(limmit), art)
	if err != nil {
		logrus.Error(err)
		return helper.ResponseList(GetStatusCode(err), nil, err.Error(), 0, 0, c)
	}

	if num <= 0 {
		num = 1
	}

	return helper.ResponseList(http.StatusOK, listAr, nil, num, (countAr), c)
}

func (Ah *CategoryHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	Category, err := Ah.CategoryUsecase.GetByID(ctx, int64(id))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}
	//handle notfound
	if Category.ID == 0 {
		return helper.Response(http.StatusNotFound, nil, nil, c)
	}
	return helper.Response(GetStatusCode(err), Category, nil, c)
}

func (Ah *CategoryHandler) Create(c echo.Context) error {
	var Category domain.Category
	ctx := c.Request().Context()
	err := c.Bind(&Category)
	if err != nil {
		return helper.Response(http.StatusUnprocessableEntity, nil, nil, c)
	}

	if newErr := validation.ValidateHandling(Category); newErr != nil {
		return helper.Response(http.StatusBadRequest, nil, newErr, c)
	}
	err = Ah.CategoryUsecase.Create(ctx, &Category)
	if err != nil {
		return helper.Response(http.StatusBadRequest, nil, err, c)
	}

	artc, err := Ah.CategoryUsecase.GetByID(ctx, int64(Category.ID))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}

	return helper.Response(http.StatusCreated, artc, nil, c)

}

func (Ah *CategoryHandler) Update(c echo.Context) error {
	var Category domain.Category
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	err := c.Bind(&Category)
	if err != nil {
		return helper.Response(http.StatusUnprocessableEntity, nil, nil, c)
	}

	err = Ah.CategoryUsecase.Update(ctx, int64(id), &Category)
	if err != nil {
		return helper.Response(http.StatusBadRequest, nil, err, c)
	}

	artc, err := Ah.CategoryUsecase.GetByID(ctx, int64(Category.ID))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}

	return helper.Response(http.StatusOK, artc, nil, c)

}

func (Ah *CategoryHandler) DeleteByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	message, err := Ah.CategoryUsecase.DeleteByID(ctx, int64(id))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}
	return helper.Response(GetStatusCode(err), message, nil, c)
}
