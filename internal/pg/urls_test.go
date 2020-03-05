package pg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
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
	u = url_shortener.Url{
		ID:   1,
		Url:  "http://google.com",
		Code: "so1gFSl5",
	}
)

func init() {
	_ = os.Chdir("../..") //pwd для тестов ./url_shortener/cmd/url_shortener, для скомпилированного файла ./url_shortener:(
	dbClient := NewClient()
	err := dbClient.Open(psqlInfo)
	dbClient.InitSchema()
	if err != nil {
		log.Fatal(err)
	}
	defer dbClient.Close()
}

func TestCreate(t *testing.T) {
	dbClient := NewClient()
	err := dbClient.Open(psqlInfo)
	if err != nil {
		t.Error(err)
	}
	dropUrlTable(dbClient.DB)
	dbClient.InitSchema()
	defer dbClient.Close()
	db := NewSQLUrlRepo(dbClient.DB)
	err = db.Create(context.Background(), &u)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, u.ID)
}

func TestById(t *testing.T) {
	dbClient := NewClient()
	err := dbClient.Open(psqlInfo)
	if err != nil {
		t.Error(err)
	}
	dropUrlTable(dbClient.DB)
	dbClient.InitSchema()
	defer dbClient.Close()
	db := NewSQLUrlRepo(dbClient.DB)
	err = db.Create(context.Background(), &u)
	if err != nil {
		t.Error(err)
	}
	dbU, err := db.ById(context.Background(), u.ID)
	assert.Equal(t, dbU, u)

}

func TestByCode(t *testing.T) {
	dbClient := Client{}
	err := dbClient.Open(psqlInfo)
	if err != nil {
		t.Error(err)
	}
	dropUrlTable(dbClient.DB)
	dbClient.InitSchema()
	defer dbClient.Close()
	db := NewSQLUrlRepo(dbClient.DB)
	err = db.Create(context.Background(), &u)
	if err != nil {
		t.Error(err)
	}
	dbU, err := db.ByCode(context.Background(), u.Code)
	assert.Equal(t, dbU, u)
}

func dropUrlTable(db *sql.DB) {
	dropUrlTable := "DROP TABLE IF EXISTS urls "
	_ = db.QueryRow(
		dropUrlTable,
	)
}
