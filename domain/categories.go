package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           uint `gorm:"primary_key" query:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `sql:"index"`
	CategoryName string         `gorm:"not null" json:"categoryName" validate:"required" query:"categoryName"`
	Articles     []*Article     `gorm:"foreignKey:Category`
}

type CategoryUsecase interface {
	FindAll(ctx context.Context, page int64, limmit int64, filter Category) ([]Category, error)
	CountPage(ctx context.Context, skip int64, limmit int64, filter Category) (res int64, err error)
	GetByID(ctx context.Context, id int64) (Category Category, err error)
	Create(ctx context.Context, artc *Category) (err error)
	Update(ctx context.Context, id int64, artc *Category) (err error)
	DeleteByID(ctx context.Context, id int64) (message string, err error)
}

//contract article repository
type CategoryRepository interface {
	FindAll(ctx context.Context, skip int64, limmit int64, filter Category) (res []Category, err error)
	CountAll(ctx context.Context, skip int64, limmit int64, filter Category) (res int64, err error)
	GetByID(ctx context.Context, id int64) (Category Category, err error)
	Store(ctx context.Context, art *Category) (err error)
	Update(ctx context.Context, id int64, art *Category) (err error)
	DeleteByID(ctx context.Context, id int64) (err error)
}
