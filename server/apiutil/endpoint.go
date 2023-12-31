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

type UntypedEndpoint interface {
	http.Handler
	Process(r *http.Request) (interface{}, http.Header, Status)
}

type Endpoint[T any] func(*http.Request) (T, http.Header, Status)

// Endpoint[T] implements UntypedEndpoint
var _ UntypedEndpoint = Endpoint[int](nil)

func LogEndpoint[T any](endpoint Endpoint[T], logger Logger) Endpoint[T] {
	return Endpoint[T](func(r *http.Request) (ret T, header http.Header, status Status) {
		defer func() {
			if r := recover(); r != nil {
				logger.Errorf("Recovered from panic: %s\n%v", r, string(debug.Stack()))
				status = Statusf(http.StatusInternalServerError, "Recovered from panic: %s", r)

			}
		}()

		ret, header, status = endpoint(r)
		if status != nil && status.StatusCode()/100 != 2 {
			logger.Errorf("%s\n", status.Error())
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

	resp, header, status := e(r)

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

	for k, v := range header {
		w.Header()[k] = v
	}
	w.WriteHeader(statusCode)
	w.Write(marshalled)
}

func (e Endpoint[T]) Process(r *http.Request) (interface{}, http.Header, Status) {
	return e(r)
}
