package activitystreams

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/brandonsides/pubblr/util"
)

var DefaultEntityUnmarshaler EntityUnmarshaler

func init() {
	DefaultEntityUnmarshaler = EntityUnmarshaler{
		unmarshalFnByType: map[string]unmarshalFn{
			"IntransitiveActivity": defaultUnmarshalFn(&DefaultEntityUnmarshaler, &IntransitiveActivity{}),
			"Activity":             defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Activity{}),
			"Accept":               defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Accept{}),
			"Announce":             defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Announce{}),
			"Add":                  defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Add{}),
			"Arrive":               defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Arrive{}),
			"Block":                defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Block{}),
			"Create":               defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Create{}),
			"Delete":               defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Delete{}),
			"Dislike":              defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Dislike{}),
			"Flag":                 defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Flag{}),
			"Follow":               defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Follow{}),
			"Ignore":               defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Ignore{}),
			"Invite":               defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Invite{}),
			"Join":                 defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Join{}),
			"Leave":                defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Leave{}),
			"Like":                 defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Like{}),
			"Listen":               defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Listen{}),
			"Move":                 defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Move{}),
			"Offer":                defaultUnmarshalFn(&DefaultEntityUnmarshaler, &Offer{}),
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

// Marshal an EntityIface to JSON
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
	targetType := reflect.TypeOf(e)
	return func(b []byte) (EntityIface, error) {
		ret, err := unmarshal(u, targetType, b)
		if err != nil {
			return nil, err
		}

		entRet, ok := ret.(EntityIface)
		if !ok {
			return nil, fmt.Errorf("unmarshal: expected EntityIface, got %T", ret)
		}

		return entRet, nil
	}
}

func unmarshal(u *EntityUnmarshaler, targetType reflect.Type, b []byte) (interface{}, error) {
	jsonUnmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	bad := map[reflect.Kind]bool{
		reflect.Chan:      true,
		reflect.Func:      true,
		reflect.Interface: true,
	}
	var ret interface{}
	var err error
	if targetType == entityIface {
		ret, err = u.Unmarshal(b)
	} else if targetType.Implements(jsonUnmarshalerType) {
		err = json.Unmarshal(b, &ret)
	} else if targetType.Kind() == reflect.Struct {
		ret, err = unmarshalStruct(u, targetType, b)
	} else if targetType.Kind() == reflect.Ptr {
		unmarshalledVal, err := unmarshal(u, targetType.Elem(), b)
		if err != nil {
			return nil, err
		}
		ptr := reflect.New(targetType.Elem())
		ptr.Elem().Set(reflect.ValueOf(unmarshalledVal))
		ret = ptr.Interface()
	} else if targetType.Kind() == reflect.Slice || targetType.Kind() == reflect.Array {
		var unmarshalledSlc []json.RawMessage
		err := json.Unmarshal(b, &unmarshalledSlc)
		if err != nil {
			return nil, err
		}
		slc := reflect.MakeSlice(targetType, 0, len(unmarshalledSlc))
		for _, val := range unmarshalledSlc {
			unmarshalledVal, err := unmarshal(u, targetType.Elem(), val)
			if err != nil {
				return nil, err
			}
			slc = reflect.Append(slc, reflect.ValueOf(unmarshalledVal))
		}
		ret = slc.Interface()
	} else if targetType.Kind() == reflect.Map && targetType.Key() == reflect.TypeOf("") {
		var unmarshalledMap map[string]json.RawMessage
		err := json.Unmarshal(b, &unmarshalledMap)
		if err != nil {
			return nil, err
		}
		m := reflect.MakeMap(targetType)
		for k, v := range unmarshalledMap {
			unmarshalledVal, err := unmarshal(u, targetType.Elem(), v)
			if err != nil {
				return nil, err
			}
			m.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(unmarshalledVal))
		}
		ret = m.Interface()
	} else if !bad[targetType.Kind()] {
		err = json.Unmarshal(b, &ret)
	} else {
		return nil, errors.New("Unsupported type: " + targetType.String())
	}
	return ret, err
}

func unmarshalStruct(u *EntityUnmarshaler, targetType reflect.Type, b []byte) (interface{}, error) {
	if targetType.Kind() != reflect.Struct {
		return nil, errors.New("targetType is not a struct")
	}

	var raw map[string]json.RawMessage
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return nil, err
	}

	ret := reflect.New(targetType).Elem()
	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		var fieldBytes json.RawMessage
		if field.Anonymous {
			fieldBytes = b
		} else {
			jsonTag := util.FromStructField(field)
			var ok bool
			fieldBytes, ok = raw[jsonTag.Name]
			if !ok {
				if jsonTag.OmitEmpty {
					continue
				} else {
					return nil, errors.New("JSON does not include required field: " + jsonTag.Name)
				}
			}
		}
		val, err := unmarshal(u, field.Type, fieldBytes)
		if err != nil {
			return nil, err
		}
		ret.Field(i).Set(reflect.ValueOf(val))
	}

	return ret.Interface(), nil
}

func (e *EntityUnmarshaler) Unmarshal(b []byte) (EntityIface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return nil, err
	}
	var t string
	tRaw, ok := raw["type"]
	if !ok {
		return nil, errors.New("no type field")
	}
	err = json.Unmarshal(tRaw, &t)
	if err != nil {
		return nil, err
	}
	fn, ok := e.unmarshalFnByType[t]
	if !ok {
		return nil, errors.New("no unmarshal function for type: " + t)
	}
	return fn(b)
}

func (e *EntityUnmarshaler) UnmarshalAs(t string, b []byte) (EntityIface, error) {
	fn, ok := e.unmarshalFnByType[t]
	if !ok {
		return nil, errors.New("no unmarshal function for type: " + t)
	}
	return fn(b)
}

func (e *EntityUnmarshaler) RegisterUnmarshalFn(t string, fn unmarshalFn) {
	e.unmarshalFnByType[t] = fn
}

func (u *EntityUnmarshaler) RegisterType(e EntityIface, t string) {
	u.RegisterUnmarshalFn(t, defaultUnmarshalFn(u, e))
}