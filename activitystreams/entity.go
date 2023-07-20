package activitystreams

import (
	"encoding/json"
	"errors"
	"fmt"
)

type EntityIface interface {
	json.Marshaler
	UnmarshalEntity(*EntityUnmarshaler, []byte) error
	Type() (string, error)
	entity() *Entity
}

// Get a reference to the Entity struct embedded in the given EntityIface
func ToEntity(e EntityIface) *Entity {
	return e.entity()
}

type Entity struct {
	dbId         uint
	Id           string
	AttributedTo []EntityIface
	Name         string
	MediaType    string
	Preview      EntityIface
}

func (e *Entity) entity() *Entity {
	return e
}

func (e *Entity) Type() (string, error) {
	return "", errors.New("Untyped ActivityStreams Entity")
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

func (e *Entity) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	var idString string
	err := json.Unmarshal(b, &idString)
	if err == nil {
		e.Id = idString
		return nil
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if id, ok := objMap["id"]; ok {
		err = json.Unmarshal(id, &e.Id)
		if err != nil {
			return err
		}
	}

	if attributedTo, ok := objMap["attributedTo"]; ok {
		var rawAttributions []json.RawMessage
		err = json.Unmarshal(attributedTo, &rawAttributions)
		if err != nil {
			return err
		}
		e.AttributedTo = make([]EntityIface, 0, len(rawAttributions))
		for _, item := range rawAttributions {
			var attributedToEntity Entity
			err = attributedToEntity.UnmarshalEntity(u, item)
			if err != nil {
				return err
			}
			e.AttributedTo = append(e.AttributedTo, &attributedToEntity)
		}
	}

	if name, ok := objMap["name"]; ok {
		err = json.Unmarshal(name, &e.Name)
		if err != nil {
			return err
		}
	}

	if mediaType, ok := objMap["mediaType"]; ok {
		err = json.Unmarshal(mediaType, &e.MediaType)
		if err != nil {
			return err
		}
	}

	return nil
}
