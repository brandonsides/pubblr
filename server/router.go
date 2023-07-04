package server

import (
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/database"
	"github.com/brandonsides/pubblr/logging"
	"github.com/brandonsides/pubblr/server/apiutil"
	"github.com/go-chi/chi"
)

type PubblrRouter struct {
	chi.Router
	Database DB
	Logger   apiutil.Logger
	baseUrl  url.URL
}

func NewPubblrRouter(cfg PubblrServerConfig, baseRouter chi.Router) chi.Router {
	router := PubblrRouter{
		Router:   chi.NewRouter(),
		Database: database.NewPubblrDatabase(cfg.Database),
		Logger:   logging.NewStandardPubblrLogger(cfg.Logger),
		baseUrl: url.URL{
			Scheme: "http",
			Host:   cfg.Host + ":" + strconv.Itoa(cfg.Port),
			Path:   cfg.MountPath,
		},
	}

	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
			h.ServeHTTP(w, r)
		})
	})

	router.Method("POST", "/create", apiutil.LogEndpoint(router.DoActivity, router.Logger))
	router.Method("GET", "/{type}/{id}", apiutil.LogEndpoint(router.GetObject, router.Logger))

	if baseRouter != nil {
		baseRouter.Mount(cfg.MountPath, router)
		return baseRouter
	}
	return router
}

func (s *PubblrRouter) ToId(ent activitystreams.EntityIface, id string) string {
	typ, err := ent.Type()
	if err != nil {
		typ = "Object"
	}
	retUrl := s.baseUrl
	retUrl.Path = path.Join(retUrl.Path, strings.ToLower(typ), id)
	return retUrl.String()
}
