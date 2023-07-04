package activitystreams

type ActivityIface interface {
	ObjectIface
	intransitiveActivity() *IntransitiveActivity
	Type() (string, error)
}

type IntransitiveActivity struct {
	Object
	Actor      EntityIface `json:"actor,omitempty"`
	Target     EntityIface `json:"target,omitempty"`
	Result     EntityIface `json:"result,omitempty"`
	Origin     EntityIface `json:"origin,omitempty"`
	Instrument EntityIface `json:"instrument,omitempty"`
}

func (a *IntransitiveActivity) intransitiveActivity() *IntransitiveActivity {
	return a
}

func (a *IntransitiveActivity) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

func (a *IntransitiveActivity) Type() (string, error) {
	return "IntransitiveActivity", nil
}

type TransitiveActivity struct {
	IntransitiveActivity
	Object EntityIface `json:"object,omitempty"`
}

func (a *TransitiveActivity) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

func (a *TransitiveActivity) Type() (string, error) {
	return "Activity", nil
}

type Accept struct {
	TransitiveActivity
}

func (a *Accept) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

func (a *Accept) Type() (string, error) {
	return "Accept", nil
}

type TentativeAccept struct {
	Accept
}

func (a *TentativeAccept) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

func (a *TentativeAccept) Type() (string, error) {
	return "TentativeAccept", nil
}

type Add struct {
	TransitiveActivity
}

func (a *Add) Type() (string, error) {
	return "Add", nil
}

func (a *Add) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

type Arrive struct {
	IntransitiveActivity
}

func (a *Arrive) Type() (string, error) {
	return "Arrive", nil
}

func (a *Arrive) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

type Create struct {
	TransitiveActivity
}

func (c *Create) Type() (string, error) {
	return "Create", nil
}

func (c *Create) MarshalJSON() ([]byte, error) {
	return MarshalEntity(c)
}

type Delete struct {
	TransitiveActivity
}

func (d *Delete) MarshalJSON() ([]byte, error) {
	return MarshalEntity(d)
}

func (d *Delete) Type() (string, error) {
	return "Delete", nil
}

type Follow struct {
	TransitiveActivity
}

func (f *Follow) Type() (string, error) {
	return "Follow", nil
}

func (f *Follow) MarshalJSON() ([]byte, error) {
	return MarshalEntity(f)
}

type Ignore struct {
	TransitiveActivity
}

func (i *Ignore) Type() (string, error) {
	return "Ignore", nil
}

func (i *Ignore) MarshalJSON() ([]byte, error) {
	return MarshalEntity(i)
}

type Join struct {
	TransitiveActivity
}

func (j *Join) Type() (string, error) {
	return "Join", nil
}

func (j *Join) MarshalJSON() ([]byte, error) {
	return MarshalEntity(j)
}

type Leave struct {
	TransitiveActivity
}

func (l *Leave) Type() (string, error) {
	return "Leave", nil
}

func (l *Leave) MarshalJSON() ([]byte, error) {
	return MarshalEntity(l)
}

type Like struct {
	TransitiveActivity
}

func (l *Like) Type() (string, error) {
	return "Like", nil
}

func (l *Like) MarshalJSON() ([]byte, error) {
	return MarshalEntity(l)
}

type Offer struct {
	TransitiveActivity
}

func (o *Offer) Type() (string, error) {
	return "Offer", nil
}

func (o *Offer) MarshalJSON() ([]byte, error) {
	return MarshalEntity(o)
}

type Invite struct {
	Offer
}

func (i *Invite) Type() (string, error) {
	return "Invite", nil
}

func (i *Invite) MarshalJSON() ([]byte, error) {
	return MarshalEntity(i)
}

type Reject struct {
	TransitiveActivity
}

func (r *Reject) Type() (string, error) {
	return "Reject", nil
}

func (r *Reject) MarshalJSON() ([]byte, error) {
	return MarshalEntity(r)
}

type TentativeReject struct {
	Reject
}

func (t *TentativeReject) Type() (string, error) {
	return "TentativeReject", nil
}

func (t *TentativeReject) MarshalJSON() ([]byte, error) {
	return MarshalEntity(t)
}

type Remove struct {
	TransitiveActivity
}

func (r *Remove) Type() (string, error) {
	return "Remove", nil
}

func (r *Remove) MarshalJSON() ([]byte, error) {
	return MarshalEntity(r)
}

type Undo struct {
	TransitiveActivity
}

func (u *Undo) Type() (string, error) {
	return "Undo", nil
}

func (u *Undo) MarshalJSON() ([]byte, error) {
	return MarshalEntity(u)
}

type Update struct {
	TransitiveActivity
}

func (u *Update) Type() (string, error) {
	return "Update", nil
}

func (u *Update) MarshalJSON() ([]byte, error) {
	return MarshalEntity(u)
}

type View struct {
	TransitiveActivity
}

func (v *View) MarshalJSON() ([]byte, error) {
	return MarshalEntity(v)
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
	return MarshalEntity(l)
}

type Read struct {
	IntransitiveActivity
}

func (r *Read) Type() (string, error) {
	return "Read", nil
}

func (r *Read) MarshalJSON() ([]byte, error) {
	return MarshalEntity(r)
}

type Move struct {
	TransitiveActivity
}

func (m *Move) Type() (string, error) {
	return "Move", nil
}

func (m *Move) MarshalJSON() ([]byte, error) {
	return MarshalEntity(m)
}

type Travel struct {
	IntransitiveActivity
}

func (t *Travel) Type() (string, error) {
	return "Travel", nil
}

func (t *Travel) MarshalJSON() ([]byte, error) {
	return MarshalEntity(t)
}

type Announce struct {
	TransitiveActivity
}

func (a *Announce) Type() (string, error) {
	return "Announce", nil
}

func (a *Announce) MarshalJSON() ([]byte, error) {
	return MarshalEntity(a)
}

type Block struct {
	Ignore
}

func (b *Block) Type() (string, error) {
	return "Block", nil
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return MarshalEntity(b)
}

type Flag struct {
	TransitiveActivity
}

func (f *Flag) Type() (string, error) {
	return "Flag", nil
}

func (f *Flag) MarshalJSON() ([]byte, error) {
	return MarshalEntity(f)
}

type Dislike struct {
	TransitiveActivity
}

func (d *Dislike) Type() (string, error) {
	return "Dislike", nil
}

func (d *Dislike) MarshalJSON() ([]byte, error) {
	return MarshalEntity(d)
}

type Question struct {
	IntransitiveActivity
}

func (q *Question) Type() (string, error) {
	return "Question", nil
}

func (q *Question) MarshalJSON() ([]byte, error) {
	return MarshalEntity(q)
}

type SingleAnswerQuestion struct {
	Question
	OneOf []EntityIface `json:"oneOf,omitempty"`
}

func (q *SingleAnswerQuestion) MarshalJSON() ([]byte, error) {
	return MarshalEntity(q)
}

type MultiAnswerQuestion struct {
	Question
	AnyOf []EntityIface `json:"anyOf,omitempty"`
}

func (q *MultiAnswerQuestion) MarshalJSON() ([]byte, error) {
	return MarshalEntity(q)
}

type ClosedQuestion struct {
	Question
	Closed EntityIface `json:"closed,omitempty"`
}

func (q *ClosedQuestion) MarshalJSON() ([]byte, error) {
	return MarshalEntity(q)
}
