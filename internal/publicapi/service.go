package publicapi

import (
	"context"
	"url_shortener"
)

func NewApiService(repo url_shortener.UrlRepository) url_shortener.PublicAPIService {
	return &service{urlRepo: repo}
}

type service struct {
	urlRepo url_shortener.UrlRepository
}

func (a *service) CreateUrl(ctx context.Context, u url_shortener.Url) (int, error) {
	err := a.urlRepo.Create(ctx, &u)
	if err != nil {
		return -1, err
	}
	return u.ID, err
}

func (a *service) GetById(ctx context.Context, i int) (u url_shortener.Url, err error) {
	u, err = a.urlRepo.ById(ctx, i)
	return
}

func (a *service) GetByCode(ctx context.Context, code string) (u url_shortener.Url, err error) {
	u, err = a.urlRepo.ByCode(ctx, code)
	return
}
