package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/geraldsamosir/myblogs/helper"

	"github.com/geraldsamosir/myblogs/interface/webserver/middleware"

	"github.com/labstack/gommon/log"
)

type userUsecase struct {
	UserRepo         domain.UserRepository
	contextTimeout   time.Duration
	passwordHandling helper.Password
	Authentication   middleware.Auth
}

func NewUserUsecase(usr domain.UserRepository, timeout time.Duration, passwordHandling helper.Password, Authentication middleware.Auth) domain.UserUsecase {
	return &userUsecase{
		UserRepo:         usr,
		contextTimeout:   timeout,
		passwordHandling: passwordHandling,
		Authentication:   Authentication,
	}
}

//this method run when web server not strating
func init() {
	// logrus.Info("initiallize")
}

func (usr *userUsecase) FindAll(c context.Context, page int64, limmit int64, filter domain.UserResponse) (res []domain.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(c, usr.contextTimeout)
	defer cancel()

	res, err = usr.UserRepo.FindAll(ctx, page, limmit, filter)
	if err != nil {
		return nil, err
	}
	return
}

func (usr *userUsecase) CountPage(c context.Context, skip int64, limmit int64, filter domain.UserResponse) (res int64, err error) {
	ctx, cancel := context.WithTimeout(c, usr.contextTimeout)
	defer cancel()
	if limmit <= 0 {
		limmit = 10
	}
	countAll, err := usr.UserRepo.CountAll(ctx, skip, limmit, filter)
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

func (usr *userUsecase) GetByID(c context.Context, id int64) (User domain.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(c, usr.contextTimeout)
	defer cancel()

	User, err = usr.UserRepo.GetByID(ctx, id)
	if err != nil {
		return User, err
	}
	return User, nil
}

func (usr *userUsecase) Register(c context.Context, usrc domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, usr.contextTimeout)
	defer cancel()
	usrc.Password = usr.passwordHandling.HashAndSalt([]byte(usrc.Password))
	err = usr.UserRepo.Store(ctx, &usrc)
	if err != nil {
		return err
	}
	return

}
func (usr *userUsecase) Login(ctx context.Context, user domain.Authentication) (auth domain.AuthenticationResponse, err error, errMessage string) {
	var usrs domain.User
	usrs, err = usr.UserRepo.GetByUsername(ctx, user.UserName)
	if err != nil {
		return auth, err, "your username /password false"
	}
	isPasswordMatch := usr.passwordHandling.ComparePassword(usrs.Password, []byte(user.Password))
	if isPasswordMatch == false {
		return auth, errors.New("your username /password false "), "your username /password false "
	}

	myAuth := domain.AuthenticationResponse{
		UserResponse: domain.UserResponse{
			ID:        usrs.ID,
			UserName:  usrs.UserName,
			RoleID:    usrs.RoleID,
			FirstName: usrs.FirstName,
			LastName:  usrs.LastName,
			Role:      usrs.Role,
		},
		Jwt: "",
	}
	//generate token
	token, err := usr.Authentication.GenerateToken(&myAuth)
	if err != nil {
		return auth, err, "error generate token"
	}
	myAuth.Jwt = token
	auth = myAuth
	return auth, nil, "success"
}
func (usr *userUsecase) Update(ctx context.Context, id int64, usrc domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, usr.contextTimeout)
	defer cancel()

	log.Info("pass", usrc.Password)
	if usrc.Password != "" {
		log.Info("pass", usrc.Password)
		usrc.Password = usr.passwordHandling.HashAndSalt([]byte(usrc.Password))
	}
	err = usr.UserRepo.Update(ctx, id, &usrc)
	if err != nil {
		return err
	}
	return
}

func (usr *userUsecase) DeleteByID(c context.Context, id int64) (message string, err error) {
	ctx, cancel := context.WithTimeout(c, usr.contextTimeout)
	defer cancel()

	err = usr.UserRepo.DeleteByID(ctx, id)
	if err != nil {
		return "", err
	}
	message = "success delete User "
	return message, err
}
