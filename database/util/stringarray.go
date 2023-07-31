package util

import (
	"database/sql/driver"
	"net/url"
	"strings"
)

type StringArray []string

func (a *StringArray) Scan(value interface{}) error {
	str := value.(string)

	splitStr := strings.Split(str, ",")
	for _, v := range splitStr {
		v, err := url.QueryUnescape(v)
		if err != nil {
			return err
		}
		*a = append(*a, v)
	}

	return nil
}

func (a *StringArray) Value() (driver.Value, error) {
	var ret string
	for _, v := range *a {
		ret += url.QueryEscape(v) + ","
	}
	return ret, nil
}
