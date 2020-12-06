package domain

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint `gorm:"primary_key" query:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
	RoleName  string         `gorm:"size:255;not null" json:"roleName" query:"roleName"`
	User      []*User        `gorm:"-"`
}
