package database

import (
	"database/sql"
	"errors"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/database/util"
)

func toDBEntity(entity activitystreams.EntityIface) (dbEntity, error) {
	var typ sql.NullString
	var err error
	typ.String, err = entity.Type()
	if err != nil {
		typ.Valid = false
	}

	e := activitystreams.ToEntity(entity)
	preview, err := toDBEntity(e.Preview)
	if err != nil {
		return dbEntity{}, err
	}
	attributedTo := make([]dbEntity, len(e.AttributedTo))
	for i, v := range e.AttributedTo {
		attributedTo[i], err = toDBEntity(v)
		if err != nil {
			return dbEntity{}, err
		}
	}

	ret := dbEntity{
		Type: typ,
		MediaType: sql.NullString{
			String: e.MediaType,
			Valid:  true,
		},
		Preview:      &preview,
		AttributedTo: attributedTo,
	}

	var rest interface{} = nil
	if typ.Valid {
		switch typ.String {
		case "Link", "Mention":
			linkIface, ok := entity.(activitystreams.LinkIface)
			if !ok {
				return dbEntity{}, errors.New("Could not convert entity to LinkIface")
			}
			dbLink := toDBLink(linkIface)
			dbLink.Entity = &ret
			rest = dbLink
		default:
			objectIface, ok := entity.(activitystreams.ObjectIface)
			if !ok {
				return dbEntity{}, errors.New("Could not convert entity to ObjectIface")
			}
			dbObj, err := toDBObject(objectIface)
			if err != nil {
				return dbEntity{}, err
			}
			dbObj.Entity = &ret
			rest = dbObj
		}
	}
	ret.Rest = rest

	return ret, nil
}

func toDBLink(link activitystreams.LinkIface) dbLink {
	l := activitystreams.ToLink(link)
	return dbLink{
		Href: sql.NullString{
			String: l.Href,
			Valid:  l.Href != "",
		},
		Hreflang: sql.NullString{
			String: l.HrefLang,
			Valid:  l.HrefLang != "",
		},
		MediaType: sql.NullString{
			String: l.MediaType,
			Valid:  l.MediaType != "",
		},
		Rel:    util.StringArray(l.Rel),
		Height: l.Height,
		Width:  l.Width,
	}
}

func toDBObject(object activitystreams.ObjectIface) (dbObject, error) {
	typ, err := object.Type()
	if err != nil {
		return dbObject{}, err
	}

	o := activitystreams.ToObject(object)

	attachments := make([]dbEntity, len(o.Attachment))
	for i, v := range o.Attachment {
		var err error
		attachments[i], err = toDBEntity(v)
		if err != nil {
			return dbObject{}, err
		}
	}

	audience := make([]dbEntity, len(o.Audience))
	for i, v := range o.Audience {
		var err error
		audience[i], err = toDBEntity(v)
		if err != nil {
			return dbObject{}, err
		}
	}

	bcc := make([]dbEntity, len(o.Bcc))
	for i, v := range o.Bcc {
		var err error
		bcc[i], err = toDBEntity(v)
		if err != nil {
			return dbObject{}, err
		}
	}

	bto := make([]dbEntity, len(o.Bto))
	for i, v := range o.Bto {
		var err error
		bto[i], err = toDBEntity(v)
		if err != nil {
			return dbObject{}, err
		}
	}

	cc := make([]dbEntity, len(o.Cc))
	for i, v := range o.Cc {
		var err error
		cc[i], err = toDBEntity(v)
		if err != nil {
			return dbObject{}, err
		}
	}

	inReplyTo := make([]dbEntity, len(o.InReplyTo))
	for i, v := range o.InReplyTo {
		var err error
		inReplyTo[i], err = toDBEntity(v)
		if err != nil {
			return dbObject{}, err
		}
	}

	tags := make([]dbEntity, len(o.Tag))
	for i, v := range o.Tag {
		var err error
		tags[i], err = toDBEntity(v)
		if err != nil {
			return dbObject{}, err
		}
	}

	to := make([]dbEntity, len(o.To))
	for i, v := range o.To {
		var err error
		to[i], err = toDBEntity(v)
		if err != nil {
			return dbObject{}, err
		}
	}

	locations := make([]dbEntity, len(o.Location))
	for i, v := range o.Location {
		var err error
		locations[i], err = toDBEntity(v)
		if err != nil {
			return dbObject{}, err
		}
	}

	context, err := toDBEntity(o.Context)
	if err != nil {
		return dbObject{}, err
	}

	generator, err := toDBEntity(o.Generator)
	if err != nil {
		return dbObject{}, err
	}

	icon, err := toDBEntity(o.Icon)
	if err != nil {
		return dbObject{}, err
	}

	image, err := toDBEntity(o.Image)
	if err != nil {
		return dbObject{}, err
	}

	var linkUrl *dbLink = nil
	var stringUrl sql.NullString
	if o.URL.IsLeft() {
		stringUrl = sql.NullString{
			String: *o.URL.Left(),
			Valid:  true,
		}
	} else {
		linkUrl = &dbLink{}
		*linkUrl = toDBLink(activitystreams.ToLink(*o.URL.Right()))
	}

	var startTime sql.NullTime
	if o.StartTime != nil {
		startTime = sql.NullTime{
			Time:  *o.StartTime,
			Valid: true,
		}
	}

	var endTime sql.NullTime
	if o.EndTime != nil {
		endTime = sql.NullTime{
			Time:  *o.EndTime,
			Valid: true,
		}
	}

	var duration sql.NullInt64
	if o.Duration != nil {
		duration = sql.NullInt64{
			Int64: int64(*o.Duration),
			Valid: true,
		}
	}

	ret := dbObject{
		Context:   &context,
		Generator: &generator,
		Icon:      &icon,
		Image:     &image,
		Location:  locations,
		LinkUrl:   linkUrl,
		StringUrl: stringUrl,
		Content: sql.NullString{
			String: o.Content,
			Valid:  o.Content != "",
		},
		Summary: sql.NullString{
			String: o.Summary,
			Valid:  o.Summary != "",
		},
		StartTime:  startTime,
		EndTime:    endTime,
		Duration:   duration,
		Attachment: attachments,
		Audience:   audience,
		Bcc:        bcc,
		Bto:        bto,
		Cc:         cc,
		InReplyTo:  inReplyTo,
		Tag:        tags,
		To:         to,
	}

	var rest interface{}
	switch typ {
	case "Object":
	case "Relationship":
		r, ok := object.(*activitystreams.Relationship)
		if !ok {
			return dbObject{}, errors.New("Could not convert object to Relationship")
		}
		dbRel, err := toDBRelationship(r)
		if err != nil {
			return dbObject{}, err
		}
		dbRel.Object = &ret
		rest = dbRel
	case "Place":
		p, ok := object.(*activitystreams.Place)
		if !ok {
			return dbObject{}, errors.New("Could not convert object to Place")
		}
		dbPlace := toDBPlace(p)
		dbPlace.Object = &ret
		rest = dbPlace
	case "Profile":
		p, ok := object.(*activitystreams.Profile)
		if !ok {
			return dbObject{}, errors.New("Could not convert object to Profile")
		}
		dbProf := toDBProfile(p)
		dbProf.Object = &ret
		rest = dbProf
	case "Person", "Service", "Group", "Organization", "Application":
		a, ok := object.(activitystreams.ActorIface)
		if !ok {
			return dbObject{}, errors.New("Could not convert object to ActorIface")
		}
		dbActor := toDBActor(a)
		dbActor.Object = &ret
		rest = dbActor
	case "Collection", "OrderedCollection":
		c, ok := object.(activitystreams.CollectionIface)
		if !ok {
			return dbObject{}, errors.New("Could not convert object to CollectionIface")
		}
		dbColl := toDBCollection(c)
		dbColl.Object = &ret
		rest = dbColl
	default:
		a, ok := object.(activitystreams.ActivityIface)
		if !ok {
			return dbObject{}, errors.New("Could not convert object to ActivityIface")
		}
		dbActv, err := toDBActivity(a)
		if err != nil {
			return dbObject{}, err
		}
		dbActv.Object = &ret
		rest = dbActv
	}
	ret.Rest = rest

	return ret, nil
}

func toDBRelationship(relationship *activitystreams.Relationship) (dbRelationship, error) {
	subj, err := toDBEntity(relationship.Subject)
	if err != nil {
		return dbRelationship{}, err
	}

	obj, err := toDBEntity(relationship.Obj)
	if err != nil {
		return dbRelationship{}, err
	}

	rel, err := toDBObject(relationship.Relationship)
	if err != nil {
		return dbRelationship{}, err
	}

	return dbRelationship{
		Subject:      &subj,
		Obj:          &obj,
		Relationship: &rel,
	}, nil
}

func toDBPlace(place *activitystreams.Place) dbPlace {
	return dbPlace{
		Accuracy: sql.NullFloat64{
			Float64: place.Accuracy,
			Valid:   place.Accuracy != 0,
		},
		Altitude: sql.NullFloat64{
			Float64: place.Altitude,
			Valid:   place.Altitude != 0,
		},
		Latitude: sql.NullFloat64{
			Float64: place.Latitude,
			Valid:   place.Latitude != 0,
		},
		Longitude: sql.NullFloat64{
			Float64: place.Longitude,
			Valid:   place.Longitude != 0,
		},
		Radius: sql.NullFloat64{
			Float64: place.Radius,
			Valid:   place.Radius != 0,
		},
		Units: sql.NullString{
			String: place.Units,
			Valid:  place.Units != "",
		},
	}
}

func toDBProfile(profile *activitystreams.Profile) dbProfile {
	describes, err := toDBEntity(profile.Describes)
	if err != nil {
		return dbProfile{}
	}

	return dbProfile{
		Describes: &describes,
	}
}

func toDBActivity(activity activitystreams.ActivityIface) (dbActivity, error) {
	typ, err := activity.Type()
	if err != nil {
		return dbActivity{}, err
	}

	a := activitystreams.ToIntransitiveActivity(activity)

	actor, err := toDBEntity(a.Actor)
	if err != nil {
		return dbActivity{}, err
	}

	objEnt, err := toDBEntity(&a.Object)
	if err != nil {
		return dbActivity{}, err
	}
	object, ok := objEnt.Rest.(dbObject)
	if !ok {
		return dbActivity{}, errors.New("Could not convert object to dbObject")
	}

	target, err := toDBEntity(a.Target)
	if err != nil {
		return dbActivity{}, err
	}

	result, err := toDBEntity(a.Result)
	if err != nil {
		return dbActivity{}, err
	}

	origin, err := toDBEntity(a.Origin)
	if err != nil {
		return dbActivity{}, err
	}

	instrument, err := toDBEntity(a.Instrument)
	if err != nil {
		return dbActivity{}, err
	}

	ret := dbActivity{
		Actor:      &actor,
		Object:     &object,
		Target:     &target,
		Result:     &result,
		Origin:     &origin,
		Instrument: &instrument,
	}

	var rest interface{}
	switch typ {
	case "IntransitiveActivity", "Arrive", "Listen", "Read", "Travel":
	case "Activity", "Accept", "Add", "Announce", "Block", "Create", "Delete", "Dislike", "Flag", "Follow", "Ignore", "Invite", "Join", "Leave", "Like", "Move", "Offer", "Reject", "Remove", "TentativeReject", "TentativeAccept", "Undo", "Update", "View":
		a, ok := activity.(activitystreams.TransitiveActivityIface)
		if !ok {
			return dbActivity{}, errors.New("Could not convert activity to TransitiveActivityIface")
		}
		dbActv, err := toDBTransitiveActivity(a)
		if err != nil {
			return dbActivity{}, err
		}
		dbActv.Activity = &ret
		rest = dbActv
	case "Question":
		q, ok := activity.(*activitystreams.Question)
		if !ok {
			return dbActivity{}, errors.New("Could not convert activity to Question")
		}
		dbQ := toDBQuestion(q)
		dbQ.Activity = &ret
		rest = dbQ
	}
	ret.Rest = rest

	return ret, nil
}

func toDBTransitiveActivity(activity activitystreams.TransitiveActivityIface) (dbTransitiveActivity, error) {
	concreteActivity := activitystreams.ToTransitiveActivity(activity)

	object, err := toDBEntity(concreteActivity.Object)
	if err != nil {
		return dbTransitiveActivity{}, err
	}

	return dbTransitiveActivity{
		Object: &object,
	}, nil
}

func toDBQuestion(question *activitystreams.Question) dbQuestion {
	return dbQuestion{}
}

func toDBActor(actor activitystreams.ActorIface) dbActor {
	dbActor := dbActor{}

	return dbActor
}

func toDBCollection(collection activitystreams.CollectionIface) dbCollection {
	return dbCollection{}
}
