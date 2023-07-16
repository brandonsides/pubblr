package activitystreams

import (
	"encoding/json"
	"errors"
	"reflect"
)

type UnmarshalFn func(*EntityUnmarshaler, []byte) (EntityIface, error)

type EntityUnmarshaler struct {
	unmarshalFnByType map[string]UnmarshalFn
}

func (u *EntityUnmarshaler) UnmarshalEntity(b []byte) (EntityIface, error) {
	var retMap map[string]interface{}
	err := json.Unmarshal(b, &retMap)
	if err != nil {
		var id string
		err := json.Unmarshal(b, &id)
		if err != nil {
			return nil, err
		}
		return &Entity{Id: id}, nil
	}

	typ, ok := retMap["type"].(string)
	if !ok {
		return nil, errors.New("Entity has no type")
	}

	fn, ok := u.unmarshalFnByType[typ]
	if !ok {
		return nil, errors.New("no unmarshal function for type: " + typ)
	}

	ret, err := fn(u, b)
	if err != nil {
		return nil, err
	}

	return ret.(EntityIface), nil
}

func (e *EntityUnmarshaler) RegisterUnmarshalFn(t string, fn UnmarshalFn) {
	if e.unmarshalFnByType == nil {
		e.unmarshalFnByType = make(map[string]UnmarshalFn)
	}
	e.unmarshalFnByType[t] = fn
}

func (u *EntityUnmarshaler) RegisterType(t string, e EntityIface) {
	u.RegisterUnmarshalFn(t, defaultUnmarshalFn(e))
}

func defaultUnmarshalFn(e EntityIface) UnmarshalFn {
	eType := reflect.TypeOf(e).Elem()
	return func(u *EntityUnmarshaler, b []byte) (EntityIface, error) {
		ret := reflect.New(eType).Interface().(EntityIface)
		err := ret.(EntityIface).UnmarshalEntity(u, b)
		if err != nil {
			return nil, err
		}
		return ret, nil
	}
}

var DefaultEntityUnmarshaler *EntityUnmarshaler = &EntityUnmarshaler{}

func init() {
	DefaultEntityUnmarshaler.RegisterUnmarshalFn("Question", func(u *EntityUnmarshaler, b []byte) (EntityIface, error) {
		var qMap map[string]interface{}
		err := json.Unmarshal(b, &qMap)
		if err != nil {
			return nil, err
		}

		var ret EntityIface
		if _, ok := qMap["oneOf"]; ok {
			ret = &SingleAnswerQuestion{}
			err = ret.UnmarshalEntity(u, b)
		} else if _, ok := qMap["anyOf"]; ok {
			ret = &MultiAnswerQuestion{}
			err = ret.UnmarshalEntity(u, b)
		} else if _, ok := qMap["closed"]; ok {
			ret = &ClosedQuestion{}
			err = ret.UnmarshalEntity(u, b)
		} else {
			ret = &Question{}
			err = ret.UnmarshalEntity(u, b)
		}
		return ret, err
	})
	DefaultEntityUnmarshaler.RegisterType("IntransitiveActivity", &IntransitiveActivity{})
	DefaultEntityUnmarshaler.RegisterType("Activity", &TransitiveActivity{})
	DefaultEntityUnmarshaler.RegisterType("Accept", &Accept{})
	DefaultEntityUnmarshaler.RegisterType("Announce", &Announce{})
	DefaultEntityUnmarshaler.RegisterType("Add", &Add{})
	DefaultEntityUnmarshaler.RegisterType("Arrive", &Arrive{})
	DefaultEntityUnmarshaler.RegisterType("Block", &Block{})
	DefaultEntityUnmarshaler.RegisterType("Create", &Create{})
	DefaultEntityUnmarshaler.RegisterType("Delete", &Delete{})
	DefaultEntityUnmarshaler.RegisterType("Dislike", &Dislike{})
	DefaultEntityUnmarshaler.RegisterType("Flag", &Flag{})
	DefaultEntityUnmarshaler.RegisterType("Follow", &Follow{})
	DefaultEntityUnmarshaler.RegisterType("Ignore", &Ignore{})
	DefaultEntityUnmarshaler.RegisterType("Invite", &Invite{})
	DefaultEntityUnmarshaler.RegisterType("Join", &Join{})
	DefaultEntityUnmarshaler.RegisterType("Leave", &Leave{})
	DefaultEntityUnmarshaler.RegisterType("Like", &Like{})
	DefaultEntityUnmarshaler.RegisterType("Listen", &Listen{})
	DefaultEntityUnmarshaler.RegisterType("Move", &Move{})
	DefaultEntityUnmarshaler.RegisterType("Offer", &Offer{})
	DefaultEntityUnmarshaler.RegisterType("Reject", &Reject{})
	DefaultEntityUnmarshaler.RegisterType("Read", &Read{})
	DefaultEntityUnmarshaler.RegisterType("Remove", &Remove{})
	DefaultEntityUnmarshaler.RegisterType("TentativeAccept", &TentativeAccept{})
	DefaultEntityUnmarshaler.RegisterType("TentativeReject", &TentativeReject{})
	DefaultEntityUnmarshaler.RegisterType("Travel", &Travel{})
	DefaultEntityUnmarshaler.RegisterType("Undo", &Undo{})
	DefaultEntityUnmarshaler.RegisterType("Update", &Update{})
	DefaultEntityUnmarshaler.RegisterType("View", &View{})
	DefaultEntityUnmarshaler.RegisterType("Application", &Application{})
	DefaultEntityUnmarshaler.RegisterType("Group", &Group{})
	DefaultEntityUnmarshaler.RegisterType("Organization", &Organization{})
	DefaultEntityUnmarshaler.RegisterType("Person", &Person{})
	DefaultEntityUnmarshaler.RegisterType("Service", &Service{})
	DefaultEntityUnmarshaler.RegisterType("Article", &Article{})
	DefaultEntityUnmarshaler.RegisterType("Audio", &Audio{})
	DefaultEntityUnmarshaler.RegisterType("Document", &Document{})
	DefaultEntityUnmarshaler.RegisterType("Event", &Event{})
	DefaultEntityUnmarshaler.RegisterType("Image", &Image{})
	DefaultEntityUnmarshaler.RegisterType("Note", &Note{})
	DefaultEntityUnmarshaler.RegisterType("Object", &Object{})
	DefaultEntityUnmarshaler.RegisterType("Page", &Page{})
	DefaultEntityUnmarshaler.RegisterType("Place", &Place{})
	DefaultEntityUnmarshaler.RegisterType("Profile", &Profile{})
	DefaultEntityUnmarshaler.RegisterType("Relationship", &Relationship{})
	DefaultEntityUnmarshaler.RegisterType("Tombstone", &Tombstone{})
	DefaultEntityUnmarshaler.RegisterType("Video", &Video{})
	DefaultEntityUnmarshaler.RegisterType("Collection", &Collection{})
	DefaultEntityUnmarshaler.RegisterType("CollectionPage", &CollectionPage{})
	DefaultEntityUnmarshaler.RegisterType("OrderedCollection", &Collection{})
	DefaultEntityUnmarshaler.RegisterType("OrderedCollectionPage", &CollectionPage{})
	DefaultEntityUnmarshaler.RegisterType("Link", &Link{})
	DefaultEntityUnmarshaler.RegisterType("Mention", &Mention{})
}

/*
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
*/
