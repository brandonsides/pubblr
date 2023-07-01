package entity

import (
	"encoding/json"
)

type EntityIface interface {
	json.Marshaler
	Type() (string, error)
	entity() *Entity
}

// Get a reference to the Entity struct embedded in the given EntityIface
func ToEntity(e EntityIface) *Entity {
	return e.entity()
}

type Entity struct {
	Id           string        `json:"id,omitempty"`
	AttributedTo []EntityIface `json:"attributedTo,omitempty"`
	Name         string        `json:"name,omitempty"`
	MediaType    string        `json:"mediaType,omitempty"`
}

func (e *Entity) entity() *Entity {
	return e
}
