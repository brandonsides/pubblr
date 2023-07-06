package apiutil

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type Endpoint[T any] func(*http.Request) (T, Status)

func LogEndpoint[T any](endpoint Endpoint[T], logger Logger) Endpoint[T] {
	return Endpoint[T](func(r *http.Request) (ret T, status Status) {
		defer func() {
			if r := recover(); r != nil {
				logger.Errorf("Recovered from panic: %s\n%v", r, string(debug.Stack()))
				status = Statusf(http.StatusInternalServerError, "Recovered from panic: %s", r)

			}
		}()

		ret, status = endpoint(r)
		if status != nil && status.StatusCode()/100 == 2 {
			logger.Errorf("%s", status.Error())
		}
		return
	})
}

func (e Endpoint[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
		}
	}()

	resp, status := e(r)

	statusCode := http.StatusOK
	if status != nil {
		statusCode = status.StatusCode()
	}

	if statusCode/100 != 2 {
		w.WriteHeader(statusCode)
		if status.StatusCode()/100 == 4 {
			w.Write([]byte(status.Error()))
		}
		return
	}

	marshalled, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(marshalled)
}
