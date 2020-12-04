package usecase

import (
	"context"
	"time"

	"github.com/geraldsamosir/myblogs/domain"
)

type articleUsecase struct {
	articleRepo    domain.ArticleRepository
	contextTimeout time.Duration
}

func NewArticleUsecase(art domain.ArticleRepository, timeout time.Duration) domain.ArticleUsecase {
	return &articleUsecase{
		articleRepo:    art,
		contextTimeout: timeout,
	}
}

func (art *articleUsecase) FindAll(c context.Context, page int64, limmit int64, filter domain.Article) (res []domain.Article, err error) {
	skip := page
	if skip <= 0 {
		skip = 0
	} else {
		skip = skip*limmit - 1
	}

	if limmit == 0 {
		limmit = 10
	}

	ctx, cancel := context.WithTimeout(c, art.contextTimeout)
	defer cancel()

	res, err = art.articleRepo.FindAll(ctx, skip, limmit, filter)
	if err != nil {
		return nil, err
	}
	return
}

func (art *articleUsecase) CountAll(c context.Context, skip int64, limmit int64, filter domain.Article) (res int64, err error) {
	ctx, cancel := context.WithTimeout(c, art.contextTimeout)
	defer cancel()
	res, err = art.articleRepo.CountAll(ctx, skip, limmit, filter)
	if err != nil {
		return 0, err
	}
	return

}
