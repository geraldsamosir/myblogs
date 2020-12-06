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

type UserMysql struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &UserMysql{Conn}
}

func (UsrRepo *UserMysql) FindAll(ctx context.Context, skip int64, limmit int64, filter domain.UserResponse) (res []domain.UserResponse, err error) {
	if err = UsrRepo.DB.Scopes(Paginate(ctx, skip, limmit, filter)).Preload(clause.Associations).Preload("Articles").Preload("Role").Find(&res).Error; err != nil {
		return nil, err
	}
	return
}

func (UsrRepo *UserMysql) CountAll(ctx context.Context, skip int64, limmit int64, filter domain.UserResponse) (res int64, err error) {
	if err = UsrRepo.DB.WithContext(ctx).Model(&domain.User{}).Where(filter).Count(&res).Error; err != nil {
		return 0, err
	}
	return
}

func (UsrRepo *UserMysql) GetByID(ctx context.Context, id int64) (usr domain.UserResponse, err error) {
	ID := uint(id)
	user := domain.UserResponse{}
	artc := UsrRepo.DB.WithContext(ctx).Preload(clause.Associations).Preload("Articles").Preload("Role").Where(domain.User{ID: ID}).Find(&user)
	if err = artc.Error; err != nil {
		return usr, err
	}
	usr = user
	return
}

func (UsrRepo *UserMysql) GetByUsername(ctx context.Context, userName string) (usr domain.User, err error) {
	user := domain.User{}
	artc := UsrRepo.DB.WithContext(ctx).Preload(clause.Associations).Preload("Role").Where(domain.User{UserName: userName}).Find(&user)
	if err = artc.Error; err != nil {
		return usr, err
	}
	usr = user
	return
}
func (UsrRepo *UserMysql) Store(ctx context.Context, art *domain.User) (err error) {
	if err = UsrRepo.DB.Create(&art).Preload(clause.Associations).Preload("Articles").Preload("Role").Error; err != nil {
		return err
	}
	return
}

func (UsrRepo *UserMysql) Update(ctx context.Context, id int64, artc *domain.User) (err error) {
	var user domain.User
	UsrRepo.DB.WithContext(ctx).First(&user, id)
	checkart, _ := UsrRepo.GetByID(ctx, id)
	if checkart.ID == 0 {
		errMessage := fmt.Sprintf("user with id %s  not exisit", string(id))
		errors.New(errMessage)
	}
	result := UsrRepo.DB.WithContext(ctx).Model(&user).Updates(artc)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (UsrRepo *UserMysql) DeleteByID(ctx context.Context, id int64) (err error) {
	artc := UsrRepo.DB.WithContext(ctx).Where("id = ?", id).Delete(&domain.User{})
	if err = artc.Error; err != nil {
		logrus.Error("err", err)
		return err
	}
	return
}
