package main

import (
	"github.com/geraldsamosir/myblogs/infrastructure/database/mysql"
	"github.com/geraldsamosir/myblogs/interface/webserver"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var ws webserver.Webserver

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		logrus.Info("Service RUN on DEBUG mode")
	}
}

func main() {
	var database mysql.Database
	db := database.DatabaseInit()
	ws.RunWebserver(db)
}
