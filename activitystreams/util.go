package activitystreams

import (
	"encoding/json"
	"errors"
)

type EntityUnmarshaler struct {
	unmarshalFnByType map[string]unmarshalFn
}

type unmarshalFn func([]byte) EntityIface

func defaultUnmarshalFn[E EntityIface](b []byte) EntityIface {
	var e E
	json.Unmarshal(b, e)
	return e
}

var DefaultEntityUnmarshaler = EntityUnmarshaler{
	unmarshalFnByType: map[string]unmarshalFn{
		"Activity":        defaultUnmarshalFn[*Activity],
		"Accept":          defaultUnmarshalFn[*Accept],
		"Add":             defaultUnmarshalFn[*Add],
		"Arrive":          defaultUnmarshalFn[*Arrive],
		"Block":           defaultUnmarshalFn[*Block],
		"Create":          defaultUnmarshalFn[*Create],
		"Delete":          defaultUnmarshalFn[*Delete],
		"Dislike":         defaultUnmarshalFn[*Dislike],
		"Flag":            defaultUnmarshalFn[*Flag],
		"Follow":          defaultUnmarshalFn[*Follow],
		"Ignore":          defaultUnmarshalFn[*Ignore],
		"Invite":          defaultUnmarshalFn[*Invite],
		"Join":            defaultUnmarshalFn[*Join],
		"Leave":           defaultUnmarshalFn[*Leave],
		"Like":            defaultUnmarshalFn[*Like],
		"Listen":          defaultUnmarshalFn[*Listen],
		"Move":            defaultUnmarshalFn[*Move],
		"Offer":           defaultUnmarshalFn[*Offer],
		"Question":        defaultUnmarshalFn[*Question],
		"Reject":          defaultUnmarshalFn[*Reject],
		"Read":            defaultUnmarshalFn[*Read],
		"Remove":          defaultUnmarshalFn[*Remove],
		"TentativeAccept": defaultUnmarshalFn[*TentativeAccept],
		"TentativeReject": defaultUnmarshalFn[*TentativeReject],
		"Undo":            defaultUnmarshalFn[*Undo],
		"Update":          defaultUnmarshalFn[*Update],
		"View":            defaultUnmarshalFn[*View],
		"Application":     defaultUnmarshalFn[*Application],
		"Group":           defaultUnmarshalFn[*Group],
		"Organization":    defaultUnmarshalFn[*Organization],
		"Person":          defaultUnmarshalFn[*Person],
		"Service":         defaultUnmarshalFn[*Service],
		"Article":         defaultUnmarshalFn[*Article],
		"Audio":           defaultUnmarshalFn[*Audio],
		"Document":        defaultUnmarshalFn[*Document],
		"Event":           defaultUnmarshalFn[*Event],
		"Image":           defaultUnmarshalFn[*Image],
		"Note":            defaultUnmarshalFn[*Note],
		"Object":          defaultUnmarshalFn[*Object],
		"Page":            defaultUnmarshalFn[*Page],
		"Place":           defaultUnmarshalFn[*Place],
		"Profile":         defaultUnmarshalFn[*Profile],
		"Relationship":    defaultUnmarshalFn[*Relationship],
		"Tombstone":       defaultUnmarshalFn[*Tombstone],
		"Video":           defaultUnmarshalFn[*Video],
		"Collection":      defaultUnmarshalFn[*Collection],
		"CollectionPage":  defaultUnmarshalFn[*CollectionPage],
		"Link":            defaultUnmarshalFn[*Link],
		"Mention":         defaultUnmarshalFn[*Mention],
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
	return fn(b), nil
}
