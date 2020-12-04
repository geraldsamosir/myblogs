package main

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/geraldsamosir/myblogs/helper"
	"github.com/geraldsamosir/myblogs/infrastructure/database/mysql"
	_repo "github.com/geraldsamosir/myblogs/infrastructure/database/mysql/models"
	_handler "github.com/geraldsamosir/myblogs/interface/webserver"
	_usecase "github.com/geraldsamosir/myblogs/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", viper.GetString("location"))
	var database mysql.Database
	db := database.DatabaseInit()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	//e.Use(middleware.Static("public/build"))
	// middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	// e.Use(middL.CORS)
	// authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	// ar := _articleRepo.NewMysqlArticleRepository(dbConn)

	// timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	// au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	// _articleHttpDelivery.NewArticleHandler(e, au)

	// domain handle
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	artRepo := _repo.NewMysqlArticleRepository(db)
	artUsecase := _usecase.NewArticleUsecase(artRepo, timeoutContext)
	fs := http.FileServer(http.Dir("public/build"))
	e.GET("/*", echo.WrapHandler(fs))

	api := e.Group("/api")
	message := helper.DefaultMessage{Message: "helo this is root api"}
	api.GET("/", func(ctx echo.Context) error {
		return helper.Response(http.StatusOK, message, nil, ctx)
	})

	_handler.NewArticleHandler(api, artUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
