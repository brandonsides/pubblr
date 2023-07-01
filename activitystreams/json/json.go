package json

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/brandonsides/pubblr/activitystreams/entity"
)

type JSONStructTag struct {
	Name      string
	OmitEmpty bool
	String    bool
	Omit      bool
}

func TagFromStructField(f reflect.StructField) JSONStructTag {
	tag := f.Tag.Get("json")
	if tag == "-" {
		return JSONStructTag{
			Omit: true,
		}
	}
	split := strings.Split(tag, ",")
	ret := JSONStructTag{
		Name: split[0],
	}
	if ret.Name == "" {
		ret.Name = f.Name
	}
	for _, opt := range split[1:] {
		switch opt {
		case "omitempty":
			ret.OmitEmpty = true
		case "string":
			ret.String = true
		}
	}
	return ret
}

func (u *InterfaceUnmarshaler) UnmarshalInterface(b []byte) (interface{}, error) {
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
	fn, ok := u.unmarshalFnByType[t]
	if !ok {
		return nil, errors.New("no unmarshal function for type: " + t)
	}
	return fn(u, b)
}

func (e *InterfaceUnmarshaler) RegisterUnmarshalFn(t string, fn unmarshalFn) {
	if e.unmarshalFnByType == nil {
		e.unmarshalFnByType = make(map[string]unmarshalFn)
	}
	e.unmarshalFnByType[t] = fn
}

func (u *InterfaceUnmarshaler) RegisterType(t string, e entity.EntityIface) {
	u.RegisterUnmarshalFn(t, defaultUnmarshalFn(e))
}
