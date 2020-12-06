package webserver

import (
	"log"
	"net/http"
	"time"

	"github.com/geraldsamosir/myblogs/helper"
	_repo "github.com/geraldsamosir/myblogs/infrastructure/database/mysql/models"
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

	// domain handle
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	artRepo := _repo.NewMysqlArticleRepository(db)
	artUsecase := _usecase.NewArticleUsecase(artRepo, timeoutContext)

	catRepo := _repo.NewMysqlCategoryRepository(db)
	catUsecase := _usecase.NewcategoryUsecase(catRepo, timeoutContext)

	roleRepo := _repo.NewMysqlRoleRepository(db)
	roleUsecase := _usecase.NewRoleUsecase(roleRepo, timeoutContext)

	fs := http.FileServer(http.Dir("public/build"))
	e.GET("/*", echo.WrapHandler(fs))

	api := e.Group("/api")
	message := helper.DefaultMessage{Message: "helo this is root api"}
	api.GET("/", func(ctx echo.Context) error {
		return helper.Response(http.StatusOK, message, nil, ctx)
	})

	NewArticleHandler(api, artUsecase)
	NewCategoryHandler(api, catUsecase)
	NewRoleHandler(api, roleUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
