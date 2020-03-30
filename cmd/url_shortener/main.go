package main

import (
	"fmt"
	"github.com/xZ4PH0Dx/url_shortener/internal/pg"
	"github.com/xZ4PH0Dx/url_shortener/internal/publicapi"
	"log"
	"net/http"
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

func main() {
	dbClient := pg.NewClient()
	defer dbClient.Close()

	err := dbClient.Open(psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	db := pg.NewSQLUrlRepo(dbClient.DB)
	r := publicapi.NewRouter(publicapi.NewAPIService(db)).Handler()

	log.Fatal(http.ListenAndServe(":8080", r))
}
