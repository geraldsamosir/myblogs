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

type RoleMsql struct {
	DB *gorm.DB
}

func NewMysqlRoleRepository(Conn *gorm.DB) domain.RoleRepository {
	return &RoleMsql{Conn}
}

func (RoleRepo *RoleMsql) FindAll(ctx context.Context, skip int64, limmit int64, filter domain.Role) (res []domain.Role, err error) {
	if err = RoleRepo.DB.Scopes(Paginate(ctx, skip, limmit, filter)).Find(&res).Error; err != nil {
		return nil, err
	}
	return
}

func (RoleRepo *RoleMsql) CountAll(ctx context.Context, skip int64, limmit int64, filter domain.Role) (res int64, err error) {
	if err = RoleRepo.DB.WithContext(ctx).Model(&domain.Role{}).Where(filter).Count(&res).Error; err != nil {
		return 0, err
	}
	return
}

func (RoleRepo *RoleMsql) GetByID(ctx context.Context, id int64) (art domain.Role, err error) {
	ID := uint(id)
	Role := domain.Role{}
	artc := RoleRepo.DB.WithContext(ctx).Preload(clause.Associations).Preload("Users.Role").Where(domain.Role{ID: ID}).Find(&Role)
	if err = artc.Error; err != nil {
		return art, err
	}
	art = Role
	return
}

func (RoleRepo *RoleMsql) Store(ctx context.Context, art *domain.Role) (err error) {
	if err = RoleRepo.DB.Create(&art).Error; err != nil {
		return err
	}
	return
}

func (RoleRepo *RoleMsql) Update(ctx context.Context, id int64, artc *domain.Role) (err error) {
	var Role domain.Role
	RoleRepo.DB.WithContext(ctx).First(&Role, id)
	checkart, _ := RoleRepo.GetByID(ctx, id)
	if checkart.ID == 0 {
		errMessage := fmt.Sprintf("Role with id %s  not exisit", string(id))
		errors.New(errMessage)
	}
	result := RoleRepo.DB.WithContext(ctx).Model(&Role).Updates(artc)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (RoleRepo *RoleMsql) DeleteByID(ctx context.Context, id int64) (err error) {
	artc := RoleRepo.DB.WithContext(ctx).Where("id = ?", id).Delete(&domain.Role{})
	if err = artc.Error; err != nil {
		logrus.Error("err", err)
		return err
	}
	return
}
