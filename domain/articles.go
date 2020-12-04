package domain

import (
	"context"
	"time"
)

type Article struct {
	ID         uint `gorm:"primary_key" query:"id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	Title      string     `json:"title" validate:"required" query:"title"`
	Content    string     `gorm:"type:TEXT" json:"content" validate:"required"`
	CreatorID  uint       `gorm:"not null;index;" json:"creator" query:"creatorID"`
	Creator    User       `gorm:"foreignKey:CreatorID;references:ID", json:"creator"`
	BannerUrl  string     `json:"bannerUrl"`
	CategoryID uint       `gorm:"not null;index" json:"categoryID" 
	query:"category"`
	Categories Category `gorm:"foreignKey:CategoryID;references:id"`
	Slug       string   `json:"slug" validate:"required" query:"slug"`
}
type ArticleUsecase interface {
	FindAll(ctx context.Context, skip int64, limmit int64) ([]Article, error)
	// Filter(ctx context.Context, filter Article) (Article, error)
	// Update(ctx context.Context, id int64, ar *Article) error
	// Create(context.Context, *Article) error
	// Delete(ctx context.Context, id int64)
}

//contract article repository
type ArticleRepository interface {
	FindAll(ctx context.Context, skip int64, limmit int64) (res []Article, err error)
	// Filter(ctx context.Context, filter Article) (Article, error)
	// Update(ctx context.Context, ar *Article) error
	// Store(ctx context.Context, a *Article) error
	// Delete(ctx context.Context, id int64) error
}
