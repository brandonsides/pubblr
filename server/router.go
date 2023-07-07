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
	GetOutboxPage(user string, page int, pageSize int) ([]activitystreams.ActivityIface, error)
	GetOutboxCount(user string) (int, error)
	GetPost(user, typ, id string) (activitystreams.ObjectIface, error)
	CreateUser(user activitystreams.ActorIface, username, password string, baseIdUrl url.URL) (activitystreams.ActorIface, error)
	GetUser(username string) (activitystreams.ActorIface, error)
}

type PubblrRouter struct {
	chi.Router
	Database DB
	Logger   apiutil.Logger
	baseUrl  url.URL
	pageSize int
}

func NewPubblrRouter(cfg PubblrServerConfig, baseRouter chi.Router) chi.Router {
	if cfg.PageSize == 0 {
		cfg.PageSize = 50
	}
	router := PubblrRouter{
		Router:   chi.NewRouter(),
		Database: database.NewPubblrDatabase(cfg.Database),
		Logger:   logging.NewStandardPubblrLogger(cfg.Logger),
		baseUrl: url.URL{
			Scheme: "http",
			Host:   cfg.Host + ":" + strconv.Itoa(cfg.Port),
			Path:   cfg.MountPath,
		},
		pageSize: cfg.PageSize,
	}

	router.Use(
		SetContentType("application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""),
	)

	// OBJECTS
	router.Method("GET", "/{actor}/{type}/{id}", apiutil.LogEndpoint(router.GetObject, router.Logger))

	// ACTORS
	router.Method("GET", "/{actor}", apiutil.LogEndpoint(router.GetUser, router.Logger))
	router.Method("POST", "/{actor}", apiutil.LogEndpoint(router.PostUser, router.Logger))

	// INBOX
	router.Method("GET", "/{actor}/inbox", apiutil.LogEndpoint(router.GetInbox, router.Logger))
	router.Method("GET", "/{actor}/inbox/page/{page}", apiutil.LogEndpoint(router.GetInboxPage, router.Logger))

	// OUTBOX
	router.Method("POST", "/{actor}/outbox", apiutil.LogEndpoint(router.PostObject, router.Logger))
	router.Method("GET", "/{actor}/outbox", apiutil.LogEndpoint(router.GetOutbox, router.Logger))
	router.Method("GET", "/{actor}/outbox/page/{page}", apiutil.LogEndpoint(router.GetOutboxPage, router.Logger))
	router.Method("GET", "/{actor}/outbox/{id}", apiutil.LogEndpoint(router.GetOutboxActivity, router.Logger))

	// STREAMS
	router.Method("GET", "/{actor}/streams", apiutil.LogEndpoint(router.GetStreams, router.Logger))
	router.Method("GET", "/{actor}/streams/{id}", apiutil.LogEndpoint(router.GetStream, router.Logger))
	router.Method("GET", "/{actor}/streams/{id}/page/{page}", apiutil.LogEndpoint(router.GetStreamPage, router.Logger))
	router.Method("GET", "/{actor}/streams/{id}/followers",
		apiutil.LogEndpoint(router.GetStreamFollowers, router.Logger))
	router.Method("GET", "/{actor}/streams/{id}/followers/page/{page}",
		apiutil.LogEndpoint(router.GetStreamFollowersPage, router.Logger))

	// FOLLOWING
	router.Method("GET", "/{actor}/following", apiutil.LogEndpoint(router.GetFollowing, router.Logger))
	router.Method("GET", "/{actor}/following/page/{page}", apiutil.LogEndpoint(router.GetFollowingPage, router.Logger))

	// FOLLOWERS
	router.Method("GET", "/{actor}/followers", apiutil.LogEndpoint(router.GetFollowers, router.Logger))
	router.Method("GET", "/{actor}/followers/page/{page}", apiutil.LogEndpoint(router.GetFollowersPage, router.Logger))

	// LIKED
	router.Method("GET", "/{actor}/liked", apiutil.LogEndpoint(router.GetLiked, router.Logger))
	router.Method("GET", "/{actor}/liked/page/{page}", apiutil.LogEndpoint(router.GetLikedPage, router.Logger))

	if baseRouter != nil {
		baseRouter.Mount(cfg.MountPath, router)
		return baseRouter
	}
	return router
}
