package activitystreams

import (
	"encoding/json"
	"fmt"
)

type ActivityIface interface {
	ObjectIface
	intransitiveActivity() *IntransitiveActivity
	Type() (string, error)
}

func ToIntransitiveActivity(a ActivityIface) *IntransitiveActivity {
	return a.intransitiveActivity()
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
	objJson, err := json.Marshal(&a.Object)
	if err != nil {
		return nil, err
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(objJson, &objMap)
	if err != nil {
		return nil, err
	}

	if a.Actor != nil {
		objMap["actor"] = []byte(fmt.Sprintf("%q", ToEntity(a.Actor).Id))
	}

	if a.Target != nil {
		objMap["target"] = []byte(fmt.Sprintf("%q", ToEntity(a.Target).Id))
	}

	if a.Result != nil {
		objMap["result"] = []byte(fmt.Sprintf("%q", ToEntity(a.Result).Id))
	}

	if a.Origin != nil {
		objMap["origin"] = []byte(fmt.Sprintf("%q", ToEntity(a.Origin).Id))
	}

	if a.Instrument != nil {
		objMap["instrument"] = []byte(fmt.Sprintf("%q", ToEntity(a.Instrument).Id))
	}

	objMap["type"] = json.RawMessage(fmt.Sprintf("%q", "IntransitiveActivity"))

	return json.Marshal(objMap)
}

func (a *IntransitiveActivity) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	err := a.Object.UnmarshalEntity(u, b)
	if err != nil {
		return err
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(b, &objMap)
	if err != nil {
		return nil
	}

	if actor, ok := objMap["actor"]; ok {
		a.Actor, err = u.UnmarshalEntity(actor)
		if err != nil {
			return err
		}
	}

	if target, ok := objMap["target"]; ok {
		a.Target, err = u.UnmarshalEntity(target)
		if err != nil {
			return err
		}
	}

	if result, ok := objMap["result"]; ok {
		a.Result, err = u.UnmarshalEntity(result)
		if err != nil {
			return err
		}
	}

	if origin, ok := objMap["origin"]; ok {
		a.Origin, err = u.UnmarshalEntity(origin)
		if err != nil {
			return err
		}
	}

	if instrument, ok := objMap["instrument"]; ok {
		a.Instrument, err = u.UnmarshalEntity(instrument)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *IntransitiveActivity) Type() (string, error) {
	return "IntransitiveActivity", nil
}

type TransitiveActivity struct {
	IntransitiveActivity
	Object EntityIface `json:"object,omitempty"`
}

func (a *TransitiveActivity) MarshalJSON() ([]byte, error) {
	intransitiveActivityJson, err := json.Marshal(&a.IntransitiveActivity)
	if err != nil {
		return nil, err
	}

	var transitiveActivityMap map[string]json.RawMessage
	err = json.Unmarshal(intransitiveActivityJson, &transitiveActivityMap)
	if err != nil {
		return nil, err
	}

	if a.Object != nil {
		transitiveActivityMap["object"] = []byte(fmt.Sprintf("%q", ToEntity(a.Object).Id))
	}

	transitiveActivityMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Activity"))

	return json.Marshal(transitiveActivityMap)
}

func (a *TransitiveActivity) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	err := a.IntransitiveActivity.UnmarshalEntity(u, b)
	if err != nil {
		return err
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(b, &objMap)
	if err != nil {
		return nil
	}

	if object, ok := objMap["object"]; ok {
		a.Object, err = u.UnmarshalEntity(object)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *TransitiveActivity) Type() (string, error) {
	return "Activity", nil
}

type Accept struct {
	TransitiveActivity
}

func (a *Accept) MarshalJSON() ([]byte, error) {
	accept, err := json.Marshal(&a.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var acceptMap map[string]json.RawMessage
	err = json.Unmarshal(accept, &acceptMap)
	if err != nil {
		return nil, err
	}

	acceptMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Accept"))

	return json.Marshal(acceptMap)
}

func (a *Accept) Type() (string, error) {
	return "Accept", nil
}

type TentativeAccept struct {
	Accept
}

func (a *TentativeAccept) MarshalJSON() ([]byte, error) {
	tAccept, err := json.Marshal(&a.Accept)
	if err != nil {
		return nil, err
	}

	var tAcceptMap map[string]json.RawMessage
	err = json.Unmarshal(tAccept, &tAcceptMap)
	if err != nil {
		return nil, err
	}

	tAcceptMap["type"] = json.RawMessage(fmt.Sprintf("%q", "TentativeAccept"))

	return json.Marshal(tAcceptMap)
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
	add, err := json.Marshal(&a.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var addMap map[string]json.RawMessage
	err = json.Unmarshal(add, &addMap)
	if err != nil {
		return nil, err
	}

	addMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Add"))

	return json.Marshal(addMap)
}

type Arrive struct {
	IntransitiveActivity
}

func (a *Arrive) Type() (string, error) {
	return "Arrive", nil
}

func (a *Arrive) MarshalJSON() ([]byte, error) {
	arrive, err := json.Marshal(&a.IntransitiveActivity)
	if err != nil {
		return nil, err
	}

	var arriveMap map[string]json.RawMessage
	err = json.Unmarshal(arrive, &arriveMap)
	if err != nil {
		return nil, err
	}

	arriveMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Arrive"))

	return json.Marshal(arriveMap)
}

type Create struct {
	TransitiveActivity
}

func (c *Create) Type() (string, error) {
	return "Create", nil
}

func (c *Create) MarshalJSON() ([]byte, error) {
	create, err := json.Marshal(&c.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var createMap map[string]json.RawMessage
	err = json.Unmarshal(create, &createMap)
	if err != nil {
		return nil, err
	}

	if c.Object != nil {
		createdObjectJSON, err := json.Marshal(&c.Object)
		if err != nil {
			return nil, err
		}
		createMap["object"] = createdObjectJSON
	}

	createMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Create"))

	return json.Marshal(createMap)
}

type Delete struct {
	TransitiveActivity
}

func (d *Delete) MarshalJSON() ([]byte, error) {
	delete, err := json.Marshal(&d.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var deleteMap map[string]json.RawMessage
	err = json.Unmarshal(delete, &deleteMap)
	if err != nil {
		return nil, err
	}

	deleteMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Delete"))

	return json.Marshal(deleteMap)
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
	follow, err := json.Marshal(&f.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var followMap map[string]json.RawMessage
	err = json.Unmarshal(follow, &followMap)
	if err != nil {
		return nil, err
	}

	followMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Follow"))

	return json.Marshal(followMap)
}

type Ignore struct {
	TransitiveActivity
}

func (i *Ignore) Type() (string, error) {
	return "Ignore", nil
}

func (i *Ignore) MarshalJSON() ([]byte, error) {
	ignore, err := json.Marshal(&i.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var ignoreMap map[string]json.RawMessage
	err = json.Unmarshal(ignore, &ignoreMap)
	if err != nil {
		return nil, err
	}

	ignoreMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Ignore"))

	return json.Marshal(ignoreMap)
}

type Join struct {
	TransitiveActivity
}

func (j *Join) Type() (string, error) {
	return "Join", nil
}

func (j *Join) MarshalJSON() ([]byte, error) {
	join, err := json.Marshal(&j.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var joinMap map[string]json.RawMessage
	err = json.Unmarshal(join, &joinMap)
	if err != nil {
		return nil, err
	}

	joinMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Join"))

	return json.Marshal(joinMap)
}

type Leave struct {
	TransitiveActivity
}

func (l *Leave) Type() (string, error) {
	return "Leave", nil
}

func (l *Leave) MarshalJSON() ([]byte, error) {
	leave, err := json.Marshal(&l.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var leaveMap map[string]json.RawMessage
	err = json.Unmarshal(leave, &leaveMap)
	if err != nil {
		return nil, err
	}

	leaveMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Leave"))

	return json.Marshal(leaveMap)
}

type Like struct {
	TransitiveActivity
}

func (l *Like) Type() (string, error) {
	return "Like", nil
}

func (l *Like) MarshalJSON() ([]byte, error) {
	like, err := json.Marshal(&l.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var likeMap map[string]json.RawMessage
	err = json.Unmarshal(like, &likeMap)
	if err != nil {
		return nil, err
	}

	likeMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Like"))

	return json.Marshal(likeMap)
}

type Offer struct {
	TransitiveActivity
}

func (o *Offer) Type() (string, error) {
	return "Offer", nil
}

func (o *Offer) MarshalJSON() ([]byte, error) {
	offer, err := json.Marshal(&o.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var offerMap map[string]json.RawMessage
	err = json.Unmarshal(offer, &offerMap)
	if err != nil {
		return nil, err
	}

	offerMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Offer"))

	return json.Marshal(offerMap)
}

type Invite struct {
	Offer
}

func (i *Invite) Type() (string, error) {
	return "Invite", nil
}

func (i *Invite) MarshalJSON() ([]byte, error) {
	invite, err := json.Marshal(&i.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var inviteMap map[string]json.RawMessage
	err = json.Unmarshal(invite, &inviteMap)
	if err != nil {
		return nil, err
	}

	inviteMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Invite"))

	return json.Marshal(inviteMap)
}

type Reject struct {
	TransitiveActivity
}

func (r *Reject) Type() (string, error) {
	return "Reject", nil
}

func (r *Reject) MarshalJSON() ([]byte, error) {
	reject, err := json.Marshal(&r.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var rejectMap map[string]json.RawMessage
	err = json.Unmarshal(reject, &rejectMap)
	if err != nil {
		return nil, err
	}

	rejectMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Reject"))

	return json.Marshal(rejectMap)
}

type TentativeReject struct {
	Reject
}

func (t *TentativeReject) Type() (string, error) {
	return "TentativeReject", nil
}

func (t *TentativeReject) MarshalJSON() ([]byte, error) {
	tentativeReject, err := json.Marshal(&t.Reject)
	if err != nil {
		return nil, err
	}

	var tentativeRejectMap map[string]json.RawMessage
	err = json.Unmarshal(tentativeReject, &tentativeRejectMap)
	if err != nil {
		return nil, err
	}

	tentativeRejectMap["type"] = json.RawMessage(fmt.Sprintf("%q", "TentativeReject"))

	return json.Marshal(tentativeRejectMap)
}

type Remove struct {
	TransitiveActivity
}

func (r *Remove) Type() (string, error) {
	return "Remove", nil
}

func (r *Remove) MarshalJSON() ([]byte, error) {
	remove, err := json.Marshal(&r.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var removeMap map[string]json.RawMessage
	err = json.Unmarshal(remove, &removeMap)
	if err != nil {
		return nil, err
	}

	removeMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Remove"))

	return json.Marshal(removeMap)
}

type Undo struct {
	TransitiveActivity
}

func (u *Undo) Type() (string, error) {
	return "Undo", nil
}

func (u *Undo) MarshalJSON() ([]byte, error) {
	undo, err := json.Marshal(&u.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var undoMap map[string]json.RawMessage
	err = json.Unmarshal(undo, &undoMap)
	if err != nil {
		return nil, err
	}

	undoMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Undo"))

	return json.Marshal(undoMap)
}

type Update struct {
	TransitiveActivity
}

func (u *Update) Type() (string, error) {
	return "Update", nil
}

func (u *Update) MarshalJSON() ([]byte, error) {
	update, err := json.Marshal(&u.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var updateMap map[string]json.RawMessage
	err = json.Unmarshal(update, &updateMap)
	if err != nil {
		return nil, err
	}

	updateMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Update"))

	return json.Marshal(updateMap)
}

type View struct {
	TransitiveActivity
}

func (v *View) Type() (string, error) {
	return "View", nil
}

func (v *View) MarshalJSON() ([]byte, error) {
	view, err := json.Marshal(&v.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var viewMap map[string]json.RawMessage
	err = json.Unmarshal(view, &viewMap)
	if err != nil {
		return nil, err
	}

	viewMap["type"] = json.RawMessage(fmt.Sprintf("%q", "View"))

	return json.Marshal(viewMap)
}

type Listen struct {
	IntransitiveActivity
}

func (l *Listen) Type() (string, error) {
	return "Listen", nil
}

func (l *Listen) MarshalJSON() ([]byte, error) {
	listen, err := json.Marshal(&l.IntransitiveActivity)
	if err != nil {
		return nil, err
	}

	var listenMap map[string]json.RawMessage
	err = json.Unmarshal(listen, &listenMap)
	if err != nil {
		return nil, err
	}

	listenMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Listen"))

	return json.Marshal(listenMap)
}

type Read struct {
	IntransitiveActivity
}

func (r *Read) Type() (string, error) {
	return "Read", nil
}

func (r *Read) MarshalJSON() ([]byte, error) {
	read, err := json.Marshal(&r.IntransitiveActivity)
	if err != nil {
		return nil, err
	}

	var readMap map[string]json.RawMessage
	err = json.Unmarshal(read, &readMap)
	if err != nil {
		return nil, err
	}

	readMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Read"))

	return json.Marshal(readMap)
}

type Move struct {
	TransitiveActivity
}

func (m *Move) Type() (string, error) {
	return "Move", nil
}

func (m *Move) MarshalJSON() ([]byte, error) {
	move, err := json.Marshal(&m.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var moveMap map[string]json.RawMessage
	err = json.Unmarshal(move, &moveMap)
	if err != nil {
		return nil, err
	}

	moveMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Move"))

	return json.Marshal(moveMap)
}

type Travel struct {
	IntransitiveActivity
}

func (t *Travel) Type() (string, error) {
	return "Travel", nil
}

func (t *Travel) MarshalJSON() ([]byte, error) {
	travel, err := json.Marshal(&t.IntransitiveActivity)
	if err != nil {
		return nil, err
	}

	var travelMap map[string]json.RawMessage
	err = json.Unmarshal(travel, &travelMap)
	if err != nil {
		return nil, err
	}

	travelMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Travel"))

	return json.Marshal(travelMap)
}

type Announce struct {
	TransitiveActivity
}

func (a *Announce) Type() (string, error) {
	return "Announce", nil
}

func (a *Announce) MarshalJSON() ([]byte, error) {
	announce, err := json.Marshal(&a.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var announceMap map[string]json.RawMessage
	err = json.Unmarshal(announce, &announceMap)
	if err != nil {
		return nil, err
	}

	announceMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Announce"))

	return json.Marshal(announceMap)
}

type Block struct {
	Ignore
}

func (b *Block) Type() (string, error) {
	return "Block", nil
}

func (b *Block) MarshalJSON() ([]byte, error) {
	block, err := json.Marshal(&b.Ignore)
	if err != nil {
		return nil, err
	}

	var blockMap map[string]json.RawMessage
	err = json.Unmarshal(block, &blockMap)
	if err != nil {
		return nil, err
	}

	blockMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Block"))

	return json.Marshal(blockMap)
}

type Flag struct {
	TransitiveActivity
}

func (f *Flag) Type() (string, error) {
	return "Flag", nil
}

func (f *Flag) MarshalJSON() ([]byte, error) {
	flag, err := json.Marshal(&f.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var flagMap map[string]json.RawMessage
	err = json.Unmarshal(flag, &flagMap)
	if err != nil {
		return nil, err
	}

	flagMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Flag"))

	return json.Marshal(flagMap)
}

type Dislike struct {
	TransitiveActivity
}

func (d *Dislike) Type() (string, error) {
	return "Dislike", nil
}

func (d *Dislike) MarshalJSON() ([]byte, error) {
	dislike, err := json.Marshal(&d.TransitiveActivity)
	if err != nil {
		return nil, err
	}

	var dislikeMap map[string]json.RawMessage
	err = json.Unmarshal(dislike, &dislikeMap)
	if err != nil {
		return nil, err
	}

	dislikeMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Dislike"))

	return json.Marshal(dislikeMap)
}

type Question struct {
	IntransitiveActivity
}

func (q *Question) Type() (string, error) {
	return "Question", nil
}

func (q *Question) MarshalJSON() ([]byte, error) {
	question, err := json.Marshal(&q.IntransitiveActivity)
	if err != nil {
		return nil, err
	}

	var questionMap map[string]json.RawMessage
	err = json.Unmarshal(question, &questionMap)
	if err != nil {
		return nil, err
	}

	questionMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Question"))

	return json.Marshal(questionMap)
}

type SingleAnswerQuestion struct {
	Question
	OneOf []EntityIface `json:"oneOf,omitempty"`
}

func (q *SingleAnswerQuestion) MarshalJSON() ([]byte, error) {
	singleAnswerQuestion, err := json.Marshal(&q.Question)
	if err != nil {
		return nil, err
	}

	var singleAnswerQuestionMap map[string]json.RawMessage
	err = json.Unmarshal(singleAnswerQuestion, &singleAnswerQuestionMap)
	if err != nil {
		return nil, err
	}

	singleAnswerQuestionMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Question"))

	if q.OneOf != nil {
		oneOf, err := json.Marshal(q.OneOf)
		if err != nil {
			return nil, err
		}

		singleAnswerQuestionMap["oneOf"] = json.RawMessage(oneOf)
	}

	return json.Marshal(singleAnswerQuestionMap)
}

func (q *SingleAnswerQuestion) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	err := q.Question.UnmarshalEntity(u, b)
	if err != nil {
		return err
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(b, &objMap)
	if err != nil {
		return nil
	}

	if oneOf, ok := objMap["oneOf"]; ok {
		err = json.Unmarshal(oneOf, &q.OneOf)
		if err != nil {
			return err
		}
	}

	return nil
}

type MultiAnswerQuestion struct {
	Question
	AnyOf []EntityIface `json:"anyOf,omitempty"`
}

func (q *MultiAnswerQuestion) MarshalJSON() ([]byte, error) {
	multiAnswerQuestion, err := json.Marshal(&q.Question)
	if err != nil {
		return nil, err
	}

	var multiAnswerQuestionMap map[string]json.RawMessage
	err = json.Unmarshal(multiAnswerQuestion, &multiAnswerQuestionMap)
	if err != nil {
		return nil, err
	}

	multiAnswerQuestionMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Question"))

	if q.AnyOf != nil {
		anyOf, err := json.Marshal(q.AnyOf)
		if err != nil {
			return nil, err
		}

		multiAnswerQuestionMap["anyOf"] = json.RawMessage(anyOf)
	}

	return json.Marshal(multiAnswerQuestionMap)
}

func (q *MultiAnswerQuestion) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	err := q.Question.UnmarshalEntity(u, b)
	if err != nil {
		return err
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(b, &objMap)
	if err != nil {
		return nil
	}

	if anyOf, ok := objMap["anyOf"]; ok {
		err = json.Unmarshal(anyOf, &q.AnyOf)
		if err != nil {
			return err
		}
	}

	return nil
}

type ClosedQuestion struct {
	Question
	Closed EntityIface `json:"closed,omitempty"`
}

func (q *ClosedQuestion) MarshalJSON() ([]byte, error) {
	closedQuestion, err := json.Marshal(&q.Question)
	if err != nil {
		return nil, err
	}

	var closedQuestionMap map[string]json.RawMessage
	err = json.Unmarshal(closedQuestion, &closedQuestionMap)
	if err != nil {
		return nil, err
	}

	closedQuestionMap["type"] = json.RawMessage(fmt.Sprintf("%q", "Question"))

	if q.Closed != nil {
		closed, err := json.Marshal(q.Closed)
		if err != nil {
			return nil, err
		}

		closedQuestionMap["closed"] = json.RawMessage(closed)
	}

	return json.Marshal(closedQuestionMap)
}

func (q *ClosedQuestion) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	err := q.Question.UnmarshalEntity(u, b)
	if err != nil {
		return err
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(b, &objMap)
	if err != nil {
		return nil
	}

	if closed, ok := objMap["closed"]; ok {
		err = json.Unmarshal(closed, &q.Closed)
		if err != nil {
			return err
		}
	}

	return nil
}
