package apiutil

import "errors"

type UntypedEndpointMiddleware interface {
	Wrap(UntypedEndpoint) (UntypedEndpoint, error)
}

type EndpointMiddleware[T, U any] func(Endpoint[T]) *Endpoint[U]

func (mw EndpointMiddleware[T, U]) Wrap(u UntypedEndpoint) (UntypedEndpoint, error) {
	e1, ok := u.(Endpoint[T])
	if !ok {
		return nil, errors.New("EndpointMiddleware: expected endpoint of type Endpoint[T]")
	}

	e2 := mw(e1)
	return e2, nil
}
