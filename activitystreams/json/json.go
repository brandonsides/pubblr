package json

import (
	"reflect"
	"strings"
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
