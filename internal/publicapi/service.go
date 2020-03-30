package publicapi

import (
	"context"
	"github.com/xZ4PH0Dx/url_shortener"
)

func NewAPIService(repo url_shortener.URLRepository) url_shortener.PublicAPIService {
	return &service{urlRepo: repo}
}

type service struct {
	urlRepo url_shortener.URLRepository
}

func (a *service) CreateURL(ctx context.Context, u url_shortener.URL) (url_shortener.URL, error) {
	err := a.urlRepo.Create(ctx, &u)
	return u, err
}

func (a *service) GetByID(ctx context.Context, i int) (u url_shortener.URL, err error) {
	return a.urlRepo.ByID(ctx, i)
}

func (a *service) GetByCode(ctx context.Context, code string) (u url_shortener.URL, err error) {
	return a.urlRepo.ByCode(ctx, code)
}
