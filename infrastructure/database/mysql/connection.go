package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/geraldsamosir/myblogs/domain"
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
}

func (database *Database) DatabaseInit() *gorm.DB {
	database.Host = viper.GetString("database.DB_HOST")
	database.Port = viper.GetString("database.DB_PORT")
	database.User = viper.GetString("database.DB_USER")
	database.Password = viper.GetString("database.DB_PASS")
	database.Name = viper.GetString("database.DB_NAME")
	//database.url = database.User + ":" + database.Password + "@(" + database.Host + ":" + database.Port + ")/" + database.Name + "?charset=utf8&parseTime=True&loc=Local"
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", database.User, database.Password, database.Host, database.Port, database.Name)
	dsn := fmt.Sprintf("%s?%s", connection)

	dbConn, err := sql.Open(`mysql`, dsn)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: dbConn,
	}), &gorm.Config{})
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("gorm:table_options", "collation_connection=utf8_general_ci")
	db.Set("gorm:auto_preload", true)
	if err != nil {
		panic(err)
	}
	log.Println("database  connect")

	// do migarion table in here

	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Article{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Role{})

	// add relation
	// db.Model(&domain.Article{}).AddForeignKey("creator", "users(id)", "CASCADE", "CASCADE")
	// db.Model(&domain.User{}).AddForeignKey("role", "roles(id)", "CASCADE", "CASCADE")
	// db.Model(&domain.Article{}).AddForeignKey("category", "categories(id)", "CASCADE", "CASCADE")
	return db
}
