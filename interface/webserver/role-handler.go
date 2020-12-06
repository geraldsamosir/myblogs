package webserver

import (
	"net/http"
	"strconv"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/geraldsamosir/myblogs/helper"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type RoleHandler struct {
	RoleUsecase domain.RoleUsecase
}

func NewRoleHandler(e *echo.Group, ArtUseCase domain.RoleUsecase) {
	handler := &RoleHandler{
		RoleUsecase: ArtUseCase,
	}
	e.GET("/Roles", handler.FindAll)
	e.GET("/Roles/:id", handler.GetByID)
	e.POST("/Roles", handler.Create)
	e.PUT("/Roles/:id", handler.Update)
	e.DELETE("/Roles/:id", handler.DeleteByID)

}

func (Ah *RoleHandler) FindAll(c echo.Context) error {
	numS := c.QueryParam("page")
	num, _ := strconv.Atoi(numS)
	limmits := c.QueryParam("limit")
	limmit, _ := strconv.Atoi(limmits)
	ctx := c.Request().Context()
	// filter allow
	ID64, _ := strconv.ParseUint(c.QueryParam("id"), 0, 32)
	ID := uint(ID64)

	art := domain.Role{
		ID:       ID,
		RoleName: c.QueryParam("roleName"),
	}

	listAr, err := Ah.RoleUsecase.FindAll(ctx, int64(num), int64(limmit), art)
	if err != nil {
		logrus.Error(err)
		return helper.ResponseList(GetStatusCode(err), nil, err.Error(), 0, 0, c)
	}

	countAr, err := Ah.RoleUsecase.CountPage(ctx, int64(num), int64(limmit), art)
	if err != nil {
		logrus.Error(err)
		return helper.ResponseList(GetStatusCode(err), nil, err.Error(), 0, 0, c)
	}

	if num <= 0 {
		num = 1
	}

	return helper.ResponseList(http.StatusOK, listAr, nil, num, (countAr), c)
}

func (Ah *RoleHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	Role, err := Ah.RoleUsecase.GetByID(ctx, int64(id))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}
	//handle notfound
	if Role.ID == 0 {
		return helper.Response(http.StatusNotFound, nil, nil, c)
	}
	return helper.Response(GetStatusCode(err), Role, nil, c)
}

func (Ah *RoleHandler) Create(c echo.Context) error {
	var Role domain.Role
	ctx := c.Request().Context()
	err := c.Bind(&Role)
	if err != nil {
		return helper.Response(http.StatusUnprocessableEntity, nil, nil, c)
	}

	if newErr := validation.ValidateHandling(Role); newErr != nil {
		return helper.Response(http.StatusBadRequest, nil, newErr, c)
	}
	err = Ah.RoleUsecase.Create(ctx, &Role)
	if err != nil {
		return helper.Response(http.StatusBadRequest, nil, err, c)
	}

	artc, err := Ah.RoleUsecase.GetByID(ctx, int64(Role.ID))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}

	return helper.Response(http.StatusCreated, artc, nil, c)

}

func (Ah *RoleHandler) Update(c echo.Context) error {
	var Role domain.Role
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	err := c.Bind(&Role)
	if err != nil {
		return helper.Response(http.StatusUnprocessableEntity, nil, nil, c)
	}

	err = Ah.RoleUsecase.Update(ctx, int64(id), &Role)
	if err != nil {
		return helper.Response(http.StatusBadRequest, nil, err, c)
	}

	artc, err := Ah.RoleUsecase.GetByID(ctx, int64(Role.ID))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}

	return helper.Response(http.StatusOK, artc, nil, c)

}

func (Ah *RoleHandler) DeleteByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	message, err := Ah.RoleUsecase.DeleteByID(ctx, int64(id))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}
	return helper.Response(GetStatusCode(err), message, nil, c)
}
