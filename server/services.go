package server

import (
	"io/ioutil"
	"net/http"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/server/apiutil"
	"github.com/go-chi/chi"
)

func (router *PubblrRouter) Create(c *activitystreams.Create) {

}

func (router *PubblrRouter) DoActivity(r *http.Request) (*activitystreams.ObjectIface, apiutil.Status) {
	var e activitystreams.TopLevelEntity

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusBadRequest, err)
	}

	err = activitystreams.DefaultEntityUnmarshaler.Unmarshal(b, &e)
	if err != nil {
		return nil, apiutil.Statusf(http.StatusBadRequest, "invalid ActivityStreams entity: %w", err)
	}

	activity, ok := e.EntityIface.(activitystreams.ActivityIface)
	if !ok {
		activity = &activitystreams.Create{
			TransitiveActivity: activitystreams.TransitiveActivity{
				Object: e.EntityIface,
			},
		}
	}

	typ, err := activity.Type()
	switch typ {
	case "Create":
		return router.Create(activity.(*activitystreams.Create))
	default:
		return nil, apiutil.Statusf(http.StatusBadRequest, "invalid ActivityStreams activity type: %s", typ)
	}
}

func (router *PubblrRouter) GetObject(r *http.Request) (*activitystreams.ObjectIface, apiutil.Status) {
	typ := chi.URLParam(r, "type")
	id := chi.URLParam(r, "id")

	post, err := router.Database.GetPostByTypeAndId(typ, id)
	if err != nil {
		return nil, apiutil.NewStatusFromError(http.StatusNotFound, err)
	}

	activitystreams.ToEntity(post).Id = router.ToId(post, id)

	return &post, nil
}
