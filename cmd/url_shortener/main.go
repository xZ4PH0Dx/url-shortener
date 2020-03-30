package main

//var (
//localIp  = GetLocalIP()
//router *gin.Engine
//appPort    = os.Getenv("APP_PORT")
//host       = os.Getenv("HOST")
//dbPort     = os.Getenv("DB_PORT")
//dbUser     = os.Getenv("DB_USER")
//dbPassword = os.Getenv("DB_PASSWORD")
//dbName     = os.Getenv("DB_NAME")

//)

func main() {
	//fmt.Println("APP_PORT", ":8080")
	//r := setupRouter()
	//_ = r.Run(appPort)
}

//func setupRouter() *gin.Engine {
//	router = gin.Default()
//	router.LoadHTMLGlob(filepath.Join("web", "*.tmpl"))
//
//	router.GET("/", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
//	})
//
//	router.GET("/byid/:id_url", func(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id_url"))
//		checkErr(err)
//
//		dbClient := pg.Client{}
//		defer dbClient.Close()
//		err = dbClient.Open(psqlInfo)
//		checkErr(err)
//		urlRepo := pg.NewSQLUrlRepo(dbClient.DB)
//		u, err := urlRepo.ByID(c, id)
//		c.Redirect(http.StatusMovedPermanently, u.URL)
//	})
//
//	router.GET("/u/:code", func(c *gin.Context) {
//		code := c.Param("code")
//		dbClient := pg.Client{}
//		defer dbClient.Close()
//		err := dbClient.Open(psqlInfo)
//		checkErr(err)
//		urlRepo := pg.NewSQLUrlRepo(dbClient.DB)
//		u, err := urlRepo.ByCode(c, code)
//		fmt.Println("URL TO REDIRECT", u.URL)
//		c.Redirect(http.StatusMovedPermanently, u.URL)
//	})
//	router.POST("/shorten", func(c *gin.Context) {
//		origUrl := c.Request.PostFormValue("url")
//		shortUri := generateCode([]byte(origUrl))
//		shortlink := "http://localhost" + appPort + "/u/" + shortUri
//
//		u := pg.URL{
//			URL:  origUrl,
//			Code: shortUri,
//		}
//
//		dbClient := pg.Client{}
//		defer dbClient.Close()
//		err := dbClient.Open(psqlInfo)
//		checkErr(err)
//
//		urlRepo := pg.NewSQLUrlRepo(dbClient.DB)
//		_, err = urlRepo.Create(c, u)
//		checkErr(err)
//
//		c.HTML(http.StatusOK, "index.tmpl", gin.H{"short_url": shortlink})
//	})
//	return router
//}
//
//func generateCode(b []byte) string {
//	var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//	s := make([]rune, 8)
//	rand.Seed(time.Now().UnixNano())
//	for i := range s {
//		s[i] = chars[rand.Intn(len(chars))]
//	}
//	return string(s)
//}

//func GetLocalIP() string {
//	addrs, err := net.InterfaceAddrs()
//	if err != nil {
//		return ""
//	}
//	for _, address := range addrs {
//		// check the address type and if it is not a loopback the display it
//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//			if ipnet.IP.To4() != nil {
//				return ipnet.IP.String()
//			}
//		}
//	}
//	return ""
//}

//func checkErr(e error) {
//	if e != nil {
//		log.Fatal(e)
//		return
//	}
//}
