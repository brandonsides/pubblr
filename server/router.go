package server

import (
	"net/url"
	"strconv"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/database"
	"github.com/brandonsides/pubblr/logging"
	"github.com/brandonsides/pubblr/server/apiutil"
	"github.com/go-chi/chi"
)

type DB interface {
	CreateObject(obj activitystreams.ObjectIface, user string, baseIdUrl url.URL) (activitystreams.ObjectIface, error)
	CreateActivity(act activitystreams.ActivityIface, user string, baseIdUrl url.URL) (activitystreams.ActivityIface, error)
	GetActivity(user, id string) (activitystreams.ActivityIface, error)
	GetPost(user, typ, id string) (activitystreams.ObjectIface, error)
	CreateUser(user activitystreams.ActorIface, username string, baseIdUrl url.URL) (activitystreams.ActorIface, error)
	GetUser(username string) (activitystreams.ActorIface, error)
}

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

	router.Use(
		SetContentType("application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""),
	)

	// OUTBOX
	router.Method("POST", "/{actor}/outbox", apiutil.LogEndpoint(router.PostObject, router.Logger))
	router.Method("GET", "/{actor}/outbox", apiutil.LogEndpoint(router.GetOutbox, router.Logger))
	router.Method("GET", "/{actor}/outbox/{id}", apiutil.LogEndpoint(router.GetOutboxActivity, router.Logger))

	// INBOX
	router.Method("GET", "/{actor}/inbox", apiutil.LogEndpoint(router.GetInbox, router.Logger))

	// OBJECTS
	router.Method("GET", "/{actor}/{type}/{id}", apiutil.LogEndpoint(router.GetObject, router.Logger))

	// ACTORS
	router.Method("GET", "/{actor}", apiutil.LogEndpoint(router.GetUser, router.Logger))
	router.Method("POST", "/{actor}", apiutil.LogEndpoint(router.PostUser, router.Logger))

	if baseRouter != nil {
		baseRouter.Mount(cfg.MountPath, router)
		return baseRouter
	}
	return router
}
