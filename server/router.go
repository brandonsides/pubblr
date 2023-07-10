package server

import (
	"net/url"
	"strconv"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/database"
	"github.com/brandonsides/pubblr/logging"
	"github.com/brandonsides/pubblr/server/apiutil"
	"github.com/brandonsides/pubblr/server/auth"
	"github.com/go-chi/chi"
)

type DB interface {
	CreateObject(obj activitystreams.ObjectIface, user string, baseIdUrl url.URL) (activitystreams.ObjectIface, error)
	CreateInboxItem(item activitystreams.ActivityIface, user string) (activitystreams.ActivityIface, error)
	CreateOutboxItem(act activitystreams.ActivityIface, user string, baseIdUrl url.URL) (activitystreams.ActivityIface, error)
	GetOutboxItem(user, id string) (activitystreams.ActivityIface, error)
	GetInboxPage(user string, page, pageSize int) ([]activitystreams.ActivityIface, error)
	GetInboxCount(user string) (int, error)
	GetInboxItem(user, id string) (activitystreams.ActivityIface, error)
	GetOutboxPage(user string, page, pageSize int) ([]activitystreams.ActivityIface, error)
	GetOutboxCount(user string) (int, error)
	GetObject(user, typ, id string) (activitystreams.ObjectIface, error)
	CreateUser(user activitystreams.ActorIface, username, password string, baseIdUrl url.URL) (activitystreams.ActorIface, error)
	GetUser(username string) (activitystreams.ActorIface, error)
	CheckPassword(username, password string) error
}

type Auth interface {
	GenerateToken(username string) (string, error)
	VerifyToken(token string) (string, error)
}

type PubblrRouter struct {
	chi.Router
	Database DB
	Logger   apiutil.Logger
	Auth     Auth
	baseUrl  url.URL
	pageSize int
}

func NewPubblrRouter(cfg PubblrServerConfig, baseRouter chi.Router) (chi.Router, error) {
	if cfg.PageSize == 0 {
		cfg.PageSize = 50
	}

	auth, err := auth.NewAuth(cfg.Auth)
	if err != nil {
		panic(err)
	}

	router := PubblrRouter{
		Router:   chi.NewRouter(),
		Database: database.NewPubblrDatabase(cfg.Database),
		Logger:   logging.NewStandardPubblrLogger(cfg.Logger),
		Auth:     auth,
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

	// AUTH
	router.Method("POST", "/login", apiutil.LogEndpoint(router.Login, router.Logger))

	// OBJECTS
	router.Method("GET", "/{actor}/{type}/{id}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetObject), router.Logger))

	// ACTORS
	router.Method("GET", "/{actor}", apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetUser), router.Logger))
	router.Method("POST", "/{actor}", apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.PostUser), router.Logger))

	// INBOX
	router.Method("POST", "/{actor}/inbox",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.PostToInbox), router.Logger))
	router.Method("GET", "/{actor}/inbox",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetInbox), router.Logger))
	router.Method("GET", "/{actor}/inbox/page/{page}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetInboxPage), router.Logger))
	router.Method("GET", "/{actor}/inbox/{id}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetInboxItem), router.Logger))

	// OUTBOX
	router.Method("POST", "/{actor}/outbox",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.PostObject), router.Logger))
	router.Method("GET", "/{actor}/outbox",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetOutbox), router.Logger))
	router.Method("GET", "/{actor}/outbox/page/{page}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetOutboxPage), router.Logger))
	router.Method("GET", "/{actor}/outbox/{id}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetOutboxActivity), router.Logger))

	// STREAMS
	router.Method("GET", "/{actor}/streams",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetStreams), router.Logger))
	router.Method("GET", "/{actor}/streams/{id}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetStream), router.Logger))
	router.Method("GET", "/{actor}/streams/{id}/page/{page}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetStreamPage), router.Logger))
	router.Method("GET", "/{actor}/streams/{id}/followers",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetStreamFollowers), router.Logger))
	router.Method("GET", "/{actor}/streams/{id}/followers/page/{page}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetStreamFollowersPage), router.Logger))

	// FOLLOWING
	router.Method("GET", "/{actor}/following",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetFollowing), router.Logger))
	router.Method("GET", "/{actor}/following/page/{page}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetFollowingPage), router.Logger))

	// FOLLOWERS
	router.Method("GET", "/{actor}/followers",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetFollowers), router.Logger))
	router.Method("GET", "/{actor}/followers/page/{page}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetFollowersPage), router.Logger))

	// LIKED
	router.Method("GET", "/{actor}/liked",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetLiked), router.Logger))
	router.Method("GET", "/{actor}/liked/page/{page}",
		apiutil.LogEndpoint(AuthMiddleware(router.Auth, router.GetLikedPage), router.Logger))

	if baseRouter != nil {
		baseRouter.Mount(cfg.MountPath, router)
		return baseRouter, nil
	}
	return router, nil
}
