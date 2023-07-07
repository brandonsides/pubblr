package server

import (
	"io/ioutil"
	"net/http"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/server/apiutil"
	"github.com/go-chi/chi"
)

type CreateAccountRequest struct {
	Password string                     `json:"password"`
	Actor    activitystreams.ActorIface `json:"actor"`
}

func (router *PubblrRouter) GetInbox(r *http.Request) ([]activitystreams.ObjectIface, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetInbox not implemented")
}

func (router *PubblrRouter) GetOutbox(r *http.Request) ([]activitystreams.ActivityIface, apiutil.Status) {
	actor := chi.URLParam(r, "actor")

	activities, err := router.Database.GetOutbox(actor)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}

	return activities, apiutil.StatusFromCode(http.StatusOK)
}

func (router *PubblrRouter) GetOutboxActivity(r *http.Request) (activitystreams.ObjectIface, apiutil.Status) {
	id := chi.URLParam(r, "id")
	user := chi.URLParam(r, "actor")

	activity, err := router.Database.GetActivity(user, id)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}

	return activity, apiutil.StatusFromCode(http.StatusOK)
}

func (router *PubblrRouter) GetUser(r *http.Request) (activitystreams.ObjectIface, apiutil.Status) {
	username := chi.URLParam(r, "actor")

	user, err := router.Database.GetUser(username)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}

	router.setEndpoints(user)

	return user, apiutil.StatusFromCode(http.StatusOK)
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

	switch typ {
	case "Create":
		return router.Create(*activityIface.(*activitystreams.Create))
	default:
		return nil, apiutil.Statusf(http.StatusBadRequest, "invalid ActivityStreams activity type: %s", typ)
	}
}

func (router *PubblrRouter) GetObject(r *http.Request) (activitystreams.ObjectIface, apiutil.Status) {
	user := chi.URLParam(r, "actor")
	typ := chi.URLParam(r, "type")
	id := chi.URLParam(r, "id")

	post, err := router.Database.GetPost(user, typ, id)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}

	return post, nil
}

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
}
