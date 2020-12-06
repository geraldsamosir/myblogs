package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primary_key" query:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
	FirstName string         `gorm:"size:255;not null" json:"firstName"`
	LastName  string         `gorm:"size:255" json:"lastName"`
	UserName  string         `gorm:"size:30;not null " json:"userName" validate:"required" query:"username"`
	Password  string         `gorm:"size:255;not null" json:"-" validate:"required"`
	Articles  []Article      `gorm:"-"; json:"articles"`
	RoleID    uint           `gorm:"not null;index" json:"role" validate:"required" query:"roleID"`
	Role      Role           `gorm:"foreignKey:RoleID;references:ID"`
}

type Authentication struct {
	UserName string `json:"username" validate:"required"`
	Password string `json: "password" validate:"required"`
}

type AuthenticationResponse struct {
	*User
	Jwt string `json:"jwt"`
}

type UserUsecase interface {
	FindAll(ctx context.Context, page int64, limmit int64, filter User) ([]User, error)
	CountPage(ctx context.Context, skip int64, limmit int64, filter Role) (res int64, err error)
	GetByID(ctx context.Context, id int64) (User User, err error)
	Register(ctx context.Context, userc *User) (err error)
	Login(ctx context.Context, user *Authentication) (ar AuthenticationResponse, err error)
	Update(ctx context.Context, id int64, userc *User) (err error)
	DeleteByID(ctx context.Context, id int64) (message string, err error)
}

//contract role repository
type UserRepository interface {
	FindAll(ctx context.Context, skip int64, limmit int64, filter User) (res []User, err error)
	CountAll(ctx context.Context, skip int64, limmit int64, filter User) (res int64, err error)
	GetByID(ctx context.Context, id int64) (user User, err error)
	Store(ctx context.Context, user *User) (err error)
	Update(ctx context.Context, id int64, user *User) (err error)
	DeleteByID(ctx context.Context, id int64) (err error)
}
