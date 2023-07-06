package activitystreams

import (
	"errors"
)

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
