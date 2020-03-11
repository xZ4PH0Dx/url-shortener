package publicapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"url_shortener"
	"url_shortener/internal/pg"
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
	respRec = httptest.NewRecorder()
	u       = url_shortener.Url{
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
	r := NewRouter(NewApiService(db))
	r.InitializeRoutes()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/url/1", nil)
	r.router.ServeHTTP(respRec, req)
	//response := executeRequest(req)
	err = json.NewDecoder(respRec.Body).Decode(&testUrl)

	assert.Equal(t, u, testUrl)
}

func TestService_CreateUrl(t *testing.T) {
	dbClient := pg.NewClient()
	err := dbClient.Open(psqlInfo)
	if err != nil {
		t.Error(err)
	}
	dropUrlTable(dbClient.DB)
	dbClient.InitSchema()
	defer dbClient.Close()

	db := pg.NewSQLUrlRepo(dbClient.DB)
	r := NewRouter(NewApiService(db))
	r.InitializeRoutes()

	mUrl, err := json.Marshal(u)
	if err != nil {
		t.Error(err)
	}
	b := bytes.NewBuffer(mUrl)

	req, _ := http.NewRequestWithContext(context.Background(), "POST", "/create", b)
	r.router.ServeHTTP(respRec, req)
	i, err := strconv.Atoi(respRec.Body.String())
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, u.ID, i)
}

func dropUrlTable(db *sqlx.DB) {
	dropUrlTable := "DROP TABLE IF EXISTS urls "
	_ = db.QueryRow(
		dropUrlTable,
	)
}
