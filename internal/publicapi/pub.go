package publicapi
//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"github.com/xZ4PH0Dx/url_shortener"
//	"github.com/xZ4PH0Dx/url_shortener/internal/mocks"
//	"io/ioutil"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestService_GetById(t *testing.T) {
//	tests := []struct {
//		name     string
//		urlID    string
//		db       *mocks.URLRepository
//		respBody string
//	}{
//		{
//			name:  "success",
//			urlID: "1",
//			db: &mocks.URLRepository{
//				ByIdFn: func(ctx context.Context, id int) (url url_shortener.Url, err error) {
//					return url_shortener.Url{
//						ID:   1,
//						Url:  "http://google.com",
//						Code: "so1gFSl5",
//					}, nil
//				},
//			},
//			respBody: `{"id":1,"url":"http://google.com","code":"so1gFSl5"}` + "\n",
//		},
//		{
//			name:  "not found",
//			urlID: "2",
//			db: &mocks.URLRepository{
//				ByIdFn: func(ctx context.Context, id int) (url url_shortener.Url, err error) {
//					return url_shortener.Url{}, errors.New("not found")
//				},
//			},
//			respBody: ``,
//		},
//		{
//			name:  "bad id format",
//			urlID: "bad_id",
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := publicapi.NewRouter(publicapi.NewApiService(tt.db)).Handler()
//			srv := httptest.NewServer(r)
//			defer srv.Close()
//
//			resp, err := http.Get(fmt.Sprintf("%s%s%s", srv.URL, "/urls/", tt.urlID))
//			if err != nil {
//				t.Fatal(err)
//			}
//
//			body, err := ioutil.ReadAll(resp.Body)
//			if err != nil {
//				t.Fatal(err)
//			}
//			if err = resp.Body.Close(); err != nil {
//				t.Fatal(err)
//			}
//			if diff := cmp.Diff(tt.respBody, string(body)); diff != "" {
//				t.Errorf(diff)
//			}
//
//		})
//	}
//}