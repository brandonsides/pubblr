package activitystreams

import (
	"encoding/json"

	"github.com/brandonsides/pubblr/util"
)

const (
	CollectionTypeOrdered   = "OrderedCollection"
	CollectionTypeUnordered = "Collection"
)

type CollectionIface interface {
	ObjectIface
	collection() *Collection
}

func ToCollection(c CollectionIface) *Collection {
	return c.collection()
}

func MarshalCollection(c CollectionIface) ([]byte, error) {
	var mapped map[string]interface{}
	j, err := MarshalObject(c)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(j, &mapped)
	if err != nil {
		return nil, err
	}

	if c.Type() == CollectionTypeOrdered {
		mapped["orderedItems"] = mapped["items"]
		delete(mapped, "items")
	}

	return json.Marshal(mapped)
}

type Collection struct {
	Object
	Ordered    bool                               `json:"-"`
	TotalItems uint64                             `json:"totalItems,omitempty"`
	Current    *util.Either[CollectionPage, Link] `json:"current,omitempty"`
	First      *util.Either[CollectionPage, Link] `json:"first,omitempty"`
	Last       *util.Either[CollectionPage, Link] `json:"last,omitempty"`
	Items      []util.Either[Object, Link]        `json:"items,omitempty"`
}

type rawCollection struct {
	*Collection
}

func (c *Collection) collection() *Collection {
	return c
}

func (c *Collection) MarshalJSON() ([]byte, error) {
	return MarshalCollection(rawCollection{
		c,
	})
}

func (c *Collection) Type() string {
	if c.Ordered {
		return CollectionTypeOrdered
	}
	return CollectionTypeUnordered
}

type CollectionPage struct {
	Collection
	PartOf *util.Either[Collection, Link]     `json:"partOf,omitempty"`
	Next   *util.Either[CollectionPage, Link] `json:"next,omitempty"`
	Prev   *util.Either[CollectionPage, Link] `json:"prev,omitempty"`
}
