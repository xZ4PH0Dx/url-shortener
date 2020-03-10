package publicapi

import (
	"context"
	"url_shortener"
)

func NewApiService() url_shortener.PublicAPIService {
	return &service{}
}

type service struct {
}

func (a *service) CreateUrl(ctx context.Context) error {
	//Extract Url from context.. and workup
	return nil
}

func (a *service) GetById(ctx context.Context) (u url_shortener.Url, err error) {

	return url_shortener.Url{}, nil
}

func (a *service) GetByCode(ctx context.Context) (u url_shortener.Url, err error) {
	return url_shortener.Url{}, nil
}
