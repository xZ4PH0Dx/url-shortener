package publicapi_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"url_shortener"
	"url_shortener/internal/mocks"
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

	db := &mocks.URLRepository{
		ByIdFn: func(ctx context.Context, id int) (url url_shortener.Url, err error) {
			url = url_shortener.Url{
				ID:   1,
				Url:  "http://google.com",
				Code: "so1gFSl5",
			}
			return
		},
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
	db := &mocks.URLRepository{
		CreateFn: func(ctx context.Context, u *url_shortener.Url) error {
			u.ID = 1
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
