package main

import (
	"net/http"
	"time"

	"github.com/brandonsides/pubblr/server"
	"github.com/brandonsides/pubblr/server/auth"
	"github.com/go-chi/chi"
)

func main() {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}
	router := chi.NewRouter()
	router.Get("/hello", handlerFunc)
	server.NewPubblrServer(server.PubblrServerConfig{
		Host:      "localhost",
		Port:      8080,
		MountPath: "/pubblr",
		PageSize:  25,
		Auth: auth.AuthConfig{
			AuthKeyLocation:       "auth.pem",
			JWTExpirationDuration: 36 * time.Hour,
		},
	}, router).ListenAndServe()
}
