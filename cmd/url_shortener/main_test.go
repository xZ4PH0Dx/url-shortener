package main

var (
	shortUrl string
)

const (
	originalUrl = "http://asd.com"
)



//func TestPostShorten(t *testing.T) {
//	data := url.Values{}
//	data.Set("url", originalUrl)
//	w := httptest.NewRecorder()
//
//	r, _ := http.NewRequest("POST", "/shorten", strings.NewReader(data.Encode()))
//	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
//	router.ServeHTTP(w, r)
//
//	doc, err := goquery.NewDocumentFromReader(w.Body)
//	if err != nil {
//		t.Fatalf("Failed to parse reponse body: %v", w.Body)
//	}
//
//	shortUrl = doc.Find("div#short_url").Text()
//	parsedUrl, err := url.Parse(shortUrl)
//	if err != nil {
//		t.Fatalf("Failed to parse url: %v", shortUrl)
//	}
//	shortCode := path.Base(parsedUrl.Path)
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, len(shortCode), 8)
//
//}
//
//func TestGetOriginalURL(t *testing.T) {
//	w := httptest.NewRecorder()
//	r, _ := http.NewRequest("GET", shortUrl, nil)
//	router.ServeHTTP(w, r)
//	assert.Equal(t, 301, w.Code)
//	assert.Equal(t, originalUrl, strings.Trim(w.Header().Get("Location"), " "))
//}
