package util

import "strings"

type JSONStructTag struct {
	Name      string
	OmitEmpty bool
	String    bool
}

func FromString(s string) JSONStructTag {
	split := strings.Split(s, ",")
	ret := JSONStructTag{
		Name: split[0],
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
