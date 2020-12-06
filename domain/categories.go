package domain

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           uint `gorm:"primary_key" query:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `sql:"index"`
	CategoryName string         `gorm:"not null" json:"categoryName" query:"categoryName"`
	Articles     []*Article     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:Category`
}
