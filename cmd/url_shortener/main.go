package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"net"
	"net/http"
	"path/filepath"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
	appPort  = ":8080"
)

var (
	localIp  = GetLocalIP()
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)
var (
	router *gin.Engine
	db     *sql.DB
)

func main() {
	r := setupRouter()
	_ = r.Run(appPort)
}

func setupRouter() *gin.Engine {
	router = gin.Default()
	router.LoadHTMLGlob(filepath.Join("web", "*.tmpl"))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	router.GET("/u/:shorturi", func(c *gin.Context) {
		searchUri := c.Param("shorturi")
		var originalUrl string
		db, _ = sql.Open("postgres", psqlInfo)
		defer db.Close()
		err := db.QueryRow("SELECT original_url FROM urls WHERE code = $1", searchUri).Scan(&originalUrl)
		if err != nil {
			c.String(http.StatusBadRequest, "URL you entered not found")
			return
		}
		c.Redirect(http.StatusMovedPermanently, originalUrl)
	})
	router.POST("/shorten", func(c *gin.Context) {
		origUrl := c.Request.PostFormValue("url")
		shortUri := generateCode([]byte(origUrl))
		shortlink := "http://localhost" + appPort + "/u/" + shortUri
		execSql := "CALL insert_data($1, $2, $3)"
		db, _ = sql.Open("postgres", psqlInfo)
		defer db.Close()
		stmt, err := db.Prepare(execSql)
		checkErr(err)
		_, err = stmt.Exec(shortUri, origUrl, c.ClientIP())
		checkErr(err)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"short_url": shortlink})
	})
	return router
}

func generateCode(b []byte) string {
	var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, 8)
	rand.Seed(time.Now().UnixNano())
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
		return
	}
}
