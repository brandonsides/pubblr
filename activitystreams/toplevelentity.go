package activitystreams

import (
	"errors"
	"fmt"

	"encoding/json"
)

// Represents an entity at the top level of an ActivityStreams document,
// including the @context field.
type TopLevelEntity struct {
	EntityIface
	Context string `json:"@context,omitempty"`
}

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

func (t *TopLevelEntity) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	var err error
	t.EntityIface, err = u.UnmarshalEntity(b)
	if err != nil {
		return err
	}

	var topLevelEntityMap map[string]json.RawMessage
	err = json.Unmarshal(b, &topLevelEntityMap)
	if err != nil {
		return err
	}

	if context, ok := topLevelEntityMap["@context"]; ok {
		err = json.Unmarshal(context, &t.Context)
		if err != nil {
			return err
		}
	}

	return nil
}
