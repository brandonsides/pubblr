package activitystreams

import (
	"encoding/json"

	jsonutil "github.com/brandonsides/pubblr/util/json"
)

var DefaultEntityUnmarshaler jsonutil.InterfaceUnmarshaler

func init() {
	DefaultEntityUnmarshaler.RegisterUnmarshalFn("Question", func(u *jsonutil.InterfaceUnmarshaler, b []byte) (interface{}, error) {
		var qMap map[string]interface{}
		json.Unmarshal(b, &qMap)
		if _, ok := qMap["oneOf"]; ok {
			ret := SingleAnswerQuestion{}
			err := u.Unmarshal(b, &ret)
			return &ret, err
		} else if _, ok := qMap["anyOf"]; ok {
			ret := MultiAnswerQuestion{}
			err := u.Unmarshal(b, &ret)
			return &ret, err
		} else if _, ok := qMap["closed"]; ok {
			ret := ClosedQuestion{}
			err := u.Unmarshal(b, &ret)
			return &ret, err
		}
		ret := Question{}
		err := u.Unmarshal(b, &ret)
		return &ret, err
	})
	DefaultEntityUnmarshaler.RegisterType("IntransitiveActivity", &IntransitiveActivity{})
	DefaultEntityUnmarshaler.RegisterType("Activity", &Activity{})
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
