package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/sirupsen/logrus"
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
	if err = ArtRepo.DB.Scopes(Paginate(ctx, skip, limmit, filter)).Preload(clause.Associations).Preload("Creator.Role").Find(&res).Error; err != nil {
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

func (ArtRepo *ArticlesMysql) GetByID(ctx context.Context, id int64) (art domain.Article, err error) {
	ID := uint(id)
	article := domain.Article{}
	artc := ArtRepo.DB.WithContext(ctx).Preload(clause.Associations).Preload("Creator.Role").Where(domain.Article{ID: ID}).Find(&article)
	if err = artc.Error; err != nil {
		return art, err
	}
	art = article
	return
}

func (ArtRepo *ArticlesMysql) Store(ctx context.Context, art *domain.Article) (err error) {
	if err = ArtRepo.DB.Create(&art).Preload(clause.Associations).Preload("Creator.Role").Error; err != nil {
		return err
	}
	return
}

func (ArtRepo *ArticlesMysql) Update(ctx context.Context, id int64, artc *domain.Article) (err error) {
	var article domain.Article
	ArtRepo.DB.WithContext(ctx).First(&article, id)
	checkart, _ := ArtRepo.GetByID(ctx, id)
	if checkart.ID == 0 {
		errMessage := fmt.Sprintf("article with id %s  not exisit", string(id))
		errors.New(errMessage)
	}
	result := ArtRepo.DB.WithContext(ctx).Model(&article).Updates(artc)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (ArtRepo *ArticlesMysql) DeleteByID(ctx context.Context, id int64) (err error) {
	artc := ArtRepo.DB.WithContext(ctx).Where("id = ?", id).Delete(&domain.Article{})
	if err = artc.Error; err != nil {
		logrus.Error("err", err)
		return err
	}
	return
}
