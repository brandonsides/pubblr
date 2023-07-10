package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/server/apiutil"
	"github.com/brandonsides/pubblr/util/either"
	"github.com/go-chi/chi"
)

// AUTH

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (router *PubblrRouter) Login(r *http.Request) (string, apiutil.Status) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", apiutil.NewStatus(http.StatusBadRequest, "Invalid JSON")
	}
	var body LoginRequest
	err = json.Unmarshal(reqBody, &body)
	if err != nil {
		return "", apiutil.NewStatus(http.StatusBadRequest, "Invalid JSON")
	}

	if body.Username == "" {
		return "", apiutil.NewStatus(http.StatusBadRequest, "Missing username")
	}

	if body.Password == "" {
		return "", apiutil.NewStatus(http.StatusBadRequest, "Missing password")
	}

	if router.Database.CheckPassword(body.Username, body.Password) != nil {
		return "", apiutil.NewStatus(http.StatusUnauthorized, "Invalid username or password")
	}

	ret, err := router.Auth.GenerateToken(body.Username)
	if err != nil {
		return "", apiutil.NewStatus(http.StatusInternalServerError, "Error generating token")
	}
	return ret, nil
}

// OBJECTS

func (router *PubblrRouter) GetObject(r *http.Request) (activitystreams.ObjectIface, apiutil.Status) {
	user := chi.URLParam(r, "actor")
	typ := chi.URLParam(r, "type")
	id := chi.URLParam(r, "id")

	post, err := router.Database.GetObject(user, typ, id)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}

	return post, nil
}

// ACTORS

func (router *PubblrRouter) GetUser(r *http.Request) (activitystreams.ObjectIface, apiutil.Status) {
	username := chi.URLParam(r, "actor")

	user, err := router.Database.GetUser(username)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}

	router.setEndpoints(user)

	return user, apiutil.StatusFromCode(http.StatusOK)
}

type CreateAccountRequest struct {
	Password string                     `json:"password"`
	Actor    activitystreams.ActorIface `json:"actor"`
}

func (router *PubblrRouter) PostUser(r *http.Request) (activitystreams.ObjectIface, apiutil.Status) {
	username := chi.URLParam(r, "actor")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusBadRequest, err)
	}

	var createAccountRequest CreateAccountRequest
	err = activitystreams.DefaultEntityUnmarshaler.Unmarshal(b, &createAccountRequest)
	if err != nil {
		return nil, apiutil.Statusf(http.StatusBadRequest, "invalid ActivityStreams actor: %w", err)
	}

	createAccountRequest.Actor, err = router.Database.CreateUser(createAccountRequest.Actor, username,
		createAccountRequest.Password, router.baseUrl)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusInternalServerError, err)
	}

	router.setEndpoints(createAccountRequest.Actor)

	return createAccountRequest.Actor, apiutil.Statusf(http.StatusCreated, "created user %s", username)
}

//INBOX

func (router *PubblrRouter) PostToInbox(r *http.Request) (*activitystreams.ActivityIface, apiutil.Status) {
	return nil, apiutil.NewStatus(http.StatusNotImplemented, "PostToInbox not yet implemented")
}

func (router *PubblrRouter) GetInbox(r *http.Request) (*activitystreams.Collection, apiutil.Status) {
	// params
	actorShortId := chi.URLParam(r, "actor")

	// Make sure user actually exists
	actorIface, err := router.Database.GetUser(actorShortId)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}
	actor := activitystreams.ToActor(actorIface)

	router.setEndpoints(actorIface)

	count, err := router.Database.GetInboxCount(actorShortId)
	if err != nil {
		count = 0
	}

	ret := &activitystreams.Collection{
		Object: activitystreams.Object{
			Entity: activitystreams.Entity{
				Id: actor.Id + "/inbox",
			},
		},
		TotalItems: uint64(count),
		Ordered:    true,
	}

	if ret.TotalItems == 0 {
		return ret, apiutil.StatusFromCode(http.StatusOK)
	}
	lastPage := ret.TotalItems / uint64(router.pageSize)

	ret.First = either.Left[*activitystreams.CollectionPage, activitystreams.LinkIface](
		&activitystreams.CollectionPage{
			Collection: activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: actor.Id + "/inbox/page/0",
					},
				},
			},
		},
	)
	ret.Last = either.Left[*activitystreams.CollectionPage, activitystreams.LinkIface](
		&activitystreams.CollectionPage{
			Collection: activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: actor.Id + "/inbox/page/" + strconv.Itoa(int(lastPage)),
					},
				},
			},
		},
	)

	return ret, apiutil.StatusFromCode(http.StatusOK)
}

func (router *PubblrRouter) GetInboxPage(r *http.Request) (*activitystreams.CollectionPage, apiutil.Status) {
	actorShortId := chi.URLParam(r, "actor")
	page, err := strconv.Atoi(chi.URLParam(r, "page"))
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusBadRequest, err)
	}

	actorIface, err := router.Database.GetUser(actorShortId)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}
	actor := activitystreams.ToActor(actorIface)

	posts, err := router.Database.GetInboxPage(actorShortId, page, router.pageSize)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusInternalServerError, err)
	}

	items := make([]*either.Either[activitystreams.ObjectIface, activitystreams.LinkIface], len(posts))
	for i, post := range posts {
		obj := activitystreams.ToObject(post)
		obj.Bcc = nil
		obj.Bto = nil
		items[i] = either.Left[activitystreams.ObjectIface, activitystreams.LinkIface](post)
	}

	return &activitystreams.CollectionPage{
		Collection: activitystreams.Collection{
			Object: activitystreams.Object{
				Entity: activitystreams.Entity{
					Id: actor.Id + "/inbox/page/" + strconv.Itoa(page),
				},
			},
			Ordered: true,
			Items:   items,
		},
		PartOf: either.Left[activitystreams.Collection, activitystreams.Link](
			activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: actor.Id + "/inbox",
					},
				},
			},
		),
	}, nil
}

func (router *PubblrRouter) GetInboxItem(r *http.Request) (activitystreams.ActivityIface, apiutil.Status) {
	id := chi.URLParam(r, "id")
	user := chi.URLParam(r, "actor")

	activity, err := router.Database.GetInboxItem(user, id)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}

	return activity, apiutil.StatusFromCode(http.StatusOK)
}

// OUTBOX
func (router *PubblrRouter) GetOutbox(r *http.Request) (*activitystreams.Collection, apiutil.Status) {
	// params
	actorShortId := chi.URLParam(r, "actor")

	// Make sure user actually exists
	actorIface, err := router.Database.GetUser(actorShortId)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}
	actor := activitystreams.ToActor(actorIface)

	router.setEndpoints(actorIface)

	count, err := router.Database.GetOutboxCount(actorShortId)
	if err != nil {
		count = 0
	}

	ret := &activitystreams.Collection{
		Object: activitystreams.Object{
			Entity: activitystreams.Entity{
				Id: actor.Id + "/outbox",
			},
		},
		TotalItems: uint64(count),
		Ordered:    true,
	}

	if ret.TotalItems == 0 {
		return ret, apiutil.StatusFromCode(http.StatusOK)
	}
	lastPage := ret.TotalItems / uint64(router.pageSize)

	ret.First = either.Left[*activitystreams.CollectionPage, activitystreams.LinkIface](
		&activitystreams.CollectionPage{
			Collection: activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: actor.Id + "/outbox/page/0",
					},
				},
			},
		},
	)
	ret.Last = either.Left[*activitystreams.CollectionPage, activitystreams.LinkIface](
		&activitystreams.CollectionPage{
			Collection: activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: actor.Id + "/outbox/page/" + strconv.Itoa(int(lastPage)),
					},
				},
			},
		},
	)

	return ret, apiutil.StatusFromCode(http.StatusOK)
}

func (router *PubblrRouter) GetOutboxPage(r *http.Request) (*activitystreams.CollectionPage, apiutil.Status) {
	actorShortId := chi.URLParam(r, "actor")
	page, err := strconv.Atoi(chi.URLParam(r, "page"))
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusBadRequest, err)
	}

	actorIface, err := router.Database.GetUser(actorShortId)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}
	actor := activitystreams.ToActor(actorIface)

	posts, err := router.Database.GetOutboxPage(actorShortId, page, router.pageSize)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusInternalServerError, err)
	}

	items := make([]*either.Either[activitystreams.ObjectIface, activitystreams.LinkIface], len(posts))
	for i, post := range posts {
		obj := activitystreams.ToObject(post)
		obj.Bcc = nil
		obj.Bto = nil
		items[i] = either.Left[activitystreams.ObjectIface, activitystreams.LinkIface](post)
	}

	return &activitystreams.CollectionPage{
		Collection: activitystreams.Collection{
			Object: activitystreams.Object{
				Entity: activitystreams.Entity{
					Id: actor.Id + "/outbox/page/" + strconv.Itoa(page),
				},
			},
			Ordered: true,
			Items:   items,
		},
		PartOf: either.Left[activitystreams.Collection, activitystreams.Link](
			activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: actor.Id + "/outbox",
					},
				},
			},
		),
	}, nil
}

func (router *PubblrRouter) PostObject(r *http.Request) (activitystreams.ObjectIface, apiutil.Status) {
	actorId := chi.URLParam(r, "actor")
	actor, err := router.Database.GetUser(actorId)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusBadRequest, err)
	}

	var e activitystreams.TopLevelEntity
	err = activitystreams.DefaultEntityUnmarshaler.Unmarshal(b, &e)
	if err != nil {
		return nil, apiutil.Statusf(http.StatusBadRequest, "invalid ActivityStreams entity: %w", err)
	}

	activityIface, ok := e.EntityIface.(activitystreams.ActivityIface)
	if !ok {
		activityIface = &activitystreams.Create{
			TransitiveActivity: activitystreams.TransitiveActivity{
				Object: e.EntityIface,
			},
		}
	}
	typ, err := activityIface.Type()
	if err != nil {
		return nil, apiutil.Statusf(http.StatusBadRequest, "invalid ActivityStreams entity: %w", err)
	}

	intransitiveActivity := activitystreams.ToIntransitiveActivity(activityIface)
	intransitiveActivity.Actor = actor
	intransitiveActivity.AttributedTo = []activitystreams.EntityIface{actor}

	var result activitystreams.ObjectIface
	var status apiutil.Status
	switch typ {
	case "Create":
		result, status = router.Create(activityIface.(*activitystreams.Create))
	default:
		status = apiutil.Statusf(http.StatusBadRequest, "invalid ActivityStreams activity type: %s", typ)
	}
	if !apiutil.IsOK(status) {
		return nil, status
	}

	router.Deliver(activityIface)

	return result, apiutil.StatusFromCode(http.StatusCreated)
}

func (router *PubblrRouter) GetOutboxActivity(r *http.Request) (activitystreams.ObjectIface, apiutil.Status) {
	id := chi.URLParam(r, "id")
	user := chi.URLParam(r, "actor")

	activity, err := router.Database.GetOutboxItem(user, id)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}

	return activity, apiutil.StatusFromCode(http.StatusOK)
}

// STREAMS

func (router *PubblrRouter) GetStreams(r *http.Request) (*activitystreams.Collection, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetStreams not implemented")
}

func (router *PubblrRouter) GetStream(r *http.Request) (*activitystreams.Collection, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetStream not implemented")
}

func (router *PubblrRouter) GetStreamPage(r *http.Request) (*activitystreams.CollectionPage, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetStreamPage not implemented")
}

func (router *PubblrRouter) GetStreamFollowers(r *http.Request) (*activitystreams.Collection, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetStreamFollowers not implemented")
}

func (router *PubblrRouter) GetStreamFollowersPage(r *http.Request) (*activitystreams.CollectionPage, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetStreamFollowersPage not implemented")
}

// FOLLOWING

func (router *PubblrRouter) GetFollowing(r *http.Request) (*activitystreams.Collection, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetFollowing not implemented")
}

func (router *PubblrRouter) GetFollowingPage(r *http.Request) (*activitystreams.CollectionPage, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetFollowingPage not implemented")
}

// FOLLOWERS

func (router *PubblrRouter) GetFollowers(r *http.Request) (*activitystreams.Collection, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetFollowers not implemented")
}

func (router *PubblrRouter) GetFollowersPage(r *http.Request) (*activitystreams.CollectionPage, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetFollowersPage not implemented")
}

// LIKED

func (router *PubblrRouter) GetLiked(r *http.Request) (*activitystreams.Collection, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetLiked not implemented")
}

func (router *PubblrRouter) GetLikedPage(r *http.Request) (*activitystreams.CollectionPage, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetLikedPage not implemented")
}

// helpers

func (router *PubblrRouter) setEndpoints(a activitystreams.ActorIface) {
	actor := activitystreams.ToActor(a)

	inbox := &activitystreams.Collection{}
	inbox.Id = actor.Id + "/inbox"
	actor.Inbox = inbox

	outbox := &activitystreams.Collection{}
	outbox.Id = actor.Id + "/outbox"
	actor.Outbox = outbox

	following := &activitystreams.Collection{}
	following.Id = actor.Id + "/following"
	actor.Following = following

	followers := &activitystreams.Collection{}
	followers.Id = actor.Id + "/followers"
	actor.Followers = followers

	liked := &activitystreams.Collection{}
	liked.Id = actor.Id + "/liked"
	actor.Liked = liked

	streams := &activitystreams.Collection{}
	streams.Id = actor.Id + "/streams"
	actor.Streams = streams
}
