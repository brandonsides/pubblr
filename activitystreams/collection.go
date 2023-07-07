package activitystreams

import (
	"encoding/json"

	"github.com/brandonsides/pubblr/util/either"
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

type Collection struct {
	Object
	Ordered    bool                                       `json:"-"`
	TotalItems uint64                                     `json:"totalItems,omitempty"`
	Current    *either.Either[*CollectionPage, LinkIface] `json:"current,omitempty"`
	First      *either.Either[*CollectionPage, LinkIface] `json:"first,omitempty"`
	Last       *either.Either[*CollectionPage, LinkIface] `json:"last,omitempty"`
	Items      []*either.Either[ObjectIface, LinkIface]   `json:"items,omitempty"`
}

func (c *Collection) collection() *Collection {
	return c
}

func (c *Collection) MarshalJSON() ([]byte, error) {
	retJson, err := MarshalEntity(c)
	if err != nil {
		return nil, err
	}
	typ, err := c.Type()
	if err != nil {
		return nil, err
	}
	if typ == CollectionTypeOrdered && len(c.Items) > 0 {
		var mapped map[string]interface{}
		err := json.Unmarshal(retJson, &mapped)
		if err != nil {
			return nil, err
		}
		mapped["orderedItems"] = mapped["items"]
		delete(mapped, "items")
		return json.Marshal(mapped)
	}
	return retJson, nil
}

func (c *Collection) Type() (string, error) {
	if c.Ordered {
		return CollectionTypeOrdered, nil
	}
	return CollectionTypeUnordered, nil
}

type CollectionPage struct {
	Collection
	PartOf *either.Either[Collection, Link]     `json:"partOf,omitempty"`
	Next   *either.Either[CollectionPage, Link] `json:"next,omitempty"`
	Prev   *either.Either[CollectionPage, Link] `json:"prev,omitempty"`
}

func (c *CollectionPage) Type() (string, error) {
	if c.Ordered {
		return "OrderedCollectionPage", nil
	}
	return "CollectionPage", nil
}

func (c *CollectionPage) MarshalJSON() ([]byte, error) {
	return MarshalEntity(c)
}
