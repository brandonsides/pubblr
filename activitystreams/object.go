package activitystreams

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/brandonsides/pubblr/util/either"
)

// ObjectIface is an interface representing any ActivityStreams object.
// It is used to allow for polymorphism for types that embed Object.
// All types must embed Object implement this interface.
type ObjectIface interface {
	EntityIface
	// unexported method implemented only by Object
	// Forces all types to embed Object in order to implement this interface
	object() *Object
}

// Get a reference to the Object struct embedded in the given ObjectIface
func ToObject[O ObjectIface](o O) *Object {
	return o.object()
}

// Concrete type representing an ActivityStreams Object
type Object struct {
	Entity
	Attachment []EntityIface                     `json:"attachment,omitempty"`
	Audience   []EntityIface                     `json:"audience,omitempty"`
	Bcc        []EntityIface                     `json:"bcc,omitempty"`
	Bto        []EntityIface                     `json:"bto,omitempty"`
	Cc         []EntityIface                     `json:"cc,omitempty"`
	Context    EntityIface                       `json:"context,omitempty"`
	Generator  EntityIface                       `json:"generator,omitempty"`
	Icon       EntityIface                       `json:"icon,omitempty"`
	Image      EntityIface                       `json:"image,omitempty"`
	InReplyTo  []EntityIface                     `json:"inReplyTo,omitempty"`
	Location   []EntityIface                     `json:"location,omitempty"`
	Preview    EntityIface                       `json:"preview,omitempty"`
	Replies    CollectionIface                   `json:"replies,omitempty"`
	Tag        []EntityIface                     `json:"tag,omitempty"`
	To         []EntityIface                     `json:"to,omitempty"`
	URL        *either.Either[string, LinkIface] `json:"url,omitempty"`
	Content    string                            `json:"content,omitempty"`
	Duration   *time.Duration                    `json:"duration,omitempty"`
	EndTime    *time.Time                        `json:"endTime,omitempty"`
	Published  *time.Time                        `json:"published,omitempty"`
	StartTime  *time.Time                        `json:"startTime,omitempty"`
	Summary    string                            `json:"summary,omitempty"`
	Updated    *time.Time                        `json:"updated,omitempty"`
}

func (o *Object) object() *Object {
	return o
}

func (o *Object) Type() (string, error) {
	return "Object", nil
}

func (o *Object) MarshalJSON() ([]byte, error) {
	object, err := json.Marshal(o.Entity)
	if err != nil {
		return nil, err
	}

	var objectMap map[string]json.RawMessage
	err = json.Unmarshal(object, &objectMap)
	if err != nil {
		return nil, err
	}

	if o.Attachment != nil {
		attachmentIds := make([]string, len(o.Attachment))
		for i, attachment := range o.Attachment {
			attachmentIds[i] = ToEntity(attachment).Id
		}
		attachment, err := json.Marshal(attachmentIds)
		if err != nil {
			return nil, err
		}
		objectMap["attachment"] = attachment
	}

	if o.Audience != nil {
		audienceIds := make([]string, len(o.Audience))
		for i, audience := range o.Audience {
			audienceIds[i] = ToEntity(audience).Id
		}
		audience, err := json.Marshal(audienceIds)
		if err != nil {
			return nil, err
		}
		objectMap["audience"] = audience
	}

	if o.Bcc != nil {
		bccIds := make([]string, len(o.Bcc))
		for i, bcc := range o.Bcc {
			bccIds[i] = ToEntity(bcc).Id
		}
		bcc, err := json.Marshal(bccIds)
		if err != nil {
			return nil, err
		}
		objectMap["bcc"] = bcc
	}

	if o.Bto != nil {
		btoIds := make([]string, len(o.Bto))
		for i, bto := range o.Bto {
			btoIds[i] = ToEntity(bto).Id
		}
		bto, err := json.Marshal(btoIds)
		if err != nil {
			return nil, err
		}
		objectMap["bto"] = bto
	}

	if o.Cc != nil {
		ccIds := make([]string, len(o.Cc))
		for i, cc := range o.Cc {
			ccIds[i] = ToEntity(cc).Id
		}
		cc, err := json.Marshal(ccIds)
		if err != nil {
			return nil, err
		}
		objectMap["cc"] = cc
	}

	if o.Context != nil {
		objectMap["context"] = []byte(fmt.Sprintf("%q", ToEntity(o.Context).Id))
	}

	if o.Generator != nil {
		objectMap["generator"] = []byte(fmt.Sprintf("%q", ToEntity(o.Generator).Id))
	}

	if o.Icon != nil {
		objectMap["icon"] = []byte(fmt.Sprintf("%q", ToEntity(o.Icon).Id))
	}

	if o.Image != nil {
		objectMap["image"] = []byte(fmt.Sprintf("%q", ToEntity(o.Image).Id))
	}

	if o.InReplyTo != nil {
		inReplyToIds := make([]string, len(o.InReplyTo))
		for i, inReplyTo := range o.InReplyTo {
			inReplyToIds[i] = ToEntity(inReplyTo).Id
		}
		inReplyTo, err := json.Marshal(inReplyToIds)
		if err != nil {
			return nil, err
		}
		objectMap["inReplyTo"] = inReplyTo
	}

	if o.Location != nil {
		locationIds := make([]string, len(o.Location))
		for i, location := range o.Location {
			locationIds[i] = ToEntity(location).Id
		}
		location, err := json.Marshal(locationIds)
		if err != nil {
			return nil, err
		}
		objectMap["location"] = location
	}

	if o.Preview != nil {
		objectMap["preview"] = []byte(fmt.Sprintf("%q", ToEntity(o.Preview).Id))
	}

	if o.Replies != nil {
		objectMap["replies"] = []byte(fmt.Sprintf("%q", ToEntity(o.Replies).Id))
	}

	if o.Tag != nil {
		tagIds := make([]string, len(o.Tag))
		for i, tag := range o.Tag {
			tagIds[i] = ToEntity(tag).Id
		}
		tag, err := json.Marshal(tagIds)
		if err != nil {
			return nil, err
		}
		objectMap["tag"] = tag
	}

	if o.To != nil {
		toIds := make([]string, len(o.To))
		for i, to := range o.To {
			toIds[i] = ToEntity(to).Id
		}
		to, err := json.Marshal(toIds)
		if err != nil {
			return nil, err
		}
		objectMap["to"] = to
	}

	if o.URL != nil {
		var url string
		if o.URL.IsLeft() {
			url = *o.URL.Left()
		} else {
			url = ToEntity(*o.URL.Right()).Id
		}
		objectMap["url"] = []byte(fmt.Sprintf("%q", url))
	}

	if o.Content != "" {
		objectMap["content"] = []byte(fmt.Sprintf("%q", o.Content))
	}

	if o.Duration != nil {
		objectMap["duration"] = []byte(fmt.Sprintf("%q", o.Duration.String()))
	}

	if o.EndTime != nil {
		objectMap["endTime"] = []byte(fmt.Sprintf("%q", o.EndTime.Format(time.RFC3339)))
	}

	if o.Published != nil {
		objectMap["published"] = []byte(fmt.Sprintf("%q", o.Published.Format(time.RFC3339)))
	}

	if o.StartTime != nil {
		objectMap["startTime"] = []byte(fmt.Sprintf("%q", o.StartTime.Format(time.RFC3339)))
	}

	if o.Summary != "" {
		objectMap["summary"] = []byte(fmt.Sprintf("%q", o.Summary))
	}

	if o.Updated != nil {
		objectMap["updated"] = []byte(fmt.Sprintf("%q", o.Updated.Format(time.RFC3339)))
	}

	objectMap["type"] = []byte(fmt.Sprintf("%q", "Object"))

	return json.Marshal(objectMap)
}

// Represents an ActivityStreams Relationship object
type Relationship struct {
	Object
	Subject      *either.Either[ObjectIface, LinkIface] `json:"subject,omitempty"`
	Obj          *either.Either[ObjectIface, LinkIface] `json:"object,omitempty"`
	Relationship ObjectIface                            `json:"relationship,omitempty"`
}

func (r *Relationship) Type() (string, error) {
	return "Relationship", nil
}

func (r *Relationship) MarshalJSON() ([]byte, error) {
	relationship, err := json.Marshal(r.Object)
	if err != nil {
		return nil, err
	}

	var relationshipMap map[string]json.RawMessage
	err = json.Unmarshal(relationship, &relationshipMap)
	if err != nil {
		return nil, err
	}

	if r.Subject != nil {
		var subject EntityIface
		if r.Subject.IsLeft() {
			subject = *r.Subject.Left()
		} else {
			subject = *r.Subject.Right()
		}
		relationshipMap["subject"] = []byte(fmt.Sprintf("%q", ToEntity(subject).Id))
	}

	if r.Obj != nil {
		var obj EntityIface
		if r.Obj.IsLeft() {
			obj = *r.Obj.Left()
		} else {
			obj = *r.Obj.Right()
		}
		relationshipMap["object"] = []byte(fmt.Sprintf("%q", ToEntity(obj).Id))
	}

	if r.Relationship != nil {
		relationshipMap["relationship"] = []byte(fmt.Sprintf("%q", ToEntity(r.Relationship).Id))
	}

	relationshipMap["type"] = []byte(fmt.Sprintf("%q", "Relationship"))

	return json.Marshal(relationshipMap)
}

// Represents an ActivityStreams Article object
type Article struct {
	Object
}

func (a *Article) Type() (string, error) {
	return "Article", nil
}

func (a *Article) MarshalJSON() ([]byte, error) {
	article, err := json.Marshal(a.Object)
	if err != nil {
		return nil, err
	}

	var articleMap map[string]json.RawMessage
	err = json.Unmarshal(article, &articleMap)
	if err != nil {
		return nil, err
	}

	articleMap["type"] = []byte(fmt.Sprintf("%q", "Article"))

	return json.Marshal(articleMap)
}

// Represents an ActivityStreams Document object
type Document struct {
	Object
}

func (d *Document) Type() (string, error) {
	return "Document", nil
}

func (d *Document) MarshalJSON() ([]byte, error) {
	document, err := json.Marshal(d.Object)
	if err != nil {
		return nil, err
	}

	var documentMap map[string]json.RawMessage
	err = json.Unmarshal(document, &documentMap)
	if err != nil {
		return nil, err
	}

	documentMap["type"] = []byte(fmt.Sprintf("%q", "Document"))

	return json.Marshal(documentMap)
}

// Represents an ActivityStreams Audio object
type Audio struct {
	Object
}

func (a *Audio) Type() (string, error) {
	return "Audio", nil
}

func (a *Audio) MarshalJSON() ([]byte, error) {
	audio, err := json.Marshal(a.Object)
	if err != nil {
		return nil, err
	}

	var audioMap map[string]json.RawMessage
	err = json.Unmarshal(audio, &audioMap)
	if err != nil {
		return nil, err
	}

	audioMap["type"] = []byte(fmt.Sprintf("%q", "Audio"))

	return json.Marshal(audioMap)
}

// Represents an ActivityStreams Image object
type Image struct {
	Object
}

func (i *Image) Type() (string, error) {
	return "Image", nil
}

func (i *Image) MarshalJSON() ([]byte, error) {
	image, err := json.Marshal(i.Object)
	if err != nil {
		return nil, err
	}

	var imageMap map[string]json.RawMessage
	err = json.Unmarshal(image, &imageMap)
	if err != nil {
		return nil, err
	}

	imageMap["type"] = []byte(fmt.Sprintf("%q", "Image"))

	return json.Marshal(imageMap)
}

// Represents an ActivityStreams Video object
type Video struct {
	Object
}

func (v *Video) Type() (string, error) {
	return "Video", nil
}

func (v *Video) MarshalJSON() ([]byte, error) {
	video, err := json.Marshal(v.Object)
	if err != nil {
		return nil, err
	}

	var videoMap map[string]json.RawMessage
	err = json.Unmarshal(video, &videoMap)
	if err != nil {
		return nil, err
	}

	videoMap["type"] = []byte(fmt.Sprintf("%q", "Video"))

	return json.Marshal(videoMap)
}

// Represents an ActivityStreams Note object
type Note struct {
	Object
}

func (n *Note) Type() (string, error) {
	return "Note", nil
}

func (n *Note) MarshalJSON() ([]byte, error) {
	note, err := json.Marshal(n.Object)
	if err != nil {
		return nil, err
	}

	var noteMap map[string]json.RawMessage
	err = json.Unmarshal(note, &noteMap)
	if err != nil {
		return nil, err
	}

	noteMap["type"] = []byte(fmt.Sprintf("%q", "Note"))

	return json.Marshal(noteMap)
}

// Represents an ActivityStreams Page object
type Page struct {
	Object
}

func (p *Page) Type() (string, error) {
	return "Page", nil
}

func (p *Page) MarshalJSON() ([]byte, error) {
	page, err := json.Marshal(p.Object)
	if err != nil {
		return nil, err
	}

	var pageMap map[string]json.RawMessage
	err = json.Unmarshal(page, &pageMap)
	if err != nil {
		return nil, err
	}

	pageMap["type"] = []byte(fmt.Sprintf("%q", "Page"))

	return json.Marshal(pageMap)
}

// Represents an ActivityStreams Event object
type Event struct {
	Object
}

func (e *Event) Type() (string, error) {
	return "Event", nil
}

func (e *Event) MarshalJSON() ([]byte, error) {
	event, err := json.Marshal(e.Object)
	if err != nil {
		return nil, err
	}

	var eventMap map[string]json.RawMessage
	err = json.Unmarshal(event, &eventMap)
	if err != nil {
		return nil, err
	}

	eventMap["type"] = []byte(fmt.Sprintf("%q", "Event"))

	return json.Marshal(eventMap)
}

// Represents an ActivityStreams Place object
type Place struct {
	Object
	Accuracy  float64 `json:"accuracy,omitempty"`
	Altitude  float64 `json:"altitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Radius    float64 `json:"radius,omitempty"`
	Units     string  `json:"units,omitempty"`
}

func (p *Place) Type() (string, error) {
	return "Place", nil
}

func (p *Place) MarshalJSON() ([]byte, error) {
	place, err := json.Marshal(p.Object)
	if err != nil {
		return nil, err
	}

	var placeMap map[string]json.RawMessage
	err = json.Unmarshal(place, &placeMap)
	if err != nil {
		return nil, err
	}

	if p.Accuracy != 0 {
		placeMap["accuracy"] = []byte(fmt.Sprintf("%f", p.Accuracy))
	}

	if p.Altitude != 0 {
		placeMap["altitude"] = []byte(fmt.Sprintf("%f", p.Altitude))
	}

	if p.Latitude != 0 {
		placeMap["latitude"] = []byte(fmt.Sprintf("%f", p.Latitude))
	}

	if p.Longitude != 0 {
		placeMap["longitude"] = []byte(fmt.Sprintf("%f", p.Longitude))
	}

	if p.Radius != 0 {
		placeMap["radius"] = []byte(fmt.Sprintf("%f", p.Radius))
	}

	if p.Units != "" {
		placeMap["units"] = []byte(fmt.Sprintf("%q", p.Units))
	}

	placeMap["type"] = []byte(fmt.Sprintf("%q", "Place"))

	return json.Marshal(placeMap)
}

// Represents an ActivityStreams Profile object
type Profile struct {
	Object
	Describes ObjectIface `json:"describes,omitempty"`
}

func (p *Profile) Type() (string, error) {
	return "Profile", nil
}

func (p *Profile) MarshalJSON() ([]byte, error) {
	profile, err := json.Marshal(p.Object)
	if err != nil {
		return nil, err
	}

	var profileMap map[string]json.RawMessage
	err = json.Unmarshal(profile, &profileMap)
	if err != nil {
		return nil, err
	}

	if p.Describes != nil {
		profileMap["describes"] = []byte(fmt.Sprintf("%q", ToEntity(p.Describes).Id))
	}

	profileMap["type"] = []byte(fmt.Sprintf("%q", "Profile"))

	return json.Marshal(profileMap)
}

// Represents an ActivityStreams Tombstone object
type Tombstone struct {
	Object
	FormerType ObjectIface `json:"formerType,omitempty"`
	Deleted    *time.Time  `json:"deleted,omitempty"`
}

func (t *Tombstone) Type() (string, error) {
	return "Tombstone", nil
}

func (t *Tombstone) MarshalJSON() ([]byte, error) {
	tombstone, err := json.Marshal(t.Object)
	if err != nil {
		return nil, err
	}

	var tombstoneMap map[string]json.RawMessage
	err = json.Unmarshal(tombstone, &tombstoneMap)
	if err != nil {
		return nil, err
	}

	if t.FormerType != nil {
		tombstoneMap["formerType"] = []byte(fmt.Sprintf("%q", ToEntity(t.FormerType).Id))
	}

	if t.Deleted != nil {
		tombstoneMap["deleted"] = []byte(fmt.Sprintf("%q", t.Deleted.Format(time.RFC3339)))
	}

	tombstoneMap["type"] = []byte(fmt.Sprintf("%q", "Tombstone"))

	return json.Marshal(tombstoneMap)
}
