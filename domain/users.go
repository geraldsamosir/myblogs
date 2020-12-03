package domain

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `gorm:"size:255" json:"first_name"`
	LastName  string    `gorm:"size:255" json:"last_name"`
	UserName  string    `gorm:"size:30" json:"user_name"`
	Password  string    `gorm:"size:255" json:"-"`
	Articles  []Article `json:"articles"`
}
