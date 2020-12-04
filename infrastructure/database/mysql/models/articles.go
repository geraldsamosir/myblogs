package models

import (
	"context"

	"github.com/geraldsamosir/myblogs/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ArticlesMysql struct {
	DB *gorm.DB
}

func NewMysqlArticleRepository(Conn *gorm.DB) domain.ArticleRepository {
	return &ArticlesMysql{Conn}
}

func (ArtRepo *ArticlesMysql) FindAll(ctx context.Context, skip int64, limmit int64, filter domain.Article) (res []domain.Article, err error) {
	if err = ArtRepo.DB.WithContext(ctx).Offset(int(skip)).Limit(int(limmit)).Where(filter).Preload(clause.Associations).Preload("Creator.Role").Find(&res).Error; err != nil {
		return nil, err
	}
	return
}

func (ArtRepo *ArticlesMysql) CountAll(ctx context.Context, skip int64, limmit int64, filter domain.Article) (res int64, err error) {
	if err = ArtRepo.DB.WithContext(ctx).Model(&domain.Article{}).Where(filter).Count(&res).Error; err != nil {
		return 0, err
	}
	return
}
