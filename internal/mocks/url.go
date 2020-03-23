package mocks

import (
	"context"
	"url_shortener"
)

type URLRepository struct {
	CreateFn func(ctx context.Context, u *url_shortener.Url) error
	ByCodeFn func(ctx context.Context, code string) (url_shortener.Url, error)
	ByIdFn   func(ctx context.Context, id int) (url_shortener.Url, error)
}

func (r *URLRepository) Create(ctx context.Context, u *url_shortener.Url) error {
	return r.CreateFn(ctx, u)
}

func (r *URLRepository) ByCode(ctx context.Context, code string) (url_shortener.Url, error) {
	return r.ByCodeFn(ctx, code)
}

func (r *URLRepository) ById(ctx context.Context, id int) (url_shortener.Url, error) {
	return r.ByIdFn(ctx, id)
}
