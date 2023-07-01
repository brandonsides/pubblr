package activitystreams

import (
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
	return MarshalEntity(o)
}

// Represents an ActivityStreams Relationship object
type Relationship struct {
	Object
	Subject      *either.Either[ObjectIface, Link] `json:"subject,omitempty"`
	Obj          *either.Either[ObjectIface, Link] `json:"object,omitempty"`
	Relationship ObjectIface                       `json:"relationship,omitempty"`
}

func (r *Relationship) Type() (string, error) {
	return "Relationship", nil
}

func (r *Relationship) MarshalJSON() ([]byte, error) {
	return MarshalEntity(r)
}

// Represents an ActivityStreams Article object
type Article struct {
	Object
}

func (a *Article) Type() (string, error) {
	return "Article", nil
}

func (a *Article) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

// Represents an ActivityStreams Document object
type Document struct {
	Object
}

func (d *Document) Type() (string, error) {
	return "Document", nil
}

func (d *Document) MarshalJSON() ([]byte, error) {
	return MarshalEntity(d)
}

// Represents an ActivityStreams Audio object
type Audio struct {
	Object
}

func (a *Audio) Type() (string, error) {
	return "Audio", nil
}

func (a *Audio) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

// Represents an ActivityStreams Image object
type Image struct {
	Object
}

func (i *Image) Type() (string, error) {
	return "Image", nil
}

func (i *Image) MarshalJSON() ([]byte, error) {
	return MarshalEntity(i)
}

// Represents an ActivityStreams Video object
type Video struct {
	Object
}

func (v *Video) Type() (string, error) {
	return "Video", nil
}

func (v *Video) MarshalJSON() ([]byte, error) {
	return MarshalEntity(v)
}

// Represents an ActivityStreams Note object
type Note struct {
	Object
}

func (n *Note) Type() (string, error) {
	return "Note", nil
}

func (n *Note) MarshalJSON() ([]byte, error) {
	return MarshalEntity(n)
}

// Represents an ActivityStreams Page object
type Page struct {
	Object
}

func (p *Page) Type() (string, error) {
	return "Page", nil
}

func (p *Page) MarshalJSON() ([]byte, error) {
	return MarshalEntity(p)
}

// Represents an ActivityStreams Event object
type Event struct {
	Object
}

func (e *Event) Type() (string, error) {
	return "Event", nil
}

func (e *Event) MarshalJSON() ([]byte, error) {
	return MarshalEntity(e)
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
	return MarshalEntity(p)
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
	return MarshalEntity(p)
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
	return MarshalEntity(t)
}
