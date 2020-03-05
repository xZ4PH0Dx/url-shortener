package url_shortener

import "context"

// Url represents ....
type Url struct {
	ID   int    `json:"id"`
	Url  string `json:"url"`
	Code string `json:"code"`
}

type UrlRepository interface {
	Create(ctx context.Context, u *Url) error
	ById(ctx context.Context, i int) (Url, error)
	ByCode(ctx context.Context, s string) (Url, error)
}
