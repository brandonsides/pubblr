package main

import (
	"flag"
	"net/http"

	"github.com/brandonsides/pubblr/server"
	"github.com/go-chi/chi"
)

func main() {
	confPath := flag.String("config", "config.json", "Path to config file")
	flag.Parse()

	config, err := server.LoadPubblrRouterConfig(*confPath)
	if err != nil {
		panic(err)
	}

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}
	router := chi.NewRouter()
	router.Get("/hello", handlerFunc)
	srv, err := server.NewPubblrServer(config, router)
	if err != nil {
		panic(err)
	}

	srv.ListenAndServe()
}
