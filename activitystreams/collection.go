package activitystreams

import (
	"encoding/json"
	"fmt"

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
	collection, err := json.Marshal(&c.Object)
	if err != nil {
		return nil, err
	}

	var collectionMap map[string]json.RawMessage
	err = json.Unmarshal(collection, &collectionMap)
	if err != nil {
		return nil, err
	}

	if c.TotalItems > 0 {
		collectionMap["totalItems"] = []byte(fmt.Sprintf("%d", c.TotalItems))
	}

	if c.Current != nil {
		var current EntityIface
		if c.Current.IsLeft() {
			current = *c.Current.Left()
		} else {
			current = *c.Current.Right()
		}
		collectionMap["current"] = []byte(fmt.Sprintf("%q", ToEntity(current).Id))
	}

	if c.First != nil {
		var first EntityIface
		if c.First.IsLeft() {
			first = *c.First.Left()
		} else {
			first = *c.First.Right()
		}
		collectionMap["first"] = []byte(fmt.Sprintf("%q", ToEntity(first).Id))
	}

	if c.Last != nil {
		var last EntityIface
		if c.Last.IsLeft() {
			last = *c.Last.Left()
		} else {
			last = *c.Last.Right()
		}
		collectionMap["last"] = []byte(fmt.Sprintf("%q", ToEntity(last).Id))
	}

	if len(c.Items) > 0 {
		itemKey := "items"
		if c.Ordered {
			itemKey = "orderedItems"
		}

		itemIds := make([]string, len(c.Items))
		for i, item := range c.Items {
			if item.IsLeft() {
				itemIds[i] = ToEntity(*item.Left()).Id
			} else {
				itemIds[i] = ToEntity(*item.Right()).Id
			}
		}

		items, err := json.Marshal(itemIds)
		if err != nil {
			return nil, err
		}

		collectionMap[itemKey] = items
	}

	if c.Ordered {
		collectionMap["type"] = []byte(fmt.Sprintf("%q", CollectionTypeOrdered))
	} else {
		collectionMap["type"] = []byte(fmt.Sprintf("%q", CollectionTypeUnordered))
	}

	return json.Marshal(collectionMap)
}

func (c *Collection) Type() (string, error) {
	if c.Ordered {
		return CollectionTypeOrdered, nil
	}
	return CollectionTypeUnordered, nil
}

type CollectionPage struct {
	Collection
	PartOf *either.Either[*Collection, LinkIface]     `json:"partOf,omitempty"`
	Next   *either.Either[*CollectionPage, LinkIface] `json:"next,omitempty"`
	Prev   *either.Either[*CollectionPage, LinkIface] `json:"prev,omitempty"`
}

func (c *CollectionPage) Type() (string, error) {
	if c.Ordered {
		return "OrderedCollectionPage", nil
	}
	return "CollectionPage", nil
}

func (c *CollectionPage) MarshalJSON() ([]byte, error) {
	collectionPage, err := json.Marshal(&c.Collection)
	if err != nil {
		return nil, err
	}

	var collectionPageMap map[string]json.RawMessage
	err = json.Unmarshal(collectionPage, &collectionPageMap)
	if err != nil {
		return nil, err
	}

	if c.PartOf != nil {
		var partOf EntityIface
		if c.PartOf.IsLeft() {
			partOf = *c.PartOf.Left()
		} else {
			partOf = *c.PartOf.Right()
		}
		collectionPageMap["partOf"] = []byte(fmt.Sprintf("%q", ToEntity(partOf).Id))
	}

	if c.Next != nil {
		var next EntityIface
		if c.Next.IsLeft() {
			next = *c.Next.Left()
		} else {
			next = *c.Next.Right()
		}
		collectionPageMap["next"] = []byte(fmt.Sprintf("%q", ToEntity(next).Id))
	}

	if c.Prev != nil {
		var prev EntityIface
		if c.Prev.IsLeft() {
			prev = *c.Prev.Left()
		} else {
			prev = *c.Prev.Right()
		}
		collectionPageMap["prev"] = []byte(fmt.Sprintf("%q", ToEntity(prev).Id))
	}

	if c.Ordered {
		collectionPageMap["type"] = []byte(fmt.Sprintf("%q", "OrderedCollectionPage"))
	} else {
		collectionPageMap["type"] = []byte(fmt.Sprintf("%q", "CollectionPage"))
	}

	return json.Marshal(collectionPageMap)
}
