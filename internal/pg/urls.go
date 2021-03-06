package pg

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/xZ4PH0Dx/url_shortener"
)

func NewSQLUrlRepo(Conn *sqlx.DB) url_shortener.UrlRepository {
	return &pgUrlRepo{Conn: Conn,}
}

type pgUrlRepo struct {
	Conn *sqlx.DB
}

func (p *pgUrlRepo) Create(ctx context.Context, u *url_shortener.Url) error {
	urlCreate := "INSERT INTO urls(code, original_url) VALUES ($1, $2) RETURNING ID;"
	return p.Conn.QueryRowContext(
		context.Background(),
		urlCreate,
		u.Code,
		u.Url,
	).Scan(&u.ID)
}

func (p *pgUrlRepo) ById(ctx context.Context, i int) (url_shortener.Url, error) {
	urlById := "SELECT id, TRIM(code) code, TRIM(original_url) original_url FROM urls WHERE id = $1;"
	u := url_shortener.Url{}
	err := p.Conn.QueryRowContext(
		ctx,
		urlById,
		i,
	).Scan(&u.ID, &u.Code, &u.Url)
	return u, err
}

func (p *pgUrlRepo) ByCode(ctx context.Context, code string) (url_shortener.Url, error) {
	urlByCode := "SELECT id, TRIM(code) code, TRIM(original_url) original_url FROM urls WHERE code = $1;"
	u := url_shortener.Url{}
	err := p.Conn.QueryRowContext(
		ctx,
		urlByCode,
		code,
	).Scan(&u.ID, &u.Code, &u.Url)
	return u, err
}
