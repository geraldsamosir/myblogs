package domain

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `gorm:"size:255;not null" json:"firstName"`
	LastName  string    `gorm:"size:255" json:"lastName"`
	UserName  string    `gorm:"size:30;not null " json:"userName"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Articles  []Article `json:"articles"`
	Role      uint      `gorm:"not null;index" json:"role"`
}
