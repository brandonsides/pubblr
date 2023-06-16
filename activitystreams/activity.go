package activitystreams

import (
	"encoding/json"
	"reflect"

	"github.com/brandonsides/pubblr/util"
)

type IntransitiveActivity struct {
	Object     `json:"-"`
	Actor      *util.Either[ObjectIface, LinkIface] `json:"actor"`
	Target     *util.Either[ObjectIface, LinkIface] `json:"target"`
	Result     *util.Either[ObjectIface, LinkIface] `json:"result"`
	Origin     *util.Either[ObjectIface, LinkIface] `json:"origin"`
	Instrument *util.Either[ObjectIface, LinkIface] `json:"instrument"`
}

func (a *IntransitiveActivity) MarshalJSON() ([]byte, error) {
	objJson, err := ToObject(a).MarshalJSON()
	if err != nil {
		return nil, err
	}

	var objMap map[string]interface{}
	err = json.Unmarshal(objJson, &objMap)
	if err != nil {
		return nil, err
	}

	IntransitiveActivityReflectType := reflect.TypeOf((*IntransitiveActivity)(nil)).Elem()
	for fieldIndex := 0; fieldIndex < IntransitiveActivityReflectType.NumField(); fieldIndex++ {
		field := IntransitiveActivityReflectType.Field(fieldIndex)
		tag := field.Tag.Get("json")
		if tag == "" {
			continue
		}

		objMap[tag] = reflect.ValueOf(a).Elem().Field(fieldIndex).Interface()
	}

	objMap["type"] = a.Type()

	return json.Marshal(objMap)
}

func (a *IntransitiveActivity) Type() string {
	return "IntransitiveActivity"
}

func (a *IntransitiveActivity) intransitiveActivity() *IntransitiveActivity {
	return a
}

type Activity struct {
	IntransitiveActivity
	Object *util.Either[ObjectIface, LinkIface] `json:"object,omitempty"`
}

type rawActivity Activity

func (a Activity) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawActivity)(&a))
}

func (a *rawActivity) Type() string {
	return "Activity"
}

func (a *Activity) Type() string {
	return "Activity"
}

type Accept Activity

func (a *Accept) Type() string {
	return "Accept"
}

type TentativeAccept Accept

func (a *TentativeAccept) Type() string {
	return "TentativeAccept"
}

type Add Activity

func (a *Add) Type() string {
	return "Add"
}

type Arrive IntransitiveActivity

func (a *Arrive) Type() string {
	return "Arrive"
}

type Create Activity

func (c *Create) Type() string {
	return "Create"
}

type Delete Activity

func (d *Delete) Type() string {
	return "Delete"
}

type Follow Activity

func (f *Follow) Type() string {
	return "Follow"
}

type Ignore Activity

func (i *Ignore) Type() string {
	return "Ignore"
}

type Join Activity

func (j *Join) Type() string {
	return "Join"
}

type Leave Activity

func (l *Leave) Type() string {
	return "Leave"
}

type Like Activity

func (l *Like) Type() string {
	return "Like"
}

type Offer Activity

func (o *Offer) Type() string {
	return "Offer"
}

type Invite Offer

func (i *Invite) Type() string {
	return "Invite"
}

type Reject struct {
	Activity
}

func (r *Reject) Type() string {
	return "Reject"
}

type TentativeReject Reject

func (t *TentativeReject) Type() string {
	return "TentativeReject"
}

type Remove struct {
	Activity
}

func (r *Remove) Type() string {
	return "Remove"
}

type Undo struct {
	Activity
}

func (u *Undo) Type() string {
	return "Undo"
}

type Update struct {
	Activity
}

func (u *Update) Type() string {
	return "Update"
}

type View struct {
	Activity
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

type Read struct {
	IntransitiveActivity
}

func (r *Read) Type() string {
	return "Read"
}

type Move struct {
	Activity
}

func (m *Move) Type() string {
	return "Move"
}

type Travel struct {
	IntransitiveActivity
}

func (t *Travel) Type() string {
	return "Travel"
}

type Announce struct {
	Activity
}

func (a *Announce) Type() string {
	return "Announce"
}

type Block struct {
	Ignore
}

func (b *Block) Type() string {
	return "Block"
}

type Flag struct {
	Activity
}

func (f *Flag) Type() string {
	return "Flag"
}

type Dislike struct {
	Activity
}

func (d *Dislike) Type() string {
	return "Dislike"
}

type Question struct {
	IntransitiveActivity
}

func (i *Question) Type() string {
	return "Question"
}

type SingleAnswerQuestion struct {
	Question
	OneOf []util.Either[Object, Link] `json:"oneOf,omitempty"`
}

type rawSingleAnswerQuestion SingleAnswerQuestion

func (q SingleAnswerQuestion) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawSingleAnswerQuestion)(&q))
}

type MultiAnswerQuestion struct {
	Question
	AnyOf []util.Either[Object, Link] `json:"anyOf,omitempty"`
}

type rawMultiAnswerQuestion MultiAnswerQuestion

func (q MultiAnswerQuestion) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawMultiAnswerQuestion)(&q))
}

type ClosedQuestion struct {
	Question
	Closed *util.Either[Object, Link] `json:"closed,omitempty"`
}

type rawClosedQuestion ClosedQuestion

func (q ClosedQuestion) MarshalJSON() ([]byte, error) {
	return MarshalObject((*rawClosedQuestion)(&q))
}
