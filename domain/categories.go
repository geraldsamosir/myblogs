package domain

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string    `gorm:"not null" json:"categoryName"`
	Articles     []Article `json:"articles"`
}
