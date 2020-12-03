package domain

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Creator uint   `gorm:"not null;index" json:"creator"`
}
type ArticleUsecase interface {
}

//contract article repository
type ArticleRepository interface {
}
