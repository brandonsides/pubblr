package server

import (
	"net/http"
	"strings"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/server/apiutil"
)

func (router *PubblrRouter) Create(create activitystreams.Create) (activitystreams.ObjectIface, apiutil.Status) {
	if create.Actor == nil {
		return nil, apiutil.NewStatus(http.StatusBadRequest, "Create activity must have an actor")
	}

	if create.Object == nil {
		return nil, apiutil.NewStatus(http.StatusBadRequest, "Create activity must have an object")
	}

	actorObjIface, ok := create.Actor.(activitystreams.ObjectIface)
	if !ok {
		return nil, apiutil.NewStatus(http.StatusBadRequest, "Create activity object must be an Object")
	}

	objectObjIface, ok := create.Object.(activitystreams.ObjectIface)
	if !ok {
		return nil, apiutil.NewStatus(http.StatusBadRequest, "Create activity object must be an Object")
	}
	object := activitystreams.ToObject(objectObjIface)

	mergedAttributedTo := merge([]activitystreams.EntityIface{actorObjIface}, object.AttributedTo, create.AttributedTo)
	create.AttributedTo = mergedAttributedTo
	object.AttributedTo = mergedAttributedTo

	actorId := activitystreams.ToObject(actorObjIface).Id
	shortId := shortId(actorId)

	router.Database.CreateActivity(&create, shortId, router.baseUrl)
	router.Database.CreateObject(objectObjIface, shortId, router.baseUrl)

	return object, apiutil.StatusFromCode(http.StatusCreated)
}

func shortId(id string) string {
	spl := strings.Split(id, "/")
	return spl[len(spl)-1]
}

func merge(slices ...[]activitystreams.EntityIface) []activitystreams.EntityIface {
	alreadyIncluded := make(map[string]bool)

	var merged []activitystreams.EntityIface
	for _, slice := range slices {
		for _, entityIface := range slice {
			entity := activitystreams.ToEntity(entityIface)
			if _, ok := alreadyIncluded[entity.Id]; !ok {
				merged = append(merged, entityIface)
				alreadyIncluded[entity.Id] = true
			}
		}
	}
	return merged
}
