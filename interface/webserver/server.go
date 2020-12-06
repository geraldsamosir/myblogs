package webserver

import (
	"log"
	"net/http"
	"time"

	"github.com/geraldsamosir/myblogs/helper"
	_repo "github.com/geraldsamosir/myblogs/infrastructure/database/mysql/models"
	_authmid "github.com/geraldsamosir/myblogs/interface/webserver/middleware"
	_usecase "github.com/geraldsamosir/myblogs/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Webserver struct {
}

func (ws *Webserver) RunWebserver(db *gorm.DB) {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	var validation helper.ValidationRequest
	var passwordHandling helper.Password
	var midlauth _authmid.Auth

	// domain handle
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	artRepo := _repo.NewMysqlArticleRepository(db)
	artUsecase := _usecase.NewArticleUsecase(artRepo, timeoutContext)

	catRepo := _repo.NewMysqlCategoryRepository(db)
	catUsecase := _usecase.NewcategoryUsecase(catRepo, timeoutContext)

	roleRepo := _repo.NewMysqlRoleRepository(db)
	roleUsecase := _usecase.NewRoleUsecase(roleRepo, timeoutContext)

	userRepo := _repo.NewMysqlUserRepository(db)
	userUsecase := _usecase.NewUserUsecase(userRepo, timeoutContext, passwordHandling, midlauth)

	//midleware
	authmidl := _authmid.InitMiddleware()
	fs := http.FileServer(http.Dir("public/build"))
	e.GET("/*", echo.WrapHandler(fs))

	api := e.Group("/api")
	message := helper.DefaultMessage{Message: "helo this is root api"}

	api.Use(authmidl.MiddlewareAuth)
	api.GET("/", func(ctx echo.Context) error {
		return helper.Response(http.StatusOK, message, nil, ctx)
	})

	NewArticleHandler(api, artUsecase, validation)
	NewCategoryHandler(api, catUsecase, validation)
	NewRoleHandler(api, roleUsecase, validation)
	NewUserHandler(api, userUsecase, validation)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
