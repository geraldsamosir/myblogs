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
	UserName  string         `gorm:"size:30;unique;not null;index " json:"userName" validate:"required" query:"username"`
	Password  string         `gorm:"size:255;not null" json:"password" validate:"required"`
	Articles  []*Article     `gorm:"foreignKey:CreatorID" json:"articles"`
	RoleID    uint           `gorm:"not null;index" json:"roleId" validate:"required" query:"roleID"`
	Role      Role           `gorm:"foreignKey:RoleID;references:ID"`
}

type UserResponse struct {
	ID        uint `gorm:"primary_key" query:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
	FirstName string         `gorm:"size:255;not null" json:"firstName"`
	LastName  string         `gorm:"size:255" json:"lastName"`
	UserName  string         `gorm:"size:30;unique;not null;unique " json:"userName" validate:"required" query:"username"`
	Password  string         `gorm:"size:255;not null" json:"-" validate:"required"`
	Articles  []Article      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CreatorID" json:"articles"`
	RoleID    uint           `gorm:"not null;index" json:"roleId" validate:"required" query:"roleID"`
	Role      Role           `gorm:"foreignKey:RoleID;references:ID"`
}

type Authentication struct {
	UserName string `json:"username" validate:"required"`
	Password string `json: "password" validate:"required"`
}

type AuthenticationResponse struct {
	UserResponse
	Jwt string `json:"jwt"`
}

func (UserResponse) TableName() string {
	return "users"
}

type UserUsecase interface {
	FindAll(ctx context.Context, page int64, limmit int64, filter UserResponse) ([]UserResponse, error)
	CountPage(ctx context.Context, skip int64, limmit int64, filter UserResponse) (res int64, err error)
	GetByID(ctx context.Context, id int64) (User UserResponse, err error)
	Register(ctx context.Context, userc User) (user User, err error)
	Login(ctx context.Context, user Authentication) (auth AuthenticationResponse, err error, errMessage string)
	Update(ctx context.Context, id int64, userc User) (err error)
	DeleteByID(ctx context.Context, id int64) (message string, err error)
}

//contract role repository
type UserRepository interface {
	FindAll(ctx context.Context, skip int64, limmit int64, filter UserResponse) (res []UserResponse, err error)
	CountAll(ctx context.Context, skip int64, limmit int64, filter UserResponse) (res int64, err error)
	GetByID(ctx context.Context, id int64) (user UserResponse, err error)
	GetByUsername(ctx context.Context, userName string) (user User, err error)
	Store(ctx context.Context, usr *User) (user User, err error)
	Update(ctx context.Context, id int64, user *User) (err error)
	DeleteByID(ctx context.Context, id int64) (err error)
}
