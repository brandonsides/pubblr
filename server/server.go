package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func NewPubblrServer(config PubblrRouterConfig, baseRouter chi.Router) *http.Server {
	router, err := NewPubblrRouter(config, baseRouter)
	if err != nil {
		panic(err)
	}
	return &http.Server{
		Addr:    config.Host + ":" + strconv.Itoa(config.Port),
		Handler: router,
	}
}
