package publicapi_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/xZ4PH0Dx/url_shortener"
	"github.com/xZ4PH0Dx/url_shortener/internal/mocks"
	"github.com/xZ4PH0Dx/url_shortener/internal/publicapi"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_GetById(t *testing.T) {
	tests := []struct {
		name     string
		urlID    string
		db       *mocks.URLRepository
		respBody string
	}{
		{
			name:  "success",
			urlID: "1",
			db: &mocks.URLRepository{
				ByIdFn: func(ctx context.Context, id int) (url url_shortener.Url, err error) {
					return url_shortener.Url{
						ID:   1,
						Url:  "http://google.com",
						Code: "so1gFSl5",
					}, nil
				},
			},
			respBody: `{"id":1,"url":"http://google.com","code":"so1gFSl5"}` + "\n",
		},
		{
			name:  "not found",
			urlID: "2",
			db: &mocks.URLRepository{
				ByIdFn: func(ctx context.Context, id int) (url url_shortener.Url, err error) {
					return
				},
			},
			respBody: `{"id":0,"url":"","code":""}` + "\n", //TODO very strange
		},
		{
			name:  "bad id format",
			urlID: "bad_id",
			db: &mocks.URLRepository{
				ByIdFn: func(ctx context.Context, id int) (url url_shortener.Url, err error) {
					return
				},
			},
			respBody: "404 page not found\n", //TODO very strange
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := publicapi.NewRouter(publicapi.NewApiService(tt.db)).Handler()
			srv := httptest.NewServer(r)
			defer srv.Close()

			resp, err := http.Get(fmt.Sprintf("%s%s%s", srv.URL, "/urls/", tt.urlID))
			if err != nil {
				t.Fatal(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}
			resp.Body.Close()

			if diff := cmp.Diff(tt.respBody, string(body)); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestService_CreateUrl(t *testing.T) {
	tests := []struct {
		name     string
		urlID    string
		url      *url_shortener.Url
		db       *mocks.URLRepository
		respBody string
	}{
		{
			name:  "success",
			urlID: "1",
			url: &url_shortener.Url{
				ID:   1,
				Url:  "http://google.com",
				Code: "so1gFSl5",
			},
			db: &mocks.URLRepository{
				ByIdFn: func(ctx context.Context, id int) (url url_shortener.Url, err error) {
					return url_shortener.Url{
						ID:   1,
						Url:  "http://google.com",
						Code: "so1gFSl5",
					}, nil
				},
				CreateFn: func(ctx context.Context, u *url_shortener.Url) error {
					return nil
				},
			},
			respBody: `{"id":1,"url":"http://google.com","code":"so1gFSl5"}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := publicapi.NewRouter(publicapi.NewApiService(tt.db)).Handler()
			srv := httptest.NewServer(r)
			defer srv.Close()

			mUrl, err := json.Marshal(tt.url)
			if err != nil {
				t.Error(err)
			}
			b := bytes.NewBuffer(mUrl)

			_, err = http.Post(fmt.Sprintf("%s%s", srv.URL, "/urls"), "application/json", b)
			if err != nil {
				t.Fatal(err)
			}
			resp, err := http.Get(fmt.Sprintf("%s%s%s", srv.URL, "/urls/", tt.urlID))
			if err != nil {
				t.Fatal(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}
			if err = resp.Body.Close(); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.respBody, string(body)); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
