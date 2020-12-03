package mysql

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dialect

	//"github.com/geraldsamosir/myblogs/domain"
	"github.com/spf13/viper"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	url      string
}

func (database *Database) DatabaseInit() *gorm.DB {
	database.Host = viper.GetString("database.DB_HOST")
	database.Port = viper.GetString("database.DB_PORT")
	database.User = viper.GetString("database.DB_USER")
	database.Password = viper.GetString("database.DB_PASS")
	database.Name = viper.GetString("database.DB_NAME")
	database.url = database.User + ":" + database.Password + "@(" + database.Host + ":" + database.Port + ")/" + database.Name + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", database.url)
	if err != nil {
		panic(err)
	}
	log.Println("database  connect")
	// db.AutoMigrate(&domain.Article{})
	//	db.AutoMigrate(&models.Book{})

	return db
}
