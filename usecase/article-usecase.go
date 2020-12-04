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

func (art *articleUsecase) FindAll(c context.Context, skip int64, limmit int64) (res []domain.Article, err error) {
	if limmit == 0 {
		limmit = 10
	}

	ctx, cancel := context.WithTimeout(c, art.contextTimeout)
	defer cancel()

	res, err = art.articleRepo.FindAll(ctx, skip, limmit)
	if err != nil {
		return nil, err
	}
	return
}
