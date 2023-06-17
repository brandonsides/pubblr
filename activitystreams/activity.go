package activitystreams

import (
	"github.com/brandonsides/pubblr/util"
)

type IntransitiveActivity struct {
	Object
	Actor      *util.Either[ObjectIface, LinkIface] `json:"actor,omitempty"`
	Target     *util.Either[ObjectIface, LinkIface] `json:"target,omitempty"`
	Result     *util.Either[ObjectIface, LinkIface] `json:"result,omitempty"`
	Origin     *util.Either[ObjectIface, LinkIface] `json:"origin,omitempty"`
	Instrument *util.Either[ObjectIface, LinkIface] `json:"instrument,omitempty"`
}

func (a *IntransitiveActivity) MarshalJSON() ([]byte, error) {
	return MarshalObject(a)
}

func (a *IntransitiveActivity) Type() string {
	return "IntransitiveActivity"
}

type Activity struct {
	IntransitiveActivity
	Object *util.Either[ObjectIface, LinkIface] `json:"object,omitempty"`
}

func (a *Activity) MarshalJSON() ([]byte, error) {
	return MarshalObject(a)
}

func (a *Activity) Type() string {
	return "Activity"
}

type Accept struct {
	Activity
}

func (a *Accept) MarshalJSON() ([]byte, error) {
	return MarshalObject(a)
}

func (a *Accept) Type() string {
	return "Accept"
}

type TentativeAccept struct {
	Accept
}

func (a *TentativeAccept) MarshalJSON() ([]byte, error) {
	return MarshalObject(a)
}

func (a *TentativeAccept) Type() string {
	return "TentativeAccept"
}

type Add struct {
	Activity
}

func (a *Add) Type() string {
	return "Add"
}

func (a *Add) MarshalJSON() ([]byte, error) {
	return MarshalObject(a)
}

type Arrive struct {
	IntransitiveActivity
}

func (a *Arrive) Type() string {
	return "Arrive"
}

func (a *Arrive) MarshalJSON() ([]byte, error) {
	return MarshalObject(a)
}

type Create struct {
	Activity
}

func (c *Create) Type() string {
	return "Create"
}

func (c *Create) MarshalJSON() ([]byte, error) {
	return MarshalObject(c)
}

type Delete struct {
	Activity
}

func (d *Delete) MarshalJSON() ([]byte, error) {
	return MarshalObject(d)
}

func (d *Delete) Type() string {
	return "Delete"
}

type Follow struct {
	Activity
}

func (f *Follow) Type() string {
	return "Follow"
}

func (f *Follow) MarshalJSON() ([]byte, error) {
	return MarshalObject(f)
}

type Ignore struct {
	Activity
}

func (i *Ignore) Type() string {
	return "Ignore"
}

func (i *Ignore) MarshalJSON() ([]byte, error) {
	return MarshalObject(i)
}

type Join struct {
	Activity
}

func (j *Join) Type() string {
	return "Join"
}

func (j *Join) MarshalJSON() ([]byte, error) {
	return MarshalObject(j)
}

type Leave struct {
	Activity
}

func (l *Leave) Type() string {
	return "Leave"
}

func (l *Leave) MarshalJSON() ([]byte, error) {
	return MarshalObject(l)
}

type Like struct {
	Activity
}

func (l *Like) Type() string {
	return "Like"
}

func (l *Like) MarshalJSON() ([]byte, error) {
	return MarshalObject(l)
}

type Offer struct {
	Activity
}

func (o *Offer) Type() string {
	return "Offer"
}

func (o *Offer) MarshalJSON() ([]byte, error) {
	return MarshalObject(o)
}

type Invite struct {
	Offer
}

func (i *Invite) Type() string {
	return "Invite"
}

func (i *Invite) MarshalJSON() ([]byte, error) {
	return MarshalObject(i)
}

type Reject struct {
	Activity
}

func (r *Reject) Type() string {
	return "Reject"
}

func (r *Reject) MarshalJSON() ([]byte, error) {
	return MarshalObject(r)
}

type TentativeReject struct {
	Reject
}

func (t *TentativeReject) Type() string {
	return "TentativeReject"
}

func (t *TentativeReject) MarshalJSON() ([]byte, error) {
	return MarshalObject(t)
}

type Remove struct {
	Activity
}

func (r *Remove) Type() string {
	return "Remove"
}

func (r *Remove) MarshalJSON() ([]byte, error) {
	return MarshalObject(r)
}

type Undo struct {
	Activity
}

func (u *Undo) Type() string {
	return "Undo"
}

func (u *Undo) MarshalJSON() ([]byte, error) {
	return MarshalObject(u)
}

type Update struct {
	Activity
}

func (u *Update) Type() string {
	return "Update"
}

func (u *Update) MarshalJSON() ([]byte, error) {
	return MarshalObject(u)
}

type View struct {
	Activity
}

func (v *View) MarshalJSON() ([]byte, error) {
	return MarshalObject(v)
}

func (v *View) Type() string {
	return "View"
}

type Listen struct {
	IntransitiveActivity
}

func (l *Listen) Type() string {
	return "Listen"
}

func (l *Listen) MarshalJSON() ([]byte, error) {
	return MarshalObject(l)
}

type Read struct {
	IntransitiveActivity
}

func (r *Read) Type() string {
	return "Read"
}

func (r *Read) MarshalJSON() ([]byte, error) {
	return MarshalObject(r)
}

type Move struct {
	Activity
}

func (m *Move) Type() string {
	return "Move"
}

func (m *Move) MarshalJSON() ([]byte, error) {
	return MarshalObject(m)
}

type Travel struct {
	IntransitiveActivity
}

func (t *Travel) Type() string {
	return "Travel"
}

func (t *Travel) MarshalJSON() ([]byte, error) {
	return MarshalObject(t)
}

type Announce struct {
	Activity
}

func (a *Announce) Type() string {
	return "Announce"
}

func (a *Announce) MarshalJSON() ([]byte, error) {
	return MarshalObject(a)
}

type Block struct {
	Ignore
}

func (b *Block) Type() string {
	return "Block"
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return MarshalObject(b)
}

type Flag struct {
	Activity
}

func (f *Flag) Type() string {
	return "Flag"
}

func (f *Flag) MarshalJSON() ([]byte, error) {
	return MarshalObject(f)
}

type Dislike struct {
	Activity
}

func (d *Dislike) Type() string {
	return "Dislike"
}

func (d *Dislike) MarshalJSON() ([]byte, error) {
	return MarshalObject(d)
}

type Question struct {
	IntransitiveActivity
}

func (i *Question) Type() string {
	return "Question"
}

func (i *Question) MarshalJSON() ([]byte, error) {
	return MarshalObject(i)
}

type SingleAnswerQuestion struct {
	Question
	OneOf []util.Either[ObjectIface, LinkIface] `json:"oneOf,omitempty"`
}

type rawSingleAnswerQuestion SingleAnswerQuestion

func (q SingleAnswerQuestion) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawSingleAnswerQuestion)(&q))
}

type MultiAnswerQuestion struct {
	Question
	AnyOf []util.Either[ObjectIface, LinkIface] `json:"anyOf,omitempty"`
}

type rawMultiAnswerQuestion MultiAnswerQuestion

func (q MultiAnswerQuestion) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawMultiAnswerQuestion)(&q))
}

type ClosedQuestion struct {
	Question
	Closed *util.Either[ObjectIface, LinkIface] `json:"closed,omitempty"`
}

type rawClosedQuestion ClosedQuestion

func (q ClosedQuestion) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawClosedQuestion)(&q))
}
