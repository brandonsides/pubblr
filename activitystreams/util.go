package activitystreams

import (
	"encoding/json"
	"errors"
)

type EntityUnmarshaler struct {
	unmarshalFnByType map[string]unmarshalFn
}

type unmarshalFn func([]byte) (EntityIface, error)

func defaultUnmarshalFn(e EntityIface) unmarshalFn {
	return func(b []byte) (EntityIface, error) {
		err := json.Unmarshal(b, e)
		return e, err
	}
}

var DefaultEntityUnmarshaler = EntityUnmarshaler{
	unmarshalFnByType: map[string]unmarshalFn{
		"Activity": defaultUnmarshalFn(&Activity{}),
		"Accept":   defaultUnmarshalFn(&Accept{}),
		"Add":      defaultUnmarshalFn(&Add{}),
		"Arrive":   defaultUnmarshalFn(&Arrive{}),
		"Block":    defaultUnmarshalFn(&Block{}),
		"Create":   defaultUnmarshalFn(&Create{}),
		"Delete":   defaultUnmarshalFn(&Delete{}),
		"Dislike":  defaultUnmarshalFn(&Dislike{}),
		"Flag":     defaultUnmarshalFn(&Flag{}),
		"Follow":   defaultUnmarshalFn(&Follow{}),
		"Ignore":   defaultUnmarshalFn(&Ignore{}),
		"Invite":   defaultUnmarshalFn(&Invite{}),
		"Join":     defaultUnmarshalFn(&Join{}),
		"Leave":    defaultUnmarshalFn(&Leave{}),
		"Like":     defaultUnmarshalFn(&Like{}),
		"Listen":   defaultUnmarshalFn(&Listen{}),
		"Move":     defaultUnmarshalFn(&Move{}),
		"Offer":    defaultUnmarshalFn(&Offer{}),
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
		"Reject":          defaultUnmarshalFn(&Reject{}),
		"Read":            defaultUnmarshalFn(&Read{}),
		"Remove":          defaultUnmarshalFn(&Remove{}),
		"TentativeAccept": defaultUnmarshalFn(&TentativeAccept{}),
		"TentativeReject": defaultUnmarshalFn(&TentativeReject{}),
		"Undo":            defaultUnmarshalFn(&Undo{}),
		"Update":          defaultUnmarshalFn(&Update{}),
		"View":            defaultUnmarshalFn(&View{}),
		"Application":     defaultUnmarshalFn(&Application{}),
		"Group":           defaultUnmarshalFn(&Group{}),
		"Organization":    defaultUnmarshalFn(&Organization{}),
		"Person":          defaultUnmarshalFn(&Person{}),
		"Service":         defaultUnmarshalFn(&Service{}),
		"Article":         defaultUnmarshalFn(&Article{}),
		"Audio":           defaultUnmarshalFn(&Audio{}),
		"Document":        defaultUnmarshalFn(&Document{}),
		"Event":           defaultUnmarshalFn(&Event{}),
		"Image":           defaultUnmarshalFn(&Image{}),
		"Note":            defaultUnmarshalFn(&Note{}),
		"Object":          defaultUnmarshalFn(&Object{}),
		"Page":            defaultUnmarshalFn(&Page{}),
		"Place":           defaultUnmarshalFn(&Place{}),
		"Profile":         defaultUnmarshalFn(&Profile{}),
		"Relationship":    defaultUnmarshalFn(&Relationship{}),
		"Tombstone":       defaultUnmarshalFn(&Tombstone{}),
		"Video":           defaultUnmarshalFn(&Video{}),
		"Collection":      defaultUnmarshalFn(&Collection{}),
		"CollectionPage":  defaultUnmarshalFn(&CollectionPage{}),
		"Link":            defaultUnmarshalFn(&Link{}),
		"Mention":         defaultUnmarshalFn(&Mention{}),
	},
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
	u.RegisterUnmarshalFn(t, defaultUnmarshalFn(e))
}
