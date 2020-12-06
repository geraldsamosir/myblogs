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

//this method run when web server not strating
func init() {
	// logrus.Info("initiallize")
}

func (art *articleUsecase) FindAll(c context.Context, page int64, limmit int64, filter domain.Article) (res []domain.Article, err error) {
	ctx, cancel := context.WithTimeout(c, art.contextTimeout)
	defer cancel()

	res, err = art.articleRepo.FindAll(ctx, page, limmit, filter)
	if err != nil {
		return nil, err
	}
	return
}

func (art *articleUsecase) CountPage(c context.Context, skip int64, limmit int64, filter domain.Article) (res int64, err error) {
	ctx, cancel := context.WithTimeout(c, art.contextTimeout)
	defer cancel()
	if limmit <= 0 {
		limmit = 10
	}
	countAll, err := art.articleRepo.CountAll(ctx, skip, limmit, filter)
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

func (art *articleUsecase) GetByID(c context.Context, id int64) (article domain.Article, err error) {
	ctx, cancel := context.WithTimeout(c, art.contextTimeout)
	defer cancel()

	article, err = art.articleRepo.GetByID(ctx, id)
	if err != nil {
		return article, err
	}
	return article, nil
}

func (art *articleUsecase) Create(c context.Context, artc domain.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, art.contextTimeout)
	defer cancel()
	err = art.articleRepo.Store(ctx, &artc)
	if err != nil {
		return err
	}
	return

}

func (art *articleUsecase) Update(ctx context.Context, id int64, artc domain.Article) (err error) {
	ctx, cancel := context.WithTimeout(ctx, art.contextTimeout)
	defer cancel()
	if artc.Title != "" {
		currentTime := time.Now().Format("01-02-2006")
		artc.Slug = artc.Title + "-" + currentTime
	}
	err = art.articleRepo.Update(ctx, id, &artc)
	if err != nil {
		return err
	}
	return
}

func (art *articleUsecase) DeleteByID(c context.Context, id int64) (message string, err error) {
	ctx, cancel := context.WithTimeout(c, art.contextTimeout)
	defer cancel()

	err = art.articleRepo.DeleteByID(ctx, id)
	if err != nil {
		return "", err
	}
	message = "success delete article "
	return message, err
}
