package main

import (
	"net/http"

	"github.com/brandonsides/pubblr/server"
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
	}, router).ListenAndServe()
}
