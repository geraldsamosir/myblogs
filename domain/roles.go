package domain

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	RoleName string `gorm:"size:255;not null" json:"roleName"`
	Users []User `json:"users"`
} 