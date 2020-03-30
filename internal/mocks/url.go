package mocks

import (
	"context"
	"github.com/xZ4PH0Dx/url_shortener"
)

type URLRepository struct {
	CreateFn func(ctx context.Context, u *url_shortener.URL) error
	ByCodeFn func(ctx context.Context, code string) (url_shortener.URL, error)
	ByIDFn   func(ctx context.Context, id int) (url_shortener.URL, error)
}

func (r *URLRepository) Create(ctx context.Context, u *url_shortener.URL) error {
	return r.CreateFn(ctx, u)
}

func (r *URLRepository) ByCode(ctx context.Context, code string) (url_shortener.URL, error) {
	return r.ByCodeFn(ctx, code)
}

func (r *URLRepository) ByID(ctx context.Context, id int) (url_shortener.URL, error) {
	return r.ByIDFn(ctx, id)
}
