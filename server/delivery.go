package server

import "github.com/brandonsides/pubblr/activitystreams"

func (router *PubblrRouter) Deliver(activity activitystreams.ActivityIface) {
	go router.deliver(activity)
}

func (router *PubblrRouter) deliver(a activitystreams.ActivityIface) {
	activity := activitystreams.ToIntransitiveActivity(a)
	recipients := merge(activity.To, activity.Bto, activity.Audience, activity.Bcc, activity.Cc)

	for _, recipient := range recipients {
		router.deliverTo(recipient, a)
	}
}

func (router *PubblrRouter) deliverTo(recipient activitystreams.EntityIface, activity activitystreams.ActivityIface) {
	if router.isLocal(recipient) {
		router.deliverToLocal(recipient, activity)
	} else {
		router.deliverToRemote(recipient, activity)
	}
}

func (router *PubblrRouter) isLocal(entity activitystreams.EntityIface) bool {
	// TODO
	return true
}

func (router *PubblrRouter) deliverToLocal(recipient activitystreams.EntityIface, activity activitystreams.ActivityIface) {
	recipientEntity := activitystreams.ToEntity(recipient)
	shordId := shortId(recipientEntity.Id)

	router.Database.CreateInboxItem(activity, shordId)
}

func (router *PubblrRouter) deliverToRemote(recipient activitystreams.EntityIface, activity activitystreams.ActivityIface) {
	router.Logger.Errorf("deliverToRemote not implemented")
}
