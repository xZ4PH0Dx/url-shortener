package pg

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"url_shortener"
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
	id int
	u  = url_shortener.Url{
		Url:  "http://google.com",
		Code: "so1gFSl5",
	}
)

func init() {
	_ = os.Chdir("../..") //pwd для тестов ./url_shortener/cmd/url_shortener, для скомпилированного файла ./url_shortener:(
}

func TestCreate(t *testing.T) {
	dbClient := Client{}
	defer dbClient.Close()

	err := dbClient.Open(psqlInfo)
	if err != nil {
		t.Error(err)
	}

	db := NewSQLUrlRepo(dbClient.DB)
	db.Clear(context.Background())
	id, err = db.Create(context.Background(), &u)
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, u.ID)
}

func TestById(t *testing.T) {
	dbClient := Client{}

	err := dbClient.Open(psqlInfo)
	if err != nil {
		t.Error(err)
	}
	defer dbClient.Close()

	db := NewSQLUrlRepo(dbClient.DB)
	dbU, err := db.ById(context.Background(), u.ID)

	assert.Equal(t, dbU, u)

}

func TestByCode(t *testing.T) {
	dbClient := Client{}

	err := dbClient.Open(psqlInfo)
	if err != nil {
		t.Error(err)
	}
	defer dbClient.Close()

	db := NewSQLUrlRepo(dbClient.DB)
	dbU, err := db.ByCode(context.Background(), u.Code)

	assert.Equal(t, dbU, u)

}
