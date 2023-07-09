package apiutil

import (
	"errors"
	"fmt"
)

type Status interface {
	error
	StatusCode() int
}

type status struct {
	e          error
	statusCode int
}

func (e *status) StatusCode() int {
	return e.statusCode
}

func (e *status) Error() string {
	if e.e == nil {
		return ""
	}
	return e.e.Error()
}

func NewStatusFromError(statusCode int, err error) Status {
	return &status{statusCode: statusCode, e: err}
}

func NewStatus(statusCode int, message string) Status {
	return &status{statusCode: statusCode, e: errors.New(message)}
}

func Statusf(statusCode int, format string, args ...interface{}) Status {
	return &status{statusCode: statusCode, e: fmt.Errorf(format, args...)}
}

func StatusFromCode(statusCode int) Status {
	return &status{statusCode: statusCode}
}

func IsOK(s Status) bool {
	return s == nil || s.StatusCode()/100 == 2
}
