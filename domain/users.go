package domain

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key" query:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	FirstName string     `gorm:"size:255;not null" json:"firstName"`
	LastName  string     `gorm:"size:255" json:"lastName"`
	UserName  string     `gorm:"size:30;not null " json:"userName" query:"username"`
	Password  string     `gorm:"size:255;not null" json:"-"`
	Articles  []Article  `gorm:"-"; json:"articles"`
	RoleID    uint       `gorm:"not null;index" json:"role" query:"role"`
	Role      Role       `gorm:"foreignKey:RoleID;references:id"`
}
