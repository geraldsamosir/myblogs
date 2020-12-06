package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID         uint `gorm:"primary_key" query:"id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `sql:"index"`
	Title      string         `json:"title" validate:"required" query:"title"`
	Content    string         `gorm:"type:TEXT" json:"content" validate:"required"`
	CreatorID  uint           `gorm:"not null;index;" json:"creatorID" validate:"required"  query:"creatorID"`
	Creator    User           `gorm:"foreignKey:CreatorID;references:ID", json:"creator"`
	BannerUrl  string         `json:"bannerUrl"`
	CategoryID uint           `gorm:"not null;index" json:"categoryID" validate:"required"  query:"category"`
	Categories Category       `gorm:"foreignKey:CategoryID;references:id"`
	Slug       string         `json:"slug" validate:"required" query:"slug"`
}
type ArticleUsecase interface {
	FindAll(ctx context.Context, page int64, limmit int64, filter Article) ([]Article, error)
	CountPage(ctx context.Context, skip int64, limmit int64, filter Article) (res int64, err error)
	GetByID(ctx context.Context, id int64) (article Article, err error)
	Create(ctx context.Context, artc *Article) (err error)
	Update(ctx context.Context, id int64, artc *Article) (err error)
	DeleteByID(ctx context.Context, id int64) (message string, err error)
}

//contract article repository
type ArticleRepository interface {
	FindAll(ctx context.Context, skip int64, limmit int64, filter Article) (res []Article, err error)
	CountAll(ctx context.Context, skip int64, limmit int64, filter Article) (res int64, err error)
	GetByID(ctx context.Context, id int64) (article Article, err error)
	Store(ctx context.Context, art *Article) (err error)
	Update(ctx context.Context, id int64, art *Article) (err error)
	DeleteByID(ctx context.Context, id int64) (err error)
}
