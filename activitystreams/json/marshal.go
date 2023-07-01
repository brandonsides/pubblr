package json

import (
	"encoding/json"
	"reflect"

	"github.com/brandonsides/pubblr/activitystreams/entity"
)

// Marshal an EntityIface to JSON
// Marshals the implementing type to JSON and adds a "type" field to the JSON
// representation with the value returned by the Type() method.
func MarshalEntity(e entity.EntityIface) ([]byte, error) {
	entMap := make(map[string]interface{})

	eElemType := reflect.TypeOf(e).Elem()
	for fieldIndex := 0; fieldIndex < eElemType.NumField(); fieldIndex++ {
		field := eElemType.Field(fieldIndex)
		if field.Anonymous {
			fieldInterface := reflect.ValueOf(e).Elem().Field(fieldIndex).Interface()
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
				// Don't overwrite existing values
				if _, ok := entMap[k]; !ok {
					entMap[k] = v
				}
			}
			continue
		}
		tag := TagFromStructField(field)
		if tag.Omit || tag.OmitEmpty && reflect.ValueOf(e).Elem().Field(fieldIndex).IsZero() {
			continue
		}

		v := reflect.ValueOf(e).Elem().Field(fieldIndex)
		if tag.String {
			entMap[tag.Name] = v.String()
		} else {
			entMap[tag.Name] = v.Interface()
		}
	}

	if eType, err := e.Type(); err == nil {
		entMap["type"] = eType
	} else {
		return nil, err
	}

	return json.Marshal(entMap)
}
