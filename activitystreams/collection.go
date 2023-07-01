package activitystreams

import (
	"encoding/json"

	pkgJson "github.com/brandonsides/pubblr/activitystreams/json"
	"github.com/brandonsides/pubblr/activitystreams/util"
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

/*
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
*/

type Collection struct {
	Object
	Ordered    bool                                     `json:"-"`
	TotalItems uint64                                   `json:"totalItems,omitempty"`
	Current    *util.Either[*CollectionPage, LinkIface] `json:"current,omitempty"`
	First      *util.Either[*CollectionPage, LinkIface] `json:"first,omitempty"`
	Last       *util.Either[*CollectionPage, LinkIface] `json:"last,omitempty"`
	Items      []*util.Either[ObjectIface, LinkIface]   `json:"items,omitempty"`
}

func (c *Collection) collection() *Collection {
	return c
}

func (c *Collection) MarshalJSON() ([]byte, error) {
	retJson, err := pkgJson.MarshalEntity(c)
	if err != nil {
		return nil, err
	}
	typ, err := c.Type()
	if err != nil {
		return nil, err
	}
	if typ == CollectionTypeOrdered {
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
	PartOf *util.Either[Collection, Link]     `json:"partOf,omitempty"`
	Next   *util.Either[CollectionPage, Link] `json:"next,omitempty"`
	Prev   *util.Either[CollectionPage, Link] `json:"prev,omitempty"`
}

func (c *CollectionPage) Type() (string, error) {
	if c.Ordered {
		return "OrderedCollectionPage", nil
	}
	return "CollectionPage", nil
}

func (c *CollectionPage) MarshalJSON() ([]byte, error) {
	return pkgJson.MarshalEntity(c)
}
