package activitystreams

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/brandonsides/pubblr/util"
)

var DefaultEntityUnmarshaler EntityUnmarshaler

func init() {
	DefaultEntityUnmarshaler = EntityUnmarshaler{
		unmarshalFnByType: map[string]unmarshalFn{
			"Activity": defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Activity{}),
			"Accept":   defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Accept{}),
			"Add":      defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Add{}),
			"Arrive":   defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Arrive{}),
			"Block":    defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Block{}),
			"Create":   defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Create{}),
			"Delete":   defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Delete{}),
			"Dislike":  defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Dislike{}),
			"Flag":     defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Flag{}),
			"Follow":   defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Follow{}),
			"Ignore":   defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Ignore{}),
			"Invite":   defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Invite{}),
			"Join":     defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Join{}),
			"Leave":    defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Leave{}),
			"Like":     defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Like{}),
			"Listen":   defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Listen{}),
			"Move":     defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Move{}),
			"Offer":    defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Offer{}),
			"Question": func(b []byte) (EntityIface, error) {
				var qMap map[string]interface{}
				json.Unmarshal(b, &qMap)
				if _, ok := qMap["oneOf"]; ok {
					ret := SingleAnswerQuestion{}
					err := json.Unmarshal(b, &ret)
					return &ret, err
				} else if _, ok := qMap["anyOf"]; ok {
					ret := MultiAnswerQuestion{}
					err := json.Unmarshal(b, &ret)
					return &ret, err
				} else if _, ok := qMap["closed"]; ok {
					ret := ClosedQuestion{}
					err := json.Unmarshal(b, &ret)
					return &ret, err
				}
				return nil, errors.New("Unknown question type")
			},
			"Reject":          defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Reject{}),
			"Read":            defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Read{}),
			"Remove":          defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Remove{}),
			"TentativeAccept": defaultUnmarshalFn(&DefaultEntityUnmarshaler, &TentativeAccept{}),
			"TentativeReject": defaultUnmarshalFn(&DefaultEntityUnmarshaler, &TentativeReject{}),
			"Undo":            defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Undo{}),
			"Update":          defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Update{}),
			"View":            defaultUnmarshalFn(&DefaultEntityUnmarshaler, &View{}),
			"Application":     defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Application{}),
			"Group":           defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Group{}),
			"Organization":    defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Organization{}),
			"Person":          defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Person{}),
			"Service":         defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Service{}),
			"Article":         defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Article{}),
			"Audio":           defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Audio{}),
			"Document":        defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Document{}),
			"Event":           defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Event{}),
			"Image":           defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Image{}),
			"Note":            defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Note{}),
			"Object":          defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Object{}),
			"Page":            defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Page{}),
			"Place":           defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Place{}),
			"Profile":         defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Profile{}),
			"Relationship":    defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Relationship{}),
			"Tombstone":       defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Tombstone{}),
			"Video":           defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Video{}),
			"Collection":      defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Collection{}),
			"CollectionPage":  defaultUnmarshalFn(&DefaultEntityUnmarshaler, &CollectionPage{}),
			"Link":            defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Link{}),
			"Mention":         defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Mention{}),
		},
	}
}

func isEntityType(t reflect.Type) bool {
	EntityIfaceType := reflect.TypeOf((*EntityIface)(nil)).Elem()
	EntityType := reflect.TypeOf(&Entity{})
	return t.Implements(EntityIfaceType) || t == EntityType
}

func isEmbeddedEntityField(f *reflect.StructField) bool {
	return f.Anonymous && (isEntityType(f.Type) || isEntityType(reflect.PointerTo(f.Type)))
}

// Marhsal an EntityIface to JSON
// Marshals the implementing type to JSON and adds a "type" field to the JSON
// representation with the value returned by the Type() method.
func MarshalEntity(e EntityIface) ([]byte, error) {
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
		tag := util.FromStructField(field)
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

type EntityUnmarshaler struct {
	unmarshalFnByType map[string]unmarshalFn
}

type unmarshalFn func([]byte) (EntityIface, error)

var entityIface reflect.Type = reflect.TypeOf((*EntityIface)(nil)).Elem()

func defaultUnmarshalFn(u *EntityUnmarshaler, e EntityIface) unmarshalFn {
	targetType := reflect.TypeOf(e).Elem()
	return func(b []byte) (EntityIface, error) {
		mapped := make(map[string]json.RawMessage)
		err := json.Unmarshal(b, &mapped)
		if err != nil {
			return nil, err
		}

		reflRet := reflect.New(targetType)
		retType := reflRet.Type()
		for i := 0; i < retType.NumField(); i++ {
			field := retType.Field(i)
			jsonTag := util.FromStructField(field)
			if jsonTag.Omit {
				continue
			}

			if field.Type.Kind() == reflect.Interface && field.Type.Implements(entityIface) {
				if val, ok := mapped[jsonTag.Name]; ok {
					unmarshalledVal, err := u.Unmarshal(val)
					if err != nil {
						return nil, err
					}
					reflRet.Elem().Field(i).Set(reflect.ValueOf(unmarshalledVal))
				} else {
					return nil, errors.New("Missing field: " + jsonTag.Name)
				}
			} else {
				if val, ok := mapped[jsonTag.Name]; ok {
					err := json.Unmarshal(val, reflRet.Elem().Field(i).Addr().Interface())
					if err != nil {
						return nil, err
					}
				}
			}
		}

		return reflRet.Interface().(EntityIface), nil
	}
}

func (e *EntityUnmarshaler) Unmarshal(b []byte) (EntityIface, error) {
	var raw map[string]interface{}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return nil, err
	}
	t, ok := raw["type"]
	if !ok {
		return nil, errors.New("no type field")
	}
	tStr := t.(string)
	if !ok {
		return nil, errors.New("type is not a string")
	}
	fn, ok := e.unmarshalFnByType[tStr]
	if !ok {
		return nil, errors.New("no unmarshal function for type " + tStr)
	}
	return fn(b)
}

func (e *EntityUnmarshaler) RegisterUnmarshalFn(t string, fn unmarshalFn) {
	e.unmarshalFnByType[t] = fn
}

func (u *EntityUnmarshaler) RegisterType(e EntityIface, t string) {
	u.RegisterUnmarshalFn(t, defaultUnmarshalFn(u, e))
}
