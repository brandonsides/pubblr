package activitystreams

import (
	"errors"

	"github.com/brandonsides/pubblr/util/json"
)

// Represents an entity at the top level of an ActivityStreams document,
// including the @context field.
type TopLevelEntity struct {
	EntityIface
	Context string `json:"@context,omitempty"`
}

type JsonTopLevelEntity TopLevelEntity

func (t *TopLevelEntity) MarshalJSON() ([]byte, error) {
	return MarshalEntity(t)
}

func (t *TopLevelEntity) Type() (string, error) {
	if t.EntityIface != nil {
		return t.EntityIface.Type()
	}
	return "", errors.New("No EntityIface set on TopLevelEntity")
}

func (t *TopLevelEntity) CustomUnmarshalJSON(u json.CustomUnmarshaler, b []byte) error {
	err := u.Unmarshal(b, (*JsonTopLevelEntity)(t))
	if err != nil {
		return err
	}

	if t.Context == "" {
		return errors.New("No @context field")
	}
	return nil
}
