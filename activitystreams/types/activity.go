package types

import (
	"github.com/brandonsides/pubblr/activitystreams"
)

type IntransitiveActivity struct {
	Object
	Actor      activitystreams.EntityIface `json:"actor,omitempty"`
	Target     activitystreams.EntityIface `json:"target,omitempty"`
	Result     activitystreams.EntityIface `json:"result,omitempty"`
	Origin     activitystreams.EntityIface `json:"origin,omitempty"`
	Instrument activitystreams.EntityIface `json:"instrument,omitempty"`
}

func (a *IntransitiveActivity) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(a)
}

func (a *IntransitiveActivity) Type() (string, error) {
	return "IntransitiveActivity", nil
}

type Activity struct {
	IntransitiveActivity
	Object activitystreams.EntityIface `json:"object,omitempty"`
}

func (a *Activity) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(a)
}

func (a *Activity) Type() (string, error) {
	return "Activity", nil
}

type Accept struct {
	Activity
}

func (a *Accept) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(a)
}

func (a *Accept) Type() (string, error) {
	return "Accept", nil
}

type TentativeAccept struct {
	Accept
}

func (a *TentativeAccept) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(a)
}

func (a *TentativeAccept) Type() (string, error) {
	return "TentativeAccept", nil
}

type Add struct {
	Activity
}

func (a *Add) Type() (string, error) {
	return "Add", nil
}

func (a *Add) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(a)
}

type Arrive struct {
	IntransitiveActivity
}

func (a *Arrive) Type() (string, error) {
	return "Arrive", nil
}

func (a *Arrive) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(a)
}

type Create struct {
	Activity
}

func (c *Create) Type() (string, error) {
	return "Create", nil
}

func (c *Create) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(c)
}

type Delete struct {
	Activity
}

func (d *Delete) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(d)
}

func (d *Delete) Type() (string, error) {
	return "Delete", nil
}

type Follow struct {
	Activity
}

func (f *Follow) Type() (string, error) {
	return "Follow", nil
}

func (f *Follow) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(f)
}

type Ignore struct {
	Activity
}

func (i *Ignore) Type() (string, error) {
	return "Ignore", nil
}

func (i *Ignore) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(i)
}

type Join struct {
	Activity
}

func (j *Join) Type() (string, error) {
	return "Join", nil
}

func (j *Join) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(j)
}

type Leave struct {
	Activity
}

func (l *Leave) Type() (string, error) {
	return "Leave", nil
}

func (l *Leave) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(l)
}

type Like struct {
	Activity
}

func (l *Like) Type() (string, error) {
	return "Like", nil
}

func (l *Like) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(l)
}

type Offer struct {
	Activity
}

func (o *Offer) Type() (string, error) {
	return "Offer", nil
}

func (o *Offer) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(o)
}

type Invite struct {
	Offer
}

func (i *Invite) Type() (string, error) {
	return "Invite", nil
}

func (i *Invite) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(i)
}

type Reject struct {
	Activity
}

func (r *Reject) Type() (string, error) {
	return "Reject", nil
}

func (r *Reject) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(r)
}

type TentativeReject struct {
	Reject
}

func (t *TentativeReject) Type() (string, error) {
	return "TentativeReject", nil
}

func (t *TentativeReject) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(t)
}

type Remove struct {
	Activity
}

func (r *Remove) Type() (string, error) {
	return "Remove", nil
}

func (r *Remove) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(r)
}

type Undo struct {
	Activity
}

func (u *Undo) Type() (string, error) {
	return "Undo", nil
}

func (u *Undo) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(u)
}

type Update struct {
	Activity
}

func (u *Update) Type() (string, error) {
	return "Update", nil
}

func (u *Update) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(u)
}

type View struct {
	Activity
}

func (v *View) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(v)
}

func (v *View) Type() (string, error) {
	return "View", nil
}

type Listen struct {
	IntransitiveActivity
}

func (l *Listen) Type() (string, error) {
	return "Listen", nil
}

func (l *Listen) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(l)
}

type Read struct {
	IntransitiveActivity
}

func (r *Read) Type() (string, error) {
	return "Read", nil
}

func (r *Read) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(r)
}

type Move struct {
	Activity
}

func (m *Move) Type() (string, error) {
	return "Move", nil
}

func (m *Move) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(m)
}

type Travel struct {
	IntransitiveActivity
}

func (t *Travel) Type() (string, error) {
	return "Travel", nil
}

func (t *Travel) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(t)
}

type Announce struct {
	Activity
}

func (a *Announce) Type() (string, error) {
	return "Announce", nil
}

func (a *Announce) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(a)
}

type Block struct {
	Ignore
}

func (b *Block) Type() (string, error) {
	return "Block", nil
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(b)
}

type Flag struct {
	Activity
}

func (f *Flag) Type() (string, error) {
	return "Flag", nil
}

func (f *Flag) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(f)
}

type Dislike struct {
	Activity
}

func (d *Dislike) Type() (string, error) {
	return "Dislike", nil
}

func (d *Dislike) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(d)
}

type Question struct {
	IntransitiveActivity
}

func (q *Question) Type() (string, error) {
	return "Question", nil
}

func (q *Question) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(q)
}

type SingleAnswerQuestion struct {
	Question
	OneOf []activitystreams.EntityIface `json:"oneOf,omitempty"`
}

func (q *SingleAnswerQuestion) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(q)
}

type MultiAnswerQuestion struct {
	Question
	AnyOf []activitystreams.EntityIface `json:"anyOf,omitempty"`
}

func (q *MultiAnswerQuestion) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(q)
}

type ClosedQuestion struct {
	Question
	Closed activitystreams.EntityIface `json:"closed,omitempty"`
}

func (q *ClosedQuestion) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(q)
}
