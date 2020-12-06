package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint `gorm:"primary_key" query:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
	RoleName  string         `gorm:"size:255;not null" json:"roleName" validate:"required" query:"roleName"`
	User      []*User        `gorm:"-"`
}

type RoleUsecase interface {
	FindAll(ctx context.Context, page int64, limmit int64, filter Role) ([]Role, error)
	CountPage(ctx context.Context, skip int64, limmit int64, filter Role) (res int64, err error)
	GetByID(ctx context.Context, id int64) (Role Role, err error)
	Create(ctx context.Context, rolec *Role) (err error)
	Update(ctx context.Context, id int64, rolec *Role) (err error)
	DeleteByID(ctx context.Context, id int64) (message string, err error)
}

//contract role repository
type RoleRepository interface {
	FindAll(ctx context.Context, skip int64, limmit int64, filter Role) (res []Role, err error)
	CountAll(ctx context.Context, skip int64, limmit int64, filter Role) (res int64, err error)
	GetByID(ctx context.Context, id int64) (role Role, err error)
	Store(ctx context.Context, role *Role) (err error)
	Update(ctx context.Context, id int64, role *Role) (err error)
	DeleteByID(ctx context.Context, id int64) (err error)
}
