package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/brandonsides/pubblr/activitystreams/entity"
)

type CustomUnmarshaler interface {
	Unmarshal([]byte, interface{}) error
}

type CustomUnmarshalerUser interface {
	CustomUnmarshalJSON(CustomUnmarshaler, []byte) error
}

type InterfaceUnmarshaler struct {
	unmarshalFnByType map[string]unmarshalFn
}

func (u *InterfaceUnmarshaler) Unmarshal(b []byte, dest interface{}) error {
	targetType := reflect.TypeOf(dest)
	if targetType.Kind() != reflect.Ptr {
		return errors.New("dest must be a pointer")
	}
	targetType = targetType.Elem()

	jsonUnmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	customUnmarshalerType := reflect.TypeOf((*CustomUnmarshalerUser)(nil)).Elem()
	bad := map[reflect.Kind]bool{
		reflect.Chan:      true,
		reflect.Func:      true,
		reflect.Interface: true,
	}
	var err error
	if reflect.TypeOf(dest).Implements(customUnmarshalerType) {
		dest.(CustomUnmarshalerUser).CustomUnmarshalJSON(u, b)
	} else if reflect.TypeOf(dest).Implements(jsonUnmarshalerType) {
		err = json.Unmarshal(b, dest)
	} else if targetType.Kind() == reflect.Interface {
		val, err := u.UnmarshalInterface(b)
		if err != nil {
			return err
		}
		reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(val))
	} else if targetType.Kind() == reflect.Struct {
		err = u.unmarshalStruct(b, dest)
		if err != nil {
			return err
		}
	} else if targetType.Kind() == reflect.Ptr {
		subDest := reflect.ValueOf(dest).Elem()
		subDest.Set(reflect.New(targetType.Elem()))
		err = u.Unmarshal(b, subDest.Interface())
	} else if targetType.Kind() == reflect.Slice || targetType.Kind() == reflect.Array {
		var unmarshalledSlc []json.RawMessage
		err := json.Unmarshal(b, &unmarshalledSlc)
		if err != nil {
			return err
		}
		slc := reflect.MakeSlice(targetType, 0, len(unmarshalledSlc))
		for _, val := range unmarshalledSlc {
			unmarshalledVal := reflect.New(targetType.Elem()).Interface()
			err := u.Unmarshal(val, unmarshalledVal)
			if err != nil {
				return err
			}
			slc = reflect.Append(slc, reflect.ValueOf(unmarshalledVal).Elem())
		}
		reflect.ValueOf(dest).Elem().Set(slc)
	} else if targetType.Kind() == reflect.Map && targetType.Key() == reflect.TypeOf("") {
		var unmarshalledMap map[string]json.RawMessage
		err := json.Unmarshal(b, &unmarshalledMap)
		if err != nil {
			return err
		}
		m := reflect.MakeMap(targetType)
		for k, v := range unmarshalledMap {
			unmarshalledVal := reflect.New(targetType.Elem()).Interface()
			err := u.Unmarshal(v, unmarshalledVal)
			if err != nil {
				return err
			}
			m.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(unmarshalledVal).Elem())
		}
		reflect.ValueOf(dest).Elem().Set(m)
	} else if !bad[targetType.Kind()] {
		err = json.Unmarshal(b, dest)
	} else {
		return errors.New("Unsupported type: " + targetType.String())
	}
	return err
}

type unmarshalFn func(*InterfaceUnmarshaler, []byte) (interface{}, error)

func defaultUnmarshalFn(e entity.EntityIface) unmarshalFn {
	targetType := reflect.TypeOf(e)
	return func(u *InterfaceUnmarshaler, b []byte) (interface{}, error) {
		ret := reflect.New(targetType)
		jsonUnmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()

		var err error
		if targetType.Implements(jsonUnmarshalerType) {
			err = ret.Interface().(json.Unmarshaler).UnmarshalJSON(b)
		} else {
			err = u.Unmarshal(b, ret.Interface())
		}
		if err != nil {
			return nil, err
		}

		entRet, ok := ret.Elem().Interface().(entity.EntityIface)
		if !ok {
			return nil, fmt.Errorf("unmarshal: expected activitystreams.EntityIface, got %T", ret)
		}

		return entRet, nil
	}
}

func (u *InterfaceUnmarshaler) unmarshalStruct(b []byte, dest interface{}) error {
	targetType := reflect.TypeOf(dest)
	if targetType.Kind() != reflect.Ptr {
		return errors.New("dest must be a pointer")
	}
	targetType = targetType.Elem()
	if targetType.Kind() != reflect.Struct {
		return errors.New("dest must be a pointer to a struct")
	}

	var raw map[string]json.RawMessage
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		var fieldBytes json.RawMessage
		if field.Anonymous {
			fieldBytes = b
		} else {
			jsonTag := TagFromStructField(field)
			if jsonTag.Omit {
				continue
			}
			var ok bool
			fieldBytes, ok = raw[jsonTag.Name]
			if !ok {
				if jsonTag.OmitEmpty {
					continue
				} else {
					return errors.New("JSON does not include required field: " + jsonTag.Name)
				}
			}
		}
		fieldDest := reflect.ValueOf(dest).Elem().Field(i).Addr().Interface()
		err := u.Unmarshal(fieldBytes, fieldDest)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *InterfaceUnmarshaler) unmarshalAs(t string, b []byte) (interface{}, error) {
	fn, ok := u.unmarshalFnByType[t]
	if !ok {
		return nil, errors.New("no unmarshal function for type: " + t)
	}
	return fn(u, b)
}