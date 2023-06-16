package util

import (
	"encoding/json"
	"unsafe"
)

type Either[A, B any] struct {
	isLeft bool
	data   unsafe.Pointer
}

func Left[A, B any](a A) *Either[A, B] {
	return &Either[A, B]{
		isLeft: true,
		data:   unsafe.Pointer(&a),
	}
}

func Right[A, B any](b B) *Either[A, B] {
	return &Either[A, B]{
		isLeft: false,
		data:   unsafe.Pointer(&b),
	}
}

func (e Either[A, B]) Left() *A {
	if !e.isLeft {
		return nil
	}
	return (*A)(e.data)
}

func (e Either[A, B]) Right() *B {
	if e.isLeft {
		return nil
	}
	return (*B)(e.data)
}

func (e Either[A, B]) MarshalJSON() ([]byte, error) {
	if e.isLeft {
		return json.Marshal(e.Left())
	}
	return json.Marshal(e.Right())
}
