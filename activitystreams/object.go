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
	Attachment []EntityIface
	Audience   []EntityIface
	Bcc        []EntityIface
	Bto        []EntityIface
	Cc         []EntityIface
	Context    EntityIface
	Generator  EntityIface
	Icon       EntityIface
	Image      EntityIface
	InReplyTo  []EntityIface
	Location   []EntityIface
	Replies    CollectionIface
	Tag        []EntityIface
	To         []EntityIface
	URL        *either.Either[string, LinkIface]
	Content    string
	Duration   *time.Duration
	EndTime    *time.Time
	Published  *time.Time
	StartTime  *time.Time
	Summary    string
	Updated    *time.Time
}

func (o *Object) object() *Object {
	return o
}

func (o *Object) Type() (string, error) {
	return "Object", nil
}

func (o *Object) MarshalJSON() ([]byte, error) {
	object, err := json.Marshal(&o.Entity)
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

func (o *Object) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	err := o.Entity.UnmarshalEntity(u, b)
	if err != nil {
		return err
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(b, &objMap)
	if err != nil {
		return nil
	}

	if attachment, ok := objMap["attachment"]; ok {
		var rawAttachments []json.RawMessage
		err = json.Unmarshal(attachment, &rawAttachments)
		if err != nil {
			return err
		}

		attachments := make([]EntityIface, len(rawAttachments))
		for i, rawAttachment := range rawAttachments {
			attachment, err := u.UnmarshalEntity(rawAttachment)
			if err != nil {
				return err
			}
			attachments[i] = attachment
		}

		o.Attachment = attachments
	}

	if audience, ok := objMap["audience"]; ok {
		var rawAudiences []json.RawMessage
		err = json.Unmarshal(audience, &rawAudiences)
		if err != nil {
			return err
		}

		audiences := make([]EntityIface, len(rawAudiences))
		for i, rawAudience := range rawAudiences {
			audience, err := u.UnmarshalEntity(rawAudience)
			if err != nil {
				return err
			}
			audiences[i] = audience
		}

		o.Audience = audiences
	}

	if bcc, ok := objMap["bcc"]; ok {
		var rawBccs []json.RawMessage
		err = json.Unmarshal(bcc, &rawBccs)
		if err != nil {
			return err
		}

		bccs := make([]EntityIface, len(rawBccs))
		for i, rawBcc := range rawBccs {
			bcc, err := u.UnmarshalEntity(rawBcc)
			if err != nil {
				return err
			}
			bccs[i] = bcc
		}

		o.Bcc = bccs
	}

	if bto, ok := objMap["bto"]; ok {
		var rawBtos []json.RawMessage
		err = json.Unmarshal(bto, &rawBtos)
		if err != nil {
			return err
		}

		btos := make([]EntityIface, len(rawBtos))
		for i, rawBto := range rawBtos {
			bto, err := u.UnmarshalEntity(rawBto)
			if err != nil {
				return err
			}
			btos[i] = bto
		}

		o.Bto = btos
	}

	if cc, ok := objMap["cc"]; ok {
		var rawCcs []json.RawMessage
		err = json.Unmarshal(cc, &rawCcs)
		if err != nil {
			return err
		}

		ccs := make([]EntityIface, len(rawCcs))
		for i, rawCc := range rawCcs {
			cc, err := u.UnmarshalEntity(rawCc)
			if err != nil {
				return err
			}
			ccs[i] = cc
		}

		o.Cc = ccs
	}

	if context, ok := objMap["context"]; ok {
		o.Context, err = u.UnmarshalEntity(context)
		if err != nil {
			return err
		}
	}

	if generator, ok := objMap["generator"]; ok {
		o.Generator, err = u.UnmarshalEntity(generator)
		if err != nil {
			return err
		}
	}

	if icon, ok := objMap["icon"]; ok {
		o.Icon, err = u.UnmarshalEntity(icon)
		if err != nil {
			return err
		}
	}

	if image, ok := objMap["image"]; ok {
		o.Image, err = u.UnmarshalEntity(image)
		if err != nil {
			return err
		}
	}

	if inReplyTo, ok := objMap["inReplyTo"]; ok {
		var rawInReplyTos []json.RawMessage
		err = json.Unmarshal(inReplyTo, &rawInReplyTos)
		if err != nil {
			return err
		}

		inReplyTos := make([]EntityIface, len(rawInReplyTos))
		for i, rawInReplyTo := range rawInReplyTos {
			inReplyTo, err := u.UnmarshalEntity(rawInReplyTo)
			if err != nil {
				return err
			}
			inReplyTos[i] = inReplyTo
		}

		o.InReplyTo = inReplyTos
	}

	if location, ok := objMap["location"]; ok {
		var rawLocations []json.RawMessage
		err = json.Unmarshal(location, &rawLocations)
		if err != nil {
			return err
		}

		locations := make([]EntityIface, len(rawLocations))
		for i, rawLocation := range rawLocations {
			location, err := u.UnmarshalEntity(rawLocation)
			if err != nil {
				return err
			}
			locations[i] = location
		}

		o.Location = locations
	}

	if preview, ok := objMap["preview"]; ok {
		o.Preview, err = u.UnmarshalEntity(preview)
		if err != nil {
			return err
		}
	}

	if replies, ok := objMap["replies"]; ok {
		repliesEntity, err := u.UnmarshalEntity(replies)
		if err != nil {
			return err
		}

		o.Replies, ok = repliesEntity.(CollectionIface)
		if !ok {
			o.Replies = &Collection{Object: Object{Entity: *ToEntity(repliesEntity)}}
		}
	}

	if tag, ok := objMap["tag"]; ok {
		var rawTags []json.RawMessage
		err = json.Unmarshal(tag, &rawTags)
		if err != nil {
			return err
		}

		tags := make([]EntityIface, len(rawTags))
		for i, rawTag := range rawTags {
			tag, err := u.UnmarshalEntity(rawTag)
			if err != nil {
				return err
			}
			tags[i] = tag
		}

		o.Tag = tags
	}

	if to, ok := objMap["to"]; ok {
		var rawTos []json.RawMessage
		err = json.Unmarshal(to, &rawTos)
		if err != nil {
			return err
		}

		tos := make([]EntityIface, len(rawTos))
		for i, rawTo := range rawTos {
			to, err := u.UnmarshalEntity(rawTo)
			if err != nil {
				return err
			}
			tos[i] = to
		}

		o.To = tos
	}

	if url, ok := objMap["url"]; ok {
		urlEntity, err := u.UnmarshalEntity(url)
		if err != nil {
			return err
		}

		urlLinkIface, ok := urlEntity.(LinkIface)
		if !ok {
			urlLinkIface = &Link{Entity: *ToEntity(urlEntity)}
		}

		o.URL = either.Right[string](urlLinkIface)
	}

	if content, ok := objMap["content"]; ok {
		err = json.Unmarshal(content, &o.Content)
		if err != nil {
			return err
		}
	}

	if duration, ok := objMap["duration"]; ok {
		var rawDuration string
		err = json.Unmarshal(duration, &rawDuration)
		if err != nil {
			return err
		}

		parsedDuration, err := time.ParseDuration(rawDuration)
		if err != nil {
			return err
		}

		o.Duration = &parsedDuration
	}

	if endTime, ok := objMap["endTime"]; ok {
		var rawEndTime string
		err = json.Unmarshal(endTime, &rawEndTime)
		if err != nil {
			return err
		}

		parsedEndTime, err := time.Parse(time.RFC3339, rawEndTime)
		if err != nil {
			return err
		}

		o.EndTime = &parsedEndTime
	}

	if published, ok := objMap["published"]; ok {
		var rawPublished string
		err = json.Unmarshal(published, &rawPublished)
		if err != nil {
			return err
		}

		parsedPublished, err := time.Parse(time.RFC3339, rawPublished)
		if err != nil {
			return err
		}

		o.Published = &parsedPublished
	}

	if startTime, ok := objMap["startTime"]; ok {
		var rawStartTime string
		err = json.Unmarshal(startTime, &rawStartTime)
		if err != nil {
			return err
		}

		parsedStartTime, err := time.Parse(time.RFC3339, rawStartTime)
		if err != nil {
			return err
		}

		o.StartTime = &parsedStartTime
	}

	if summary, ok := objMap["summary"]; ok {
		err = json.Unmarshal(summary, &o.Summary)
		if err != nil {
			return err
		}
	}

	if updated, ok := objMap["updated"]; ok {
		var rawUpdated string
		err = json.Unmarshal(updated, &rawUpdated)
		if err != nil {
			return err
		}

		parsedUpdated, err := time.Parse(time.RFC3339, rawUpdated)
		if err != nil {
			return err
		}

		o.Updated = &parsedUpdated
	}

	return nil
}

// Represents an ActivityStreams Relationship object
type Relationship struct {
	Object
	Subject      EntityIface `json:"subject,omitempty"`
	Obj          EntityIface `json:"object,omitempty"`
	Relationship ObjectIface `json:"relationship,omitempty"`
}

func (r *Relationship) Type() (string, error) {
	return "Relationship", nil
}

func (r *Relationship) MarshalJSON() ([]byte, error) {
	relationship, err := json.Marshal(&r.Object)
	if err != nil {
		return nil, err
	}

	var relationshipMap map[string]json.RawMessage
	err = json.Unmarshal(relationship, &relationshipMap)
	if err != nil {
		return nil, err
	}

	if r.Subject != nil {
		relationshipMap["subject"] = []byte(fmt.Sprintf("%q", ToEntity(r.Subject).Id))
	}

	if r.Obj != nil {
		relationshipMap["object"] = []byte(fmt.Sprintf("%q", ToEntity(r.Obj).Id))
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
	article, err := json.Marshal(&a.Object)
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
	document, err := json.Marshal(&d.Object)
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
	audio, err := json.Marshal(&a.Object)
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
	image, err := json.Marshal(&i.Object)
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
	video, err := json.Marshal(&v.Object)
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
	note, err := json.Marshal(&n.Object)
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
	page, err := json.Marshal(&p.Object)
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
	event, err := json.Marshal(&e.Object)
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
	place, err := json.Marshal(&p.Object)
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
	profile, err := json.Marshal(&p.Object)
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
	tombstone, err := json.Marshal(&t.Object)
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
