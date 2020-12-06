package usecase

import (
	"context"
	"time"

	"github.com/geraldsamosir/myblogs/domain"
)

type categoryUsecase struct {
	CategoryRepo   domain.CategoryRepository
	contextTimeout time.Duration
}

func NewcategoryUsecase(cat domain.CategoryRepository, timeout time.Duration) domain.CategoryUsecase {
	return &categoryUsecase{
		CategoryRepo:   cat,
		contextTimeout: timeout,
	}
}

//this method run when web server not strating
func init() {
	// log.Println("initiallize")
}

func (cat *categoryUsecase) FindAll(c context.Context, page int64, limmit int64, filter domain.Category) (res []domain.Category, err error) {
	ctx, cancel := context.WithTimeout(c, cat.contextTimeout)
	defer cancel()

	res, err = cat.CategoryRepo.FindAll(ctx, page, limmit, filter)
	if err != nil {
		return nil, err
	}
	return
}

func (cat *categoryUsecase) CountPage(c context.Context, skip int64, limmit int64, filter domain.Category) (res int64, err error) {
	ctx, cancel := context.WithTimeout(c, cat.contextTimeout)
	defer cancel()
	if limmit <= 0 {
		limmit = 10
	}
	countAll, err := cat.CategoryRepo.CountAll(ctx, skip, limmit, filter)
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

func (cat *categoryUsecase) GetByID(c context.Context, id int64) (Category domain.Category, err error) {
	ctx, cancel := context.WithTimeout(c, cat.contextTimeout)
	defer cancel()

	Category, err = cat.CategoryRepo.GetByID(ctx, id)
	if err != nil {
		return Category, err
	}
	return Category, nil
}

func (cat *categoryUsecase) Create(c context.Context, catc *domain.Category) (err error) {
	ctx, cancel := context.WithTimeout(c, cat.contextTimeout)
	defer cancel()
	err = cat.CategoryRepo.Store(ctx, catc)
	if err != nil {
		return err
	}
	return

}

func (cat *categoryUsecase) Update(ctx context.Context, id int64, catc *domain.Category) (err error) {
	ctx, cancel := context.WithTimeout(ctx, cat.contextTimeout)
	defer cancel()
	err = cat.CategoryRepo.Update(ctx, id, catc)
	if err != nil {
		return err
	}
	return
}

func (cat *categoryUsecase) DeleteByID(c context.Context, id int64) (message string, err error) {
	ctx, cancel := context.WithTimeout(c, cat.contextTimeout)
	defer cancel()

	err = cat.CategoryRepo.DeleteByID(ctx, id)
	if err != nil {
		return "", err
	}
	message = "success delete Category "
	return message, err
}
