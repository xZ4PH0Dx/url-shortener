package publicapi_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/xZ4PH0Dx/url_shortener"
	"github.com/xZ4PH0Dx/url_shortener/internal/mocks"
	"github.com/xZ4PH0Dx/url_shortener/internal/publicapi"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

//var (
//host       = "localhost"
//dbPort     = 5432
//dbUser     = "postgres"
//dbPassword = "example"
//dbName     = "postgres"
//psqlInfo   = fmt.Sprintf("host=%s port=%d user=%s "+
//	"password=%s dbname=%s sslmode=disable",
//	host, dbPort, dbUser, dbPassword, dbName)
//)

var (
	u = url_shortener.Url{
		ID:   1,
		Url:  "http://google.com",
		Code: "so1gFSl5",
	}
)

func TestService_GetById(t *testing.T) {
	var testUrl url_shortener.Url

	type args struct {
		x interface{}
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "success",
			args: args{
				x: 1,
			},
			want: u,
		},
		{
			name: "success nil",
			args: args{
				x: 2,
			},
			want: url_shortener.Url{},
		},
	}

	db := &mocks.URLRepository{
		ByIdFn: func(ctx context.Context, id int) (url url_shortener.Url, err error) {
			if id == 1 {
				url = url_shortener.Url{
					ID:   1,
					Url:  "http://google.com",
					Code: "so1gFSl5",
				}
			}
			return
		},
	}

	r := publicapi.NewRouter(publicapi.NewApiService(db)).Handler()
	srv := httptest.NewServer(r)
	defer srv.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var strId string
			switch v := tt.args.x.(type) {
			case int:
				strId = strconv.Itoa(v)
			case string:
				strId = v
			}
			uri := "/urls/" + strId
			resp, err := http.Get(srv.URL + uri)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			err = json.NewDecoder(resp.Body).Decode(&testUrl)
			if err != nil {
				t.Error(err)
			}
			if testUrl != tt.want {
				t.Errorf("GetById() = %v, want %v", testUrl, tt.want)
			}
		})
	}
	//assert.Equal(t, u, testUrl)
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
