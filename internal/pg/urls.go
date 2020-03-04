package pg

import (
	"context"
	"database/sql"
	"url_shortener"
)

func NewSQLUrlRepo(Conn *sql.DB) url_shortener.UrlRepository {
	return &pgUrlRepo{Conn: Conn,}
}

type pgUrlRepo struct {
	Conn *sql.DB
}

func (p *pgUrlRepo) Clear(ctx context.Context) {
	sqlStr := "TRUNCATE TABLE urls"

	_ = p.Conn.QueryRowContext(
		ctx,
		sqlStr,
	)
}

func (p *pgUrlRepo) Create(ctx context.Context, u *url_shortener.Url) (int, error) {
	sqlStr := "INSERT INTO urls(code, original_url) VALUES ($1, $2) RETURNING ID;"
	err := p.Conn.QueryRowContext(
		context.Background(),
		sqlStr,
		u.Code,
		u.Url,
	).Scan(&u.ID)
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (p *pgUrlRepo) ById(ctx context.Context, i int) (url_shortener.Url, error) {
	sqlStr := "SELECT id, TRIM(code) code, TRIM(original_url) original_url FROM urls WHERE id = $1;"
	u := url_shortener.Url{}
	err := p.Conn.QueryRowContext(
		ctx,
		sqlStr,
		i,
	).Scan(&u.ID, &u.Code, &u.Url)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (p *pgUrlRepo) ByCode(ctx context.Context, s string) (url_shortener.Url, error) {
	sqlStr := "SELECT id, TRIM(code) code, TRIM(original_url) original_url FROM urls WHERE code = $1;"
	u := url_shortener.Url{}
	err := p.Conn.QueryRowContext(
		ctx,
		sqlStr,
		s,
	).Scan(&u.ID, &u.Code, &u.Url)
	if err != nil {
		return u, err
	}
	return u, nil
}
