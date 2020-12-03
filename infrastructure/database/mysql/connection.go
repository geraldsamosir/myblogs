package mysql

import (
	"log"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dialect

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

	// do migarion table in here

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Article{})
	db.AutoMigrate(&domain.Role{})
	db.AutoMigrate(&domain.Category{})

	// add relation
	db.Model(&domain.Article{}).AddForeignKey("creator", "users(id)", "CASCADE", "CASCADE")
	db.Model(&domain.User{}).AddForeignKey("role", "roles(id)", "CASCADE", "CASCADE")
	db.Model(&domain.Article{}).AddForeignKey("category", "categories(id)", "CASCADE", "CASCADE")
	return db
}
