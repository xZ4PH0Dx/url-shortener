package pg

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/xZ4PH0Dx/url_shortener"
)

func NewSQLUrlRepo(conn *sqlx.DB) url_shortener.URLRepository {
	return &pgURLRepo{Conn: conn,}
}

type pgURLRepo struct {
	Conn *sqlx.DB
}

func (p *pgURLRepo) Create(ctx context.Context, u *url_shortener.URL) error {
	urlCreate := "INSERT INTO urls(code, original_url) VALUES ($1, $2) RETURNING ID;"

	return p.Conn.QueryRowContext(
		context.Background(),
		urlCreate,
		u.Code,
		u.URL,
	).Scan(&u.ID)
}

func (p *pgURLRepo) ByID(ctx context.Context, i int) (url_shortener.URL, error) {
	urlByID := "SELECT id, TRIM(code) code, TRIM(original_url) original_url FROM urls WHERE id = $1;"
	u := url_shortener.URL{}
	err := p.Conn.QueryRowContext(
		ctx,
		urlByID,
		i,
	).Scan(&u.ID, &u.Code, &u.URL)

	return u, err
}

func (p *pgURLRepo) ByCode(ctx context.Context, code string) (url_shortener.URL, error) {
	urlByCode := "SELECT id, TRIM(code) code, TRIM(original_url) original_url FROM urls WHERE code = $1;"
	u := url_shortener.URL{}
	err := p.Conn.QueryRowContext(
		ctx,
		urlByCode,
		code,
	).Scan(&u.ID, &u.Code, &u.URL)

	return u, err
}
