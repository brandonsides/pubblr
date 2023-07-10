package server

import (
	"context"
	"net/http"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/server/apiutil"
	"github.com/go-chi/chi"
)

type Middleware func(http.Handler) http.Handler

func SetContentType(contentType string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", contentType)
			next.ServeHTTP(w, r)
		})
	}
}

func AuthMiddleware[T any](auth Auth, next apiutil.Endpoint[T]) apiutil.Endpoint[*T] {
	return apiutil.Endpoint[*T](func(r *http.Request) (*T, apiutil.Status) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			return nil, apiutil.NewStatus(http.StatusUnauthorized, "Missing authorization header")
		}

		owner := chi.URLParam(r, "actor")

		username, err := auth.VerifyToken(tokenString)
		if err == nil {
			r = r.WithContext(context.WithValue(r.Context(), "username", username))
		}

		var ret interface{}
		ret, status := next(r)
		if !apiutil.IsOK(status) {
			return nil, status
		}

		retObject, ok := ret.(activitystreams.ObjectIface)
		if !ok {
			return ret.(*T), status
		}

		if !intendedFor(username, retObject) {
			return nil, apiutil.NewStatus(http.StatusForbidden, "You are not authorized to access this resource")
		}

		if owner != "" && owner != username {
			object := activitystreams.ToObject(retObject)
			object.Bcc = nil
			object.Bto = nil
		}

		return ret.(*T), status
	})
}

func intendedFor(username string, objectIface activitystreams.ObjectIface) bool {
	object := activitystreams.ToObject(objectIface)
	return in(username, object.To, object.Bto, object.Audience)
}

func in(username string, entities ...[]activitystreams.EntityIface) bool {
	for _, entityList := range entities {
		for _, entityIface := range entityList {
			entity := activitystreams.ToEntity(entityIface)
			// TODO: Also check that this entity is local
			if shortId(entity.Id) == username {
				return true
			}
		}
	}
	return false
}
