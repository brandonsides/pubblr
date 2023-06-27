package activitystreams

import (
	"encoding/json"
	"errors"
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

// Represents an entity at the top level of an ActivityStreams document,
// including the @context field.
type TopLevelEntity struct {
	EntityIface
	Context string `json:"@context,omitempty"`
}

func (t *TopLevelEntity) MarshalJSON() ([]byte, error) {
	return MarshalEntity(t)
}

func (t *TopLevelEntity) Type() (string, error) {
	if t.EntityIface != nil {
		return t.EntityIface.Type()
	}
	return "", errors.New("No EntityIface set on TopLevelEntity")
}
