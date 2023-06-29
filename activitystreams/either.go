package activitystreams

import (
	"encoding/json"
	"fmt"
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

func (e *Either[A, B]) CustomUnmarshalJSON(u *EntityUnmarshaler, data []byte) error {
	var a A
	var b B

	var aIface interface{} = a
	if err := u.Unmarshal(data, &aIface); err == nil {
		*e = *Left[A, B](a)
		return nil
	}
	var bIface interface{} = b
	if err := json.Unmarshal(data, &bIface); err == nil {
		*e = *Right[A](b)
		return nil
	}
	return fmt.Errorf("Could not unmarshal Either[%T, %T]", a, b)
}