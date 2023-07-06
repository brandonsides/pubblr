package either

import (
	"encoding/json"
	"fmt"

	jsonutil "github.com/brandonsides/pubblr/util/json"
)

type Either[A, B any] struct {
	isLeft bool
	a      *A
	b      *B
}

func Left[A, B any](a A) *Either[A, B] {
	return &Either[A, B]{
		isLeft: true,
		a:      &a,
	}
}

func Right[A, B any](b B) *Either[A, B] {
	return &Either[A, B]{
		isLeft: false,
		b:      &b,
	}
}

func (e Either[A, B]) Left() *A {
	return e.a
}

func (e Either[A, B]) Right() *B {
	return e.b
}

func (e Either[A, B]) MarshalJSON() ([]byte, error) {
	if e.isLeft {
		return json.Marshal(e.Left())
	}
	return json.Marshal(e.Right())
}

func (e *Either[A, B]) UnmarshalJSON(data []byte) error {
	var a A
	var b B

	if err := json.Unmarshal(data, &a); err == nil {
		*e = *Left[A, B](a)
		return nil
	}
	if err := json.Unmarshal(data, &b); err == nil {
		*e = *Right[A](b)
		return nil
	}
	return fmt.Errorf("Could not unmarshal Either[%T, %T]", a, b)
}

func (e *Either[A, B]) CustomUnmarshalJSON(u jsonutil.CustomUnmarshaler, data []byte) error {
	var a A
	var b B

	if err := u.Unmarshal(data, &a); err == nil {
		*e = *Left[A, B](a)
		return nil
	}
	if err := u.Unmarshal(data, &b); err == nil {
		*e = *Right[A](b)
		return nil
	}
	return fmt.Errorf("Could not unmarshal Either[%T, %T]", a, b)
}
