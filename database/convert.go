package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/database/util"
	"github.com/brandonsides/pubblr/util/either"
)

func toDBEntity(entity activitystreams.EntityIface) (dbEntity, error) {
	typ, err := entity.Type()
	if err != nil {
		return dbEntity{}, err
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
	name := sql.NullString{
		String: e.Name,
		Valid:  e.Name != "",
	}

	ret := dbEntity{
		Type: typ,
		Name: name,
		MediaType: sql.NullString{
			String: e.MediaType,
			Valid:  true,
		},
		Preview:      &preview,
		AttributedTo: attributedTo,
	}

	var rest interface{} = nil
	switch typ {
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
	ret.Rest = rest

	return ret, nil
}

func fromDBEntity(entity dbEntity) (activitystreams.EntityIface, error) {
	var retEntity activitystreams.Entity
	retEntity.MediaType = entity.MediaType.String
	retEntity.Name = entity.Name.String
	var err error
	retEntity.Preview, err = fromDBEntity(*entity.Preview)
	if err != nil {
		return nil, err
	}
	retEntity.AttributedTo = make([]activitystreams.EntityIface, len(entity.AttributedTo))
	for i, v := range entity.AttributedTo {
		retEntity.AttributedTo[i], err = fromDBEntity(v)
		if err != nil {
			return nil, err
		}
	}

	var ret activitystreams.EntityIface
	switch entity.Type {
	case "Link", "Mention":
		var err error
		restLink, ok := entity.Rest.(dbLink)
		if !ok {
			return nil, errors.New("Could not convert entity.Rest to dbLink")
		}
		retLink, err := fromDBLink(restLink)
		if err != nil {
			return nil, err
		}
		activitystreams.ToLink(retLink).Entity = retEntity
		ret = retLink
	default:
		var err error
		restObject, ok := entity.Rest.(dbObject)
		if !ok {
			return nil, errors.New("Could not convert entity.Rest to dbObject")
		}
		retObject, err := fromDBObject(restObject)
		if err != nil {
			return nil, err
		}
		activitystreams.ToObject(retObject).Entity = retEntity
		ret = retObject
	}
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

func fromDBLink(link dbLink) (activitystreams.LinkIface, error) {
	var l activitystreams.LinkIface = &activitystreams.Link{
		Href:     link.Href.String,
		HrefLang: link.Hreflang.String,
		Rel:      util.StringArray(link.Rel),
		Height:   link.Height,
		Width:    link.Width,
	}

	if link.Entity.Type == "Mention" {
		l = &activitystreams.Mention{
			Link: *activitystreams.ToLink(l),
		}
	}

	return l, nil
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

func fromDBObject(object dbObject) (activitystreams.ObjectIface, error) {
	var retObject activitystreams.Object

	retObject.Attachment = make([]activitystreams.EntityIface, len(object.Attachment))
	for _, v := range object.Attachment {
		attachment, err := fromDBEntity(v)
		if err != nil {
			return nil, err
		}
		retObject.Attachment = append(retObject.Attachment, attachment)
	}

	retObject.Audience = make([]activitystreams.EntityIface, len(object.Audience))
	for _, v := range object.Audience {
		audienceMember, err := fromDBEntity(v)
		if err != nil {
			return nil, err
		}
		retObject.Audience = append(retObject.Audience, audienceMember)
	}

	retObject.Bcc = make([]activitystreams.EntityIface, len(object.Bcc))
	for _, v := range object.Bcc {
		bccMember, err := fromDBEntity(v)
		if err != nil {
			return nil, err
		}
		retObject.Bcc = append(retObject.Bcc, bccMember)
	}

	retObject.Bto = make([]activitystreams.EntityIface, len(object.Bto))
	for _, v := range object.Bto {
		btoMember, err := fromDBEntity(v)
		if err != nil {
			return nil, err
		}
		retObject.Bto = append(retObject.Bto, btoMember)
	}

	retObject.Cc = make([]activitystreams.EntityIface, len(object.Cc))
	for _, v := range object.Cc {
		ccMember, err := fromDBEntity(v)
		if err != nil {
			return nil, err
		}
		retObject.Cc = append(retObject.Cc, ccMember)
	}

	retObject.InReplyTo = make([]activitystreams.EntityIface, len(object.InReplyTo))
	for _, v := range object.InReplyTo {
		inReplyTo, err := fromDBEntity(v)
		if err != nil {
			return nil, err
		}
		retObject.InReplyTo = append(retObject.InReplyTo, inReplyTo)
	}

	retObject.Location = make([]activitystreams.EntityIface, len(object.Location))
	for _, v := range object.Location {
		location, err := fromDBEntity(v)
		if err != nil {
			return nil, err
		}
		retObject.Location = append(retObject.Location, location)
	}

	retObject.Tag = make([]activitystreams.EntityIface, len(object.Tag))
	for _, v := range object.Tag {
		tag, err := fromDBEntity(v)
		if err != nil {
			return nil, err
		}
		retObject.Tag = append(retObject.Tag, tag)
	}

	retObject.To = make([]activitystreams.EntityIface, len(object.To))
	for _, v := range object.To {
		toMember, err := fromDBEntity(v)
		if err != nil {
			return nil, err
		}
		retObject.To = append(retObject.To, toMember)
	}

	var err error
	retObject.Context, err = fromDBEntity(*object.Context)
	if err != nil {
		return nil, err
	}

	retObject.Generator, err = fromDBEntity(*object.Generator)
	if err != nil {
		return nil, err
	}

	retObject.Icon, err = fromDBEntity(*object.Icon)
	if err != nil {
		return nil, err
	}

	retObject.Image, err = fromDBEntity(*object.Image)
	if err != nil {
		return nil, err
	}

	if object.LinkUrl != nil {
		linkUrl, err := fromDBLink(*object.LinkUrl)
		if err != nil {
			return nil, err
		}
		retObject.URL = either.Right[string](linkUrl)
	} else {
		retObject.URL = either.Left[string, activitystreams.LinkIface](object.StringUrl.String)
	}

	retObject.Content = object.Content.String

	retObject.Duration = (*time.Duration)(&object.Duration.Int64)

	retObject.EndTime = &object.EndTime.Time

	retObject.Published = &object.Entity.CreatedAt

	retObject.StartTime = &object.StartTime.Time

	retObject.Summary = object.Summary.String

	retObject.Updated = &object.Entity.UpdatedAt

	var ret activitystreams.ObjectIface
	switch object.Entity.Type {
	case "Object":
		ret = &retObject
	case "Relationship":
		rel, ok := object.Rest.(dbRelationship)
		if !ok {
			return nil, errors.New("Could not convert object.Rest to dbRelationship")
		}
		retRel, err := fromDBRelationship(rel)
		if err != nil {
			return nil, err
		}
		retRel.Object = retObject
		ret = retRel
	case "Article":
		ret = &activitystreams.Article{
			Object: retObject,
		}
	case "Document":
		ret = &activitystreams.Document{
			Object: retObject,
		}
	case "Audio":
		ret = &activitystreams.Audio{
			Object: retObject,
		}
	case "Image":
		ret = &activitystreams.Image{
			Object: retObject,
		}
	case "Video":
		ret = &activitystreams.Video{
			Object: retObject,
		}
	case "Note":
		ret = &activitystreams.Note{
			Object: retObject,
		}
	case "Page":
		ret = &activitystreams.Page{
			Object: retObject,
		}
	case "Event":
		ret = &activitystreams.Event{
			Object: retObject,
		}
	case "Place":
		place, ok := object.Rest.(dbPlace)
		if !ok {
			return nil, errors.New("Could not convert object.Rest to dbPlace")
		}
		retPlace := fromDBPlace(place)
		retPlace.Object = retObject
		ret = retPlace
	case "Profile":
		profile, ok := object.Rest.(dbProfile)
		if !ok {
			return nil, errors.New("Could not convert object.Rest to dbProfile")
		}
		retProfile, err := fromDBProfile(profile)
		if err != nil {
			return nil, err
		}
		retProfile.Object = retObject
		ret = retProfile
	case "Person", "Service", "Group", "Organization", "Application":
		actor, ok := object.Rest.(dbActor)
		if !ok {
			return nil, errors.New("Could not convert object.Rest to dbActor")
		}
		retActor := fromDBActor(actor)
		activitystreams.ToActor(retActor).Object = retObject
		ret = retActor
	case "Collection", "OrderedCollection":
		collection, ok := object.Rest.(dbCollection)
		if !ok {
			return nil, errors.New("Could not convert object.Rest to dbCollection")
		}
		retCollection := fromDBCollection(collection)
		activitystreams.ToCollection(retCollection).Object = retObject
		ret = retCollection
	default:
		activity, ok := object.Rest.(dbActivity)
		if !ok {
			return nil, errors.New("Could not convert object.Rest to dbActivity")
		}
		retActivity, err := fromDBActivity(activity)
		if err != nil {
			return nil, err
		}
		activitystreams.ToIntransitiveActivity(retActivity).Object = retObject
		ret = retActivity
	}

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

func fromDBRelationship(relationship dbRelationship) (*activitystreams.Relationship, error) {
	subj, err := fromDBEntity(*relationship.Subject)
	if err != nil {
		return nil, err
	}

	obj, err := fromDBEntity(*relationship.Obj)
	if err != nil {
		return nil, err
	}

	rel, err := fromDBObject(*relationship.Relationship)
	if err != nil {
		return nil, err
	}

	return &activitystreams.Relationship{
		Subject:      subj,
		Obj:          obj,
		Relationship: rel,
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

func fromDBPlace(place dbPlace) *activitystreams.Place {
	return &activitystreams.Place{
		Accuracy:  place.Accuracy.Float64,
		Altitude:  place.Altitude.Float64,
		Latitude:  place.Latitude.Float64,
		Longitude: place.Longitude.Float64,
		Radius:    place.Radius.Float64,
		Units:     place.Units.String,
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

func fromDBProfile(profile dbProfile) (*activitystreams.Profile, error) {
	describedEntity, err := fromDBEntity(*profile.Describes)
	if err != nil {
		return nil, err
	}
	describes, ok := describedEntity.(activitystreams.ObjectIface)
	if !ok {
		return nil, errors.New("Could not convert described entity to ObjectIface")
	}

	return &activitystreams.Profile{
		Describes: describes,
	}, nil
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
		dbQ, err := toDBQuestion(q)
		if err != nil {
			return dbActivity{}, err
		}
		dbQ.Activity = &ret
		rest = dbQ
	}
	ret.Rest = rest

	return ret, nil
}

func fromDBActivity(activity dbActivity) (activitystreams.ActivityIface, error) {
	var a activitystreams.IntransitiveActivity

	var err error
	a.Actor, err = fromDBEntity(*activity.Actor)
	if err != nil {
		return nil, err
	}

	a.Target, err = fromDBEntity(*activity.Target)
	if err != nil {
		return nil, err
	}

	a.Result, err = fromDBEntity(*activity.Result)
	if err != nil {
		return nil, err
	}

	a.Origin, err = fromDBEntity(*activity.Origin)
	if err != nil {
		return nil, err
	}

	a.Instrument, err = fromDBEntity(*activity.Instrument)
	if err != nil {
		return nil, err
	}

	var ret activitystreams.ActivityIface
	switch activity.Object.Entity.Type {
	case "IntransitiveActivity":
		ret = &a
	case "Arrive":
		ret = &activitystreams.Arrive{
			IntransitiveActivity: a,
		}
	case "Listen":
		ret = &activitystreams.Listen{
			IntransitiveActivity: a,
		}
	case "Read":
		ret = &activitystreams.Read{
			IntransitiveActivity: a,
		}
	case "Travel":
		ret = &activitystreams.Travel{
			IntransitiveActivity: a,
		}
	case "Question":
		question, ok := activity.Rest.(dbQuestion)
		if !ok {
			return nil, errors.New("Could not convert activity.Rest to dbQuestion")
		}
		retQuestion, err := fromDBQuestion(question)
		if err != nil {
			return nil, err
		}
		activitystreams.ToQuestion(retQuestion).IntransitiveActivity = a
		ret = retQuestion
	default:
		transitiveActivity, ok := activity.Rest.(dbTransitiveActivity)
		if !ok {
			return nil, errors.New("Could not convert activity.Rest to dbTransitiveActivity")
		}
		retTransitiveActivity, err := fromDBTransitiveActivity(transitiveActivity)
		if err != nil {
			return nil, err
		}
		activitystreams.ToTransitiveActivity(retTransitiveActivity).IntransitiveActivity = a
		ret = retTransitiveActivity
	}

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

func fromDBTransitiveActivity(activity dbTransitiveActivity) (activitystreams.TransitiveActivityIface, error) {
	var a activitystreams.TransitiveActivity

	var err error
	a.Object, err = fromDBEntity(*activity.Object)
	if err != nil {
		return nil, err
	}

	var ret activitystreams.TransitiveActivityIface
	switch activity.Activity.Object.Entity.Type {
	case "Activity":
		ret = &a
	case "Accept":
		ret = &activitystreams.Accept{
			TransitiveActivity: a,
		}
	case "TentativeAccept":
		ret = &activitystreams.TentativeAccept{
			Accept: activitystreams.Accept{
				TransitiveActivity: a,
			},
		}
	case "Add":
		ret = &activitystreams.Add{
			TransitiveActivity: a,
		}
	case "Create":
		ret = &activitystreams.Create{
			TransitiveActivity: a,
		}
	case "Delete":
		ret = &activitystreams.Delete{
			TransitiveActivity: a,
		}
	case "Follow":
		ret = &activitystreams.Follow{
			TransitiveActivity: a,
		}
	case "Ignore":
		ret = &activitystreams.Ignore{
			TransitiveActivity: a,
		}
	case "Join":
		ret = &activitystreams.Join{
			TransitiveActivity: a,
		}
	case "Leave":
		ret = &activitystreams.Leave{
			TransitiveActivity: a,
		}
	case "Like":
		ret = &activitystreams.Like{
			TransitiveActivity: a,
		}
	case "Offer":
		ret = &activitystreams.Offer{
			TransitiveActivity: a,
		}
	case "Invite":
		ret = &activitystreams.Invite{
			Offer: activitystreams.Offer{
				TransitiveActivity: a,
			},
		}
	case "Reject":
		ret = &activitystreams.Reject{
			TransitiveActivity: a,
		}
	case "TentativeReject":
		ret = &activitystreams.TentativeReject{
			Reject: activitystreams.Reject{
				TransitiveActivity: a,
			},
		}
	case "Remove":
		ret = &activitystreams.Remove{
			TransitiveActivity: a,
		}
	case "Undo":
		ret = &activitystreams.Undo{
			TransitiveActivity: a,
		}
	case "Update":
		ret = &activitystreams.Update{
			TransitiveActivity: a,
		}
	case "View":
		ret = &activitystreams.View{
			TransitiveActivity: a,
		}
	case "Move":
		ret = &activitystreams.Move{
			TransitiveActivity: a,
		}
	case "Announce":
		ret = &activitystreams.Announce{
			TransitiveActivity: a,
		}
	case "Block":
		ret = &activitystreams.Block{
			Ignore: activitystreams.Ignore{
				TransitiveActivity: a,
			},
		}
	case "Flag":
		ret = &activitystreams.Flag{
			TransitiveActivity: a,
		}
	case "Dislike":
		ret = &activitystreams.Dislike{
			TransitiveActivity: a,
		}
	default:
		return nil, errors.New("Unknown activity type")
	}

	return ret, nil
}

func toDBQuestion(question activitystreams.QuestionIface) (dbQuestion, error) {
	var ret dbQuestion

	multiQ, ok := question.(*activitystreams.MultiAnswerQuestion)
	if ok {
		ret.QuestionType = sql.NullString{
			String: "MultiAnswerQuestion",
			Valid:  true,
		}
		for _, answer := range multiQ.AnyOf {
			dbAnswer, err := toDBEntity(answer)
			if err != nil {
				return dbQuestion{}, err
			}
			ret.Answers = append(ret.Answers, dbAnswer)
		}
		return ret, nil
	}

	singleQ, ok := question.(*activitystreams.SingleAnswerQuestion)
	if ok {
		ret.QuestionType = sql.NullString{
			String: "SingleAnswerQuestion",
			Valid:  true,
		}
		for _, answer := range singleQ.OneOf {
			dbAnswer, err := toDBEntity(answer)
			if err != nil {
				return dbQuestion{}, err
			}
			ret.Answers = append(ret.Answers, dbAnswer)
		}
		return ret, nil
	}

	closedQ, ok := question.(*activitystreams.ClosedQuestion)
	if ok {
		ret.QuestionType = sql.NullString{
			String: "ClosedQuestion",
			Valid:  true,
		}
		answer, err := toDBEntity(closedQ.Closed)
		if err != nil {
			return dbQuestion{}, err
		}
		ret.Answers = append(ret.Answers, answer)
		return ret, nil
	}

	ret.QuestionType = sql.NullString{
		String: "Question",
		Valid:  true,
	}
	return ret, nil
}

func fromDBQuestion(question dbQuestion) (activitystreams.QuestionIface, error) {
	var answers []activitystreams.EntityIface
	for _, dbanswer := range question.Answers {
		answer, err := fromDBEntity(dbanswer)
		if err != nil {
			return nil, err
		}
		answers = append(answers, answer)
	}

	switch question.QuestionType.String {
	case "MultiAnswerQuestion":
		return &activitystreams.MultiAnswerQuestion{
			AnyOf: answers,
		}, nil
	case "SingleAnswerQuestion":
		return &activitystreams.SingleAnswerQuestion{
			OneOf: answers,
		}, nil
	case "ClosedQuestion":
		if len(answers) != 1 {
			return nil, errors.New("ClosedQuestion must have exactly one answer")
		}
		return &activitystreams.ClosedQuestion{
			Closed: answers[0],
		}, nil
	default:
		return &activitystreams.Question{}, nil
	}
}

func toDBActor(actor activitystreams.ActorIface) dbActor {
	concreteActor := activitystreams.ToActor(actor)
	return dbActor{
		PreferredUsername: sql.NullString{
			String: concreteActor.PreferredUsername,
			Valid:  concreteActor.PreferredUsername != "",
		},
	}
}

func fromDBActor(dbactor dbActor) activitystreams.ActorIface {
	actor := activitystreams.Actor{
		PreferredUsername: dbactor.PreferredUsername.String,
	}
	switch dbactor.Object.Entity.Type {
	case "Person":
		return &activitystreams.Person{
			Actor: actor,
		}
	case "Service":
		return &activitystreams.Service{
			Actor: actor,
		}
	case "Group":
		return &activitystreams.Group{
			Actor: actor,
		}
	case "Organization":
		return &activitystreams.Organization{
			Actor: actor,
		}
	case "Application":
		return &activitystreams.Application{
			Actor: actor,
		}
	default:
		return &actor
	}
}

func toDBCollection(collection activitystreams.CollectionIface) dbCollection {
	concreteCollection := activitystreams.ToCollection(collection)

	items := make([]dbEntity, len(concreteCollection.Items))
	for i, v := range concreteCollection.Items {
		var err error
		items[i], err = toDBEntity(v)
		if err != nil {
			return dbCollection{}
		}
	}

	return dbCollection{
		Items:   items,
		Ordered: concreteCollection.Ordered,
	}
}

func fromDBCollection(collection dbCollection) activitystreams.CollectionIface {
	items := make([]activitystreams.EntityIface, len(collection.Items))
	for i, v := range collection.Items {
		var err error
		items[i], err = fromDBEntity(v)
		if err != nil {
			return nil
		}
	}

	return &activitystreams.Collection{
		Items:   items,
		Ordered: collection.Ordered,
	}
}
