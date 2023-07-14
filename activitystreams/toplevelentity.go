package activitystreams

import (
	"errors"
	"fmt"

	"encoding/json"

	jsonutil "github.com/brandonsides/pubblr/util/json"
)

// Represents an entity at the top level of an ActivityStreams document,
// including the @context field.
type TopLevelEntity struct {
	EntityIface
	Context string `json:"@context,omitempty"`
}

type JsonTopLevelEntity TopLevelEntity

func (t *TopLevelEntity) MarshalJSON() ([]byte, error) {
	if t.EntityIface == nil {
		return nil, errors.New("No EntityIface set on TopLevelEntity")
	}
	topLevelEntity, err := json.Marshal(t.EntityIface)
	if err != nil {
		return nil, err
	}

	var topLevelEntityMap map[string]json.RawMessage
	err = json.Unmarshal(topLevelEntity, &topLevelEntityMap)
	if err != nil {
		return nil, err
	}

	if t.Context != "" {
		topLevelEntityMap["@context"] = []byte(fmt.Sprintf("%q", t.Context))
	}

	return json.Marshal(topLevelEntityMap)
}

func (t *TopLevelEntity) Type() (string, error) {
	if t.EntityIface != nil {
		return t.EntityIface.Type()
	}
	return "", errors.New("No EntityIface set on TopLevelEntity")
}

func (t *TopLevelEntity) CustomUnmarshalJSON(u jsonutil.CustomUnmarshaler, b []byte) error {
	err := u.Unmarshal(b, (*JsonTopLevelEntity)(t))
	if err != nil {
		return err
	}

	if t.Context == "" {
		return errors.New("No @context field")
	}
	return nil
}
