package models

import (
	"context"

	"github.com/geraldsamosir/myblogs/domain"
	"gorm.io/gorm"
)

type ArticlesMysql struct {
	DB *gorm.DB
}

func NewMysqlArticleRepository(Conn *gorm.DB) domain.ArticleRepository {
	return &ArticlesMysql{Conn}
}

func (ArtRepo *ArticlesMysql) FindAll(ctx context.Context, skip int64, limmit int64) (res []domain.Article, err error) {
	if err = ArtRepo.DB.WithContext(ctx).Model(&domain.Article{}).Offset(int(skip)).Limit(int(limmit)).Find(&res).Error; err != nil {
		return nil, err
	}
	return

	// if err = ArtRepo.DB.WithContext(ctx).Preload("Article.User.Role").Find(&res).Error; err != nil {
	// 	return nil, err
	// }
	// return

	// rows, err := ArtRepo.DB.WithContext(ctx).Table("articles").Joins("users").Find(&res).Rows()
	// if err != nil {
	// 	return nil, err
	// }
	// artcl := &domain.Article{}
	// for rows.Next() {
	// 	err = rows.Scan(&artcl.ID, &artcl.Title, &artcl.BannerUrl)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// }

	// return
}
