package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/geraldsamosir/myblogs/infrastructure/database/mysql"
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
	fmt.Println(db)
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	// middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	// e.Use(middL.CORS)
	// authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	// ar := _articleRepo.NewMysqlArticleRepository(dbConn)

	// timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	// au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	// _articleHttpDelivery.NewArticleHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
