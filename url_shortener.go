package url_shortener

import (
	"context"
)

// URL represents ....
type URL struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Code string `json:"code"`
}

type URLRepository interface {
	Create(ctx context.Context, u *URL) error
	ByID(ctx context.Context, i int) (URL, error)
	ByCode(ctx context.Context, s string) (URL, error)
}
type PublicAPIService interface {
	CreateURL(ctx context.Context, u URL) (URL, error)
	GetByID(ctx context.Context, i int) (URL, error)
	GetByCode(ctx context.Context, code string) (URL, error)
}
