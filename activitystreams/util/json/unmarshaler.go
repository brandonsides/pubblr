package json

import (
	"encoding/json"

	"github.com/brandonsides/pubblr/activitystreams/types"
	jsonutil "github.com/brandonsides/pubblr/util/json"
)

var DefaultEntityUnmarshaler jsonutil.InterfaceUnmarshaler

func init() {
	DefaultEntityUnmarshaler.RegisterUnmarshalFn("Question", func(u *jsonutil.InterfaceUnmarshaler, b []byte) (interface{}, error) {
		var qMap map[string]interface{}
		json.Unmarshal(b, &qMap)
		if _, ok := qMap["oneOf"]; ok {
			ret := types.SingleAnswerQuestion{}
			err := u.Unmarshal(b, &ret)
			return &ret, err
		} else if _, ok := qMap["anyOf"]; ok {
			ret := types.MultiAnswerQuestion{}
			err := u.Unmarshal(b, &ret)
			return &ret, err
		} else if _, ok := qMap["closed"]; ok {
			ret := types.ClosedQuestion{}
			err := u.Unmarshal(b, &ret)
			return &ret, err
		}
		ret := types.Question{}
		err := u.Unmarshal(b, &ret)
		return &ret, err
	})
	DefaultEntityUnmarshaler.RegisterType("IntransitiveActivity", &types.IntransitiveActivity{})
	DefaultEntityUnmarshaler.RegisterType("Activity", &types.Activity{})
	DefaultEntityUnmarshaler.RegisterType("Accept", &types.Accept{})
	DefaultEntityUnmarshaler.RegisterType("Announce", &types.Announce{})
	DefaultEntityUnmarshaler.RegisterType("Add", &types.Add{})
	DefaultEntityUnmarshaler.RegisterType("Arrive", &types.Arrive{})
	DefaultEntityUnmarshaler.RegisterType("Block", &types.Block{})
	DefaultEntityUnmarshaler.RegisterType("Create", &types.Create{})
	DefaultEntityUnmarshaler.RegisterType("Delete", &types.Delete{})
	DefaultEntityUnmarshaler.RegisterType("Dislike", &types.Dislike{})
	DefaultEntityUnmarshaler.RegisterType("Flag", &types.Flag{})
	DefaultEntityUnmarshaler.RegisterType("Follow", &types.Follow{})
	DefaultEntityUnmarshaler.RegisterType("Ignore", &types.Ignore{})
	DefaultEntityUnmarshaler.RegisterType("Invite", &types.Invite{})
	DefaultEntityUnmarshaler.RegisterType("Join", &types.Join{})
	DefaultEntityUnmarshaler.RegisterType("Leave", &types.Leave{})
	DefaultEntityUnmarshaler.RegisterType("Like", &types.Like{})
	DefaultEntityUnmarshaler.RegisterType("Listen", &types.Listen{})
	DefaultEntityUnmarshaler.RegisterType("Move", &types.Move{})
	DefaultEntityUnmarshaler.RegisterType("Offer", &types.Offer{})
	DefaultEntityUnmarshaler.RegisterType("Reject", &types.Reject{})
	DefaultEntityUnmarshaler.RegisterType("Read", &types.Read{})
	DefaultEntityUnmarshaler.RegisterType("Remove", &types.Remove{})
	DefaultEntityUnmarshaler.RegisterType("TentativeAccept", &types.TentativeAccept{})
	DefaultEntityUnmarshaler.RegisterType("TentativeReject", &types.TentativeReject{})
	DefaultEntityUnmarshaler.RegisterType("Travel", &types.Travel{})
	DefaultEntityUnmarshaler.RegisterType("Undo", &types.Undo{})
	DefaultEntityUnmarshaler.RegisterType("Update", &types.Update{})
	DefaultEntityUnmarshaler.RegisterType("View", &types.View{})
	DefaultEntityUnmarshaler.RegisterType("Application", &types.Application{})
	DefaultEntityUnmarshaler.RegisterType("Group", &types.Group{})
	DefaultEntityUnmarshaler.RegisterType("Organization", &types.Organization{})
	DefaultEntityUnmarshaler.RegisterType("Person", &types.Person{})
	DefaultEntityUnmarshaler.RegisterType("Service", &types.Service{})
	DefaultEntityUnmarshaler.RegisterType("Article", &types.Article{})
	DefaultEntityUnmarshaler.RegisterType("Audio", &types.Audio{})
	DefaultEntityUnmarshaler.RegisterType("Document", &types.Document{})
	DefaultEntityUnmarshaler.RegisterType("Event", &types.Event{})
	DefaultEntityUnmarshaler.RegisterType("Image", &types.Image{})
	DefaultEntityUnmarshaler.RegisterType("Note", &types.Note{})
	DefaultEntityUnmarshaler.RegisterType("Object", &types.Object{})
	DefaultEntityUnmarshaler.RegisterType("Page", &types.Page{})
	DefaultEntityUnmarshaler.RegisterType("Place", &types.Place{})
	DefaultEntityUnmarshaler.RegisterType("Profile", &types.Profile{})
	DefaultEntityUnmarshaler.RegisterType("Relationship", &types.Relationship{})
	DefaultEntityUnmarshaler.RegisterType("Tombstone", &types.Tombstone{})
	DefaultEntityUnmarshaler.RegisterType("Video", &types.Video{})
	DefaultEntityUnmarshaler.RegisterType("Collection", &types.Collection{})
	DefaultEntityUnmarshaler.RegisterType("CollectionPage", &types.CollectionPage{})
	DefaultEntityUnmarshaler.RegisterType("OrderedCollection", &types.Collection{})
	DefaultEntityUnmarshaler.RegisterType("OrderedCollectionPage", &types.CollectionPage{})
	DefaultEntityUnmarshaler.RegisterType("Link", &types.Link{})
	DefaultEntityUnmarshaler.RegisterType("Mention", &types.Mention{})
}
