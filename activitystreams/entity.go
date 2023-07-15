package activitystreams

import (
	"encoding/json"
	"fmt"
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

func (e *Entity) MarshalJSON() ([]byte, error) {
	entityMap := make(map[string]json.RawMessage)

	if e.Id != "" {
		entityMap["id"] = []byte(fmt.Sprintf("%q", e.Id))
	}

	if len(e.AttributedTo) > 0 {
		attributedToIds := make([]string, len(e.AttributedTo))
		for i, attributedTo := range e.AttributedTo {
			attributedToIds[i] = ToEntity(attributedTo).Id
		}
		attributedTo, err := json.Marshal(attributedToIds)
		if err != nil {
			return nil, err
		}
		entityMap["attributedTo"] = attributedTo
	}

	if e.Name != "" {
		entityMap["name"] = []byte(fmt.Sprintf("%q", e.Name))
	}

	if e.MediaType != "" {
		entityMap["mediaType"] = []byte(fmt.Sprintf("%q", e.MediaType))
	}

	return json.Marshal(entityMap)
}
