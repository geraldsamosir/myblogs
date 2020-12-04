package domain

import (
	"time"
)

type Role struct {
	ID        uint `gorm:"primary_key" query:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	RoleName  string     `gorm:"size:255;not null" json:"roleName" query:"roleName"`
	User      []*User    `gorm:"-"`
}
