package webserver

import (
	"log"
	"net/http"
	"time"

	"github.com/geraldsamosir/myblogs/helper"
	_repo "github.com/geraldsamosir/myblogs/infrastructure/database/mysql/models"
	_cloudinary "github.com/geraldsamosir/myblogs/infrastructure/filesystem/cloudinary"
	_authmid "github.com/geraldsamosir/myblogs/interface/webserver/middleware"
	_usecase "github.com/geraldsamosir/myblogs/usecase"
	"github.com/komfy/cloudinary"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	_ "github.com/geraldsamosir/myblogs/interface/webserver/docs"
)

type Webserver struct {
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /api

func (ws *Webserver) RunWebserver(db *gorm.DB) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))
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
	cloudiSvc, err := cloudinary.NewService(viper.GetString("Cloudinary_URL"))
	if err != nil {
		logrus.Error(err)
	}
	cloudinyInterface := _cloudinary.NewCloudinary(cloudiSvc)

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
	api.GET("/swagger/*", echoSwagger.WrapHandler)
	NewArticleHandler(api, artUsecase, validation)
	NewCategoryHandler(api, catUsecase, validation)
	NewRoleHandler(api, roleUsecase, validation)
	NewUserHandler(api, userUsecase, validation)
	NewFilesystemHandler(api, cloudinyInterface, validation)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
