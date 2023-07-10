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
		owner := chi.URLParam(r, "actor")

		var username string
		var err error
		tokenString := r.Header.Get("Authorization")
		if tokenString != "" {
			username, err = auth.VerifyToken(tokenString)
			if err == nil {
				r = r.WithContext(context.WithValue(r.Context(), "username", username))
			}
		}

		if r.Method == "POST" && owner != "" && owner != username {
			return nil, apiutil.NewStatus(http.StatusForbidden, "You are not authorized act on behalf of this user")
		}

		ret, status := next(r)
		var retInterface interface{} = ret
		if !apiutil.IsOK(status) {
			return nil, status
		}

		retObject, ok := retInterface.(activitystreams.ObjectIface)
		if !ok {
			return &ret, status
		}

		if !intendedFor(username, owner, retObject) {
			return nil, apiutil.NewStatus(http.StatusForbidden, "You are not authorized to access this resource")
		}

		if owner != "" && owner != username {
			object := activitystreams.ToObject(retObject)
			object.Bcc = nil
			object.Bto = nil
		}

		return &ret, status
	})
}

func intendedFor(username string, owner string, objectIface activitystreams.ObjectIface) bool {
	object := activitystreams.ToObject(objectIface)

	return username == owner || in(
		username, mapitems(
			func(e activitystreams.EntityIface) string {
				entity := activitystreams.ToEntity(e)
				return shortId(entity.Id)
			}, object.To, object.Cc, object.Bto, object.Bcc, object.Audience,
		),
	) || in(
		"https://www.w3.org/ns/activitystreams#Public", mapitems(
			func(e activitystreams.EntityIface) string {
				entity := activitystreams.ToEntity(e)
				return entity.Id
			}, object.To, object.Cc, object.Bto, object.Bcc, object.Audience,
		),
	)
}

func in[T comparable](i T, tArrs ...[]T) bool {
	for _, tArr := range tArrs {
		for _, t := range tArr {
			if i == t {
				return true
			}
		}
	}
	return false
}

func mapitems[T, U any](f func(T) U, items ...[]T) []U {
	var ret []U
	for _, itemlist := range items {
		for _, item := range itemlist {
			ret = append(ret, f(item))
		}
	}
	return ret
}
