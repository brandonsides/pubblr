package server

import (
	"io/ioutil"
	"net/http"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/server/apiutil"
	"github.com/go-chi/chi"
)

func (router *PubblrRouter) GetInbox(r *http.Request) ([]activitystreams.ObjectIface, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetInbox not implemented")
}

func (router *PubblrRouter) GetOutbox(r *http.Request) ([]activitystreams.ObjectIface, apiutil.Status) {
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetOutbox not implemented")
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
	return nil, apiutil.Statusf(http.StatusNotImplemented, "GetUser not implemented")
}

func (router *PubblrRouter) PostUser(r *http.Request) (activitystreams.ObjectIface, apiutil.Status) {
	username := chi.URLParam(r, "actor")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusBadRequest, err)
	}

	var actor activitystreams.ActorIface
	err = activitystreams.DefaultEntityUnmarshaler.Unmarshal(b, &actor)
	if err != nil {
		return nil, apiutil.Statusf(http.StatusBadRequest, "invalid ActivityStreams actor: %w", err)
	}

	actor, err = router.Database.CreateUser(actor, username, router.baseUrl)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusInternalServerError, err)
	}

	return actor, apiutil.Statusf(http.StatusCreated, "created user %s", username)
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
