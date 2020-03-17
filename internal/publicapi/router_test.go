package publicapi_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"url_shortener"
	"url_shortener/internal/mocks"
	"url_shortener/internal/pg"
	"url_shortener/internal/publicapi"
)

var (
	host       = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "example"
	dbName     = "postgres"
	psqlInfo   = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, dbPort, dbUser, dbPassword, dbName)
)

var (
	u = url_shortener.Url{
		ID:   1,
		Url:  "http://google.com",
		Code: "so1gFSl5",
	}
)

func TestService_GetById(t *testing.T) {
	var testUrl url_shortener.Url
	dbClient := pg.NewClient()
	err := dbClient.Open(psqlInfo)
	if err != nil {
		t.Error(err)
	}
	dropUrlTable(dbClient.DB)
	dbClient.InitSchema()
	defer dbClient.Close()
	db := pg.NewSQLUrlRepo(dbClient.DB)
	err = db.Create(context.Background(), &u)
	if err != nil {
		t.Error(err)
	}
	r := publicapi.NewRouter(publicapi.NewApiService(db)).Handler()
	srv := httptest.NewServer(r)
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/urls/1")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&testUrl)

	assert.Equal(t, u, testUrl)
}

func TestService_CreateUrl(t *testing.T) {
	var testUrl url_shortener.Url
	dbClient := pg.NewClient()
	err := dbClient.Open(psqlInfo)
	if err != nil {
		t.Error(err)
	}
	dropUrlTable(dbClient.DB)
	dbClient.InitSchema()
	defer dbClient.Close()
	db := pg.NewSQLUrlRepo(dbClient.DB)
	r := publicapi.NewRouter(publicapi.NewApiService(db)).Handler()
	srv := httptest.NewServer(r)
	defer srv.Close()

	mUrl, err := json.Marshal(u)
	if err != nil {
		t.Error(err)
	}
	b := bytes.NewBuffer(mUrl)

	resp, err := http.Post(srv.URL+"/urls", "application/json", b)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&testUrl)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, u.ID, testUrl.ID)
}

func dropUrlTable(db *sqlx.DB) {
	dropUrlTable := "DROP TABLE IF EXISTS urls "
	_ = db.QueryRow(
		dropUrlTable,
	)
}

func TestService_MockCreateUrl(t *testing.T) {
	var testUrl url_shortener.Url
	db := &mocks.URLRepository{
		CreateFn: func(ctx context.Context, u *url_shortener.Url) error {
			u = &url_shortener.Url{
				ID:   1,
				Url:  u.Url,
				Code: u.Code,
			}
			return nil
		},
	}
	r := publicapi.NewRouter(publicapi.NewApiService(db)).Handler()
	srv := httptest.NewServer(r)
	defer srv.Close()

	mUrl, err := json.Marshal(u)
	if err != nil {
		t.Error(err)
	}
	b := bytes.NewBuffer(mUrl)

	resp, err := http.Post(srv.URL+"/urls", "application/json", b)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&testUrl)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, u.ID, testUrl.ID)
}
