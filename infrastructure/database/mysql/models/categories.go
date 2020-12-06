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

type CategoryMsql struct {
	DB *gorm.DB
}

func NewMysqlCategoryRepository(Conn *gorm.DB) domain.CategoryRepository {
	return &CategoryMsql{Conn}
}

func (CatRepo *CategoryMsql) FindAll(ctx context.Context, skip int64, limmit int64, filter domain.Category) (res []domain.Category, err error) {
	if err = CatRepo.DB.Scopes(Paginate(ctx, skip, limmit, filter)).Find(&res).Error; err != nil {
		return nil, err
	}
	return
}

func (CatRepo *CategoryMsql) CountAll(ctx context.Context, skip int64, limmit int64, filter domain.Category) (res int64, err error) {
	if err = CatRepo.DB.WithContext(ctx).Model(&domain.Category{}).Where(filter).Count(&res).Error; err != nil {
		return 0, err
	}
	return
}

func (CatRepo *CategoryMsql) GetByID(ctx context.Context, id int64) (art domain.Category, err error) {
	ID := uint(id)
	category := domain.Category{}
	artc := CatRepo.DB.WithContext(ctx).Preload(clause.Associations).Preload("Articles.Creator.Role").Preload("Articles.Categories").Where(domain.Category{ID: ID}).Find(&category)
	if err = artc.Error; err != nil {
		return art, err
	}
	art = category
	return
}

func (CatRepo *CategoryMsql) Store(ctx context.Context, art *domain.Category) (err error) {
	if err = CatRepo.DB.Create(&art).Error; err != nil {
		return err
	}
	return
}

func (CatRepo *CategoryMsql) Update(ctx context.Context, id int64, artc *domain.Category) (err error) {
	var category domain.Category
	CatRepo.DB.WithContext(ctx).First(&category, id)
	checkart, _ := CatRepo.GetByID(ctx, id)
	if checkart.ID == 0 {
		errMessage := fmt.Sprintf("category with id %s  not exisit", string(id))
		errors.New(errMessage)
	}
	result := CatRepo.DB.WithContext(ctx).Model(&category).Updates(artc)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (CatRepo *CategoryMsql) DeleteByID(ctx context.Context, id int64) (err error) {
	artc := CatRepo.DB.WithContext(ctx).Where("id = ?", id).Delete(&domain.Category{})
	if err = artc.Error; err != nil {
		logrus.Error("err", err)
		return err
	}
	return
}
