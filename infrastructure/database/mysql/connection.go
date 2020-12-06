package mysql

import (
	"fmt"
	"net/url"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	url      string
	loc      string
}

func (database *Database) DatabaseInit() *gorm.DB {
	database.Host = viper.GetString("database.DB_HOST")
	database.Port = viper.GetString("database.DB_PORT")
	database.User = viper.GetString("database.DB_USER")
	database.Password = viper.GetString("database.DB_PASS")
	database.Name = viper.GetString("database.DB_NAME")
	database.loc = viper.GetString("location")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", database.User, database.Password, database.Host, database.Port, database.Name)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("gorm:table_options", "collation_connection=utf8_general_ci")
	db.Set("gorm:auto_preload", true)
	if err != nil {
		panic(err)
	}
	logrus.Info("database  connect")

	// do migarion table in here

	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Role{})
	db.AutoMigrate(&domain.Article{})
	return db
}
