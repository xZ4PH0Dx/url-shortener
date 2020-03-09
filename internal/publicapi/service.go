package publicapi

import (
	"context"
	"fmt"
	"net/http"
	"url_shortener"
)

func NewApiService(handler http.Handler) *url_shortener.PublicAPIServer {
	return &apiService{}
}

type apiService struct {
	//TODO some vars
}

func (a *apiService) CreateUrl(ctx context.Context) error {
	q := ctx.Value("")
	fmt.Println(q)
	return nil
}
