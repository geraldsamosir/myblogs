package webserver

import (
	"net/http"
	"strconv"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/geraldsamosir/myblogs/helper"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
	validation  helper.ValidationRequest
}

func NewUserHandler(e *echo.Group, ArtUseCase domain.UserUsecase, valreq helper.ValidationRequest) {
	handler := &UserHandler{
		UserUsecase: ArtUseCase,
		validation:  valreq,
	}
	e.GET("/Users", handler.FindAll)
	e.GET("/Users/:id", handler.GetByID)
	e.POST("/Users", handler.Create)
	e.PUT("/Users/:id", handler.Update)
	e.DELETE("/Users/:id", handler.DeleteByID)
	e.POST("/Users/Login", handler.Login)

}

func (Uh *UserHandler) FindAll(c echo.Context) error {
	numS := c.QueryParam("page")
	num, _ := strconv.Atoi(numS)
	limmits := c.QueryParam("limit")
	limmit, _ := strconv.Atoi(limmits)
	ctx := c.Request().Context()
	// filter allow
	RoleID64, _ := strconv.ParseUint(c.QueryParam("roleID"), 0, 32)
	RoleID64ID := uint(RoleID64)

	usr := domain.UserResponse{
		UserName: c.QueryParam("userName"),
		RoleID:   RoleID64ID,
	}

	listAr, err := Uh.UserUsecase.FindAll(ctx, int64(num), int64(limmit), usr)
	if err != nil {
		logrus.Error(err)
		return helper.ResponseList(GetStatusCode(err), nil, err.Error(), 0, 0, c)
	}

	countAr, err := Uh.UserUsecase.CountPage(ctx, int64(num), int64(limmit), usr)
	if err != nil {
		logrus.Error(err)
		return helper.ResponseList(GetStatusCode(err), nil, err.Error(), 0, 0, c)
	}

	if num <= 0 {
		num = 1
	}

	return helper.ResponseList(http.StatusOK, listAr, nil, num, (countAr), c)
}

func (Uh *UserHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	User, err := Uh.UserUsecase.GetByID(ctx, int64(id))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}
	//handle notfound
	if User.ID == 0 {
		return helper.Response(http.StatusNotFound, nil, nil, c)
	}
	return helper.Response(GetStatusCode(err), User, nil, c)
}

func (Uh *UserHandler) Create(c echo.Context) error {
	var User domain.User
	ctx := c.Request().Context()
	err := c.Bind(&User)
	if err != nil {
		return helper.Response(http.StatusUnprocessableEntity, nil, nil, c)
	}

	if newErr := Uh.validation.ValidateHandling(User); newErr != nil {
		return helper.Response(http.StatusBadRequest, nil, newErr, c)
	}
	User, err = Uh.UserUsecase.Register(ctx, User)
	if err != nil {
		return helper.Response(http.StatusBadRequest, nil, err, c)
	}

	artc, err := Uh.UserUsecase.GetByID(ctx, int64(User.ID))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}

	return helper.Response(http.StatusCreated, artc, nil, c)

}

func (Uh *UserHandler) Login(c echo.Context) error {
	var User domain.Authentication
	ctx := c.Request().Context()
	err := c.Bind(&User)
	if err != nil {
		return helper.Response(http.StatusUnprocessableEntity, nil, nil, c)
	}

	if newErr := Uh.validation.ValidateHandling(User); newErr != nil {
		return helper.Response(http.StatusUnauthorized, nil, newErr, c)
	}
	user, err, message := Uh.UserUsecase.Login(ctx, User)
	if err != nil {
		return helper.Response(http.StatusUnauthorized, nil, message, c)
	}

	return helper.Response(GetStatusCode(err), user, nil, c)

}

func (Uh *UserHandler) Update(c echo.Context) error {
	var User domain.User
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	err := c.Bind(&User)
	if err != nil {
		return helper.Response(http.StatusUnprocessableEntity, nil, nil, c)
	}

	err = Uh.UserUsecase.Update(ctx, int64(id), User)
	if err != nil {
		return helper.Response(http.StatusBadRequest, nil, err, c)
	}

	artc, err := Uh.UserUsecase.GetByID(ctx, int64(User.ID))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}

	return helper.Response(http.StatusOK, artc, nil, c)

}

func (Uh *UserHandler) DeleteByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	message, err := Uh.UserUsecase.DeleteByID(ctx, int64(id))
	if err != nil {
		return helper.Response(GetStatusCode(err), nil, err.Error(), c)
	}
	return helper.Response(GetStatusCode(err), message, nil, c)
}
