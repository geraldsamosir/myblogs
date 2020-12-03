package domain

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title     string `json:"title" validate:"required"`
	Content   string `gorm:"type:TEXT" json:"content" validate:"required"`
	Creator   uint   `gorm:"not null;index" json:"creator"`
	BannerUrl string `json:"bannerUrl"`
	Category  uint   `gorm:"not null;index" json:"Article"`
	Slug      string `json:"slug" validate:"required"`
}
type ArticleUsecase interface {
}

//contract article repository
type ArticleRepository interface {
}
