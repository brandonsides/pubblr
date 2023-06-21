package activitystreams

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/brandonsides/pubblr/util"
)

type EntityIface interface {
	entity() *Entity
	Type() (string, error)
}

// Get a reference to the Entity struct embedded in the given EntityIface
func ToEntity(e EntityIface) *Entity {
	return e.entity()
}

func IsEmbeddedEntityField(f *reflect.StructField) bool {
	EntityIfaceType := reflect.TypeOf((*EntityIface)(nil)).Elem()
	return f.Anonymous && (f.Type == EntityIfaceType || reflect.PointerTo(f.Type).Implements(EntityIfaceType))
}

// Marhsal an EntityIface to JSON
// Marshals the implementing type to JSON and adds a "type" field to the JSON
// representation with the value returned by the Type() method.
func MarshalEntity(e EntityIface) ([]byte, error) {
	entMap := make(map[string]interface{})

	//EntityIfaceType := reflect.TypeOf((*EntityIface)(nil)).Elem()

	eElemType := reflect.TypeOf(e).Elem()
	for fieldIndex := 0; fieldIndex < eElemType.NumField(); fieldIndex++ {
		field := eElemType.Field(fieldIndex)
		if IsEmbeddedEntityField(&field) {
			fieldInterface := reflect.ValueOf(e).Elem().Field(fieldIndex).Interface()
			if entity, ok := fieldInterface.(Entity); ok {
				fieldInterface = (aliasedEntity)(entity)
			}
			var nestedMap map[string]interface{}
			nestedJson, err := json.Marshal(fieldInterface)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(nestedJson, &nestedMap)
			if err != nil {
				return nil, err
			}

			for k, v := range nestedMap {
				entMap[k] = v
			}
			continue
		}
		tag := util.FromString(field.Tag.Get("json"))
		if tag.Name == "" {
			tag.Name = strings.ToLower(field.Name[:1]) + field.Name[1:]
		}
		if tag.Name == "-" || tag.OmitEmpty && reflect.ValueOf(e).Elem().Field(fieldIndex).IsZero() {
			continue
		}

		v := reflect.ValueOf(e).Elem().Field(fieldIndex)
		if tag.String {
			entMap[tag.Name] = v.String()
		} else {
			entMap[tag.Name] = v.Interface()
		}
	}

	return json.Marshal(entMap)
}

type Entity struct {
	Id           string        `json:"id,omitempty"`
	AttributedTo []EntityIface `json:"attributedTo,omitempty"`
	Name         string        `json:"name,omitempty"`
	MediaType    string        `json:"mediaType,omitempty"`
}

type aliasedEntity Entity

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