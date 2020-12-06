package usecase

import (
	"context"
	"time"

	"github.com/geraldsamosir/myblogs/domain"
)

type roleUsecase struct {
	roleRepo       domain.RoleRepository
	contextTimeout time.Duration
}

func NewRoleUsecase(rol domain.RoleRepository, timeout time.Duration) domain.RoleUsecase {
	return &roleUsecase{
		roleRepo:       rol,
		contextTimeout: timeout,
	}
}

//this method run when web server not strating
func init() {
	// logrus.Info("initiallize")
}

func (rol *roleUsecase) FindAll(c context.Context, page int64, limmit int64, filter domain.Role) (res []domain.Role, err error) {
	ctx, cancel := context.WithTimeout(c, rol.contextTimeout)
	defer cancel()

	res, err = rol.roleRepo.FindAll(ctx, page, limmit, filter)
	if err != nil {
		return nil, err
	}
	return
}

func (rol *roleUsecase) CountPage(c context.Context, skip int64, limmit int64, filter domain.Role) (res int64, err error) {
	ctx, cancel := context.WithTimeout(c, rol.contextTimeout)
	defer cancel()
	if limmit <= 0 {
		limmit = 10
	}
	countAll, err := rol.roleRepo.CountAll(ctx, skip, limmit, filter)
	if err != nil {
		return 0, err
	}
	if (countAll / limmit) == 0 {
		res = 1
	} else {
		res = countAll / limmit
	}
	return

}

func (rol *roleUsecase) GetByID(c context.Context, id int64) (role domain.Role, err error) {
	ctx, cancel := context.WithTimeout(c, rol.contextTimeout)
	defer cancel()

	role, err = rol.roleRepo.GetByID(ctx, id)
	if err != nil {
		return role, err
	}
	return role, nil
}

func (rol *roleUsecase) Create(c context.Context, rolc *domain.Role) (err error) {
	ctx, cancel := context.WithTimeout(c, rol.contextTimeout)
	defer cancel()
	err = rol.roleRepo.Store(ctx, rolc)
	if err != nil {
		return err
	}
	return

}

func (rol *roleUsecase) Update(ctx context.Context, id int64, rolc *domain.Role) (err error) {
	ctx, cancel := context.WithTimeout(ctx, rol.contextTimeout)
	defer cancel()
	err = rol.roleRepo.Update(ctx, id, rolc)
	if err != nil {
		return err
	}
	return
}

func (rol *roleUsecase) DeleteByID(c context.Context, id int64) (message string, err error) {
	ctx, cancel := context.WithTimeout(c, rol.contextTimeout)
	defer cancel()

	err = rol.roleRepo.DeleteByID(ctx, id)
	if err != nil {
		return "", err
	}
	message = "success delete role "
	return message, err
}
