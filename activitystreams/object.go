package activitystreams

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/brandonsides/pubblr/util"
)

// ObjectIface is an interface representing any ActivityStreams object.
// It is used to allow for polymorphism for types that embed Object.
// All types must embed Object implement this interface.
type ObjectIface interface {
	// unexported method implemented only by Object
	// Forces all types to embed Object in order to implement this interface
	object() *Object
	// Get the type of the object
	// This is used to set the "type" field in the JSON representation in lieu of
	// an object.Type field which may not correspond with the type of the object.
	// Object provides a default implementation which returns an empty string;
	// types that embed Object should override this method to return the correct
	// type.
	Type() string
}

// Get a reference to the Object struct embedded in the given ObjectIface
func ToObject(o ObjectIface) *Object {
	return o.object()
}

// Marhsal an ObjectIface to JSON
// Marshals the implementing type to JSON and adds a "type" field to the JSON
// representation with the value returned by the Type() method.
func MarshalObject(o ObjectIface) ([]byte, error) {
	objJson, err := json.Marshal((*rawObject)(ToObject(o)))
	if err != nil {
		return nil, err
	}

	var objMap map[string]interface{}
	err = json.Unmarshal(objJson, &objMap)
	if err != nil {
		return nil, err
	}

	IntransitiveActivityReflectType := reflect.TypeOf(o).Elem()
	for fieldIndex := 0; fieldIndex < IntransitiveActivityReflectType.NumField(); fieldIndex++ {
		field := IntransitiveActivityReflectType.Field(fieldIndex)
		tag := util.FromString(field.Tag.Get("json"))
		if tag.Name == "-" || tag.OmitEmpty && reflect.ValueOf(o).Elem().Field(fieldIndex).IsZero() {
			continue
		}

		v := reflect.ValueOf(o).Elem().Field(fieldIndex)
		if tag.String {
			objMap[tag.Name] = v.String()
		} else {
			objMap[tag.Name] = v.Interface()
		}
	}

	objMap["type"] = o.Type()

	return json.Marshal(objMap)
}

// Concrete type representing an ActivityStreams Object
type Object struct {
	Id           string                                `json:"id,omitempty"`
	Attachment   []util.Either[ObjectIface, LinkIface] `json:"attachment,omitempty"`
	AttributedTo []util.Either[ObjectIface, LinkIface] `json:"attributedTo,omitempty"`
	Audience     []util.Either[ObjectIface, LinkIface] `json:"audience,omitempty"`
	Bcc          []util.Either[ObjectIface, LinkIface] `json:"bcc,omitempty"`
	Bto          []util.Either[ObjectIface, LinkIface] `json:"bto,omitempty"`
	Cc           []util.Either[ObjectIface, LinkIface] `json:"cc,omitempty"`
	Context      *util.Either[ObjectIface, LinkIface]  `json:"context,omitempty"`
	Generator    *util.Either[ObjectIface, LinkIface]  `json:"generator,omitempty"`
	Icon         *util.Either[Image, LinkIface]        `json:"icon,omitempty"`
	Image        *util.Either[Image, LinkIface]        `json:"image,omitempty"`
	InReplyTo    []util.Either[ObjectIface, LinkIface] `json:"inReplyTo,omitempty"`
	Location     []util.Either[ObjectIface, LinkIface] `json:"location,omitempty"`
	Preview      *util.Either[ObjectIface, LinkIface]  `json:"preview,omitempty"`
	Replies      CollectionIface                       `json:"replies,omitempty"`
	Tag          []util.Either[ObjectIface, LinkIface] `json:"tag,omitempty"`
	To           []util.Either[ObjectIface, LinkIface] `json:"to,omitempty"`
	URL          *util.Either[string, LinkIface]       `json:"url,omitempty"`
	Content      string                                `json:"content,omitempty"`
	Name         string                                `json:"name,omitempty"`
	Duration     *time.Duration                        `json:"duration,omitempty"`
	MediaType    string                                `json:"mediaType,omitempty"`
	EndTime      *time.Time                            `json:"endTime,omitempty"`
	Published    *time.Time                            `json:"published,omitempty"`
	StartTime    *time.Time                            `json:"startTime,omitempty"`
	Summary      string                                `json:"summary,omitempty"`
	Updated      *time.Time                            `json:"updated,omitempty"`
}

type rawObject Object

func (o *Object) object() *Object {
	return o
}

func (o *Object) Type() string {
	return "Object"
}

func (o Object) MarshalJSON() ([]byte, error) {
	return MarshalObject(&o)
}

// Represents an object at the top level of an ActivityStreams document,
// including the @context field.
type TopLevelObject struct {
	ObjectIface
	Context string
}

func (t TopLevelObject) MarshalJSON() ([]byte, error) {
	j, err := MarshalObject(t.ObjectIface)
	if err != nil {
		return nil, err
	}

	jMap := make(map[string]interface{})
	err = json.Unmarshal(j, &jMap)
	if err != nil {
		return nil, err
	}

	if t.Context != "" {
		jMap["@context"] = t.Context
	}

	return json.Marshal(jMap)
}

// Represents an ActivityStreams Relationship object
type Relationship struct {
	Object
	Subject      *util.Either[ObjectIface, Link] `json:"subject,omitempty"`
	Obj          *util.Either[ObjectIface, Link] `json:"object,omitempty"`
	Relationship ObjectIface                     `json:"relationship,omitempty"`
}

type rawRelationship Relationship

func (r *Relationship) Type() string {
	return "Relationship"
}

func (r *rawRelationship) Type() string {
	return "Relationship"
}

func (r Relationship) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawRelationship)(&r))
}

// Represents an ActivityStreams Article object
type Article struct {
	Object
}

type rawArticle Article

func (a *Article) Type() string {
	return "Article"
}

func (p *rawArticle) Type() string {
	return "Article"
}

// Represents an ActivityStreams Document object
type Document struct {
	Object
}

type rawDocument Document

func (d *Document) Type() string {
	return "Document"
}

func (p *rawDocument) Type() string {
	return "Document"
}

func (d Document) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawDocument)(&d))
}

// Represents an ActivityStreams Audio object
type Audio struct {
	Object
}

type rawAudio Audio

func (a *Audio) Type() string {
	return "Audio"
}

func (p *rawAudio) Type() string {
	return "Audio"
}

func (a Audio) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawAudio)(&a))
}

// Represents an ActivityStreams Image object
type Image struct {
	Object
}

type rawImage Image

func (i *Image) Type() string {
	return "Image"
}

func (p *rawImage) Type() string {
	return "Image"
}

func (i Image) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawImage)(&i))
}

// Represents an ActivityStreams Video object
type Video struct {
	Object
}

type rawVideo Video

func (v *Video) Type() string {
	return "Video"
}

func (p *rawVideo) Type() string {
	return "Video"
}

func (v Video) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawVideo)(&v))
}

// Represents an ActivityStreams Note object
type Note struct {
	Object
}

type rawNote Note

func (n *Note) Type() string {
	return "Note"
}

func (p *rawNote) Type() string {
	return "Note"
}

func (n Note) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawNote)(&n))
}

// Represents an ActivityStreams Page object
type Page struct {
	Object
}

type rawPage Page

func (p *Page) Type() string {
	return "Page"
}

func (p *rawPage) Type() string {
	return "Page"
}

func (p Page) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawPage)(&p))
}

// Represents an ActivityStreams Event object
type Event struct {
	Object
}

type rawEvent Event

func (p *rawEvent) Type() string {
	return "Event"
}

func (e *Event) Type() string {
	return "Event"
}

func (e Event) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawEvent)(&e))
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

type rawPlace Place

func (p *rawPlace) Type() string {
	return "Place"
}

func (p *Place) Type() string {
	return "Place"
}

func (p Place) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawPlace)(&p))
}

// Represents an ActivityStreams Profile object
type Profile struct {
	Object
	Describes ObjectIface `json:"describes,omitempty"`
}

type rawProfile Profile

func (p *rawProfile) Type() string {
	return "Profile"
}

func (p *Profile) Type() string {
	return "Profile"
}

func (p Profile) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawProfile)(&p))
}

// Represents an ActivityStreams Tombstone object
type Tombstone struct {
	Object
	FormerType ObjectIface `json:"formerType,omitempty"`
	Deleted    *time.Time  `json:"deleted,omitempty"`
}

type rawTombstone Tombstone

func (p *rawTombstone) Type() string {
	return "Tombstone"
}

func (t *Tombstone) Type() string {
	return "Tombstone"
}

func (t Tombstone) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawTombstone)(&t))
}
