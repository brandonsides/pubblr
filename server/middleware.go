package server

import "net/http"

type Middleware func(http.Handler) http.Handler

func SetContentType(contentType string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", contentType)
			next.ServeHTTP(w, r)
		})
	}
}
