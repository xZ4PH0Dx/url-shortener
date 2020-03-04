package pg

import (
	"context"
	"database/sql"
	"fmt"
)

func NewSQLUrlRepo(Conn *sql.DB) UrlRepository {
	return &pgUrlRepo{Conn: Conn,}
}

type pgUrlRepo struct {
	Conn *sql.DB
}

func (p *pgUrlRepo) Create(ctx context.Context, u Url) (int, error) {
	sqlStr := "INSERT INTO urls(code, original_url) VALUES ($1, $2);"

	err := p.Conn.QueryRowContext(
		ctx,
		sqlStr,
		u.Code,
		u.Url,
	).Scan(&u.ID)
	if err != nil {
		return fmt.Println(err.Error())
	}
	return u.ID, nil
}

func (p *pgUrlRepo) ById(ctx context.Context, i int) (Url, error) {
	sqlStr := "SELECT TRIM(id), TRIM(code), TRIM(original_url) FROM urls WHERE id = $1;"
	u := Url{}
	err := p.Conn.QueryRowContext(
		ctx,
		sqlStr,
		i,
	).Scan(&u.ID, &u.Url, &u.Code)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (p *pgUrlRepo) ByCode(ctx context.Context, s string) (Url, error) {
	sqlStr := "SELECT id, code, original_url FROM urls WHERE code = $1;"
	u := Url{}
	err := p.Conn.QueryRowContext(
		ctx,
		sqlStr,
		s,
	).Scan(&u.ID, &u.Url, &u.Code)
	if err != nil {
		return u, err
	}
	return u, nil
}
