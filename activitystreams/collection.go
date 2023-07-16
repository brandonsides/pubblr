package activitystreams

import (
	"encoding/json"
	"errors"
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
	Items      []EntityIface                              `json:"items,omitempty"`
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
			itemIds[i] = ToEntity(item).Id
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

func (c *Collection) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	err := c.Object.UnmarshalEntity(u, b)
	if err != nil {
		return err
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(b, &objMap)
	if err != nil {
		return nil
	}

	jsonType, ok := objMap["type"]
	if !ok {
		return errors.New("Collection has no type")
	}
	var typ string
	err = json.Unmarshal(jsonType, &typ)
	if err != nil {
		return err
	}

	switch typ {
	case CollectionTypeOrdered:
		c.Ordered = true
	case CollectionTypeUnordered:
		c.Ordered = false
	default:
		return errors.New("Collection has invalid type: " + typ)
	}

	itemsKey := "items"
	if c.Ordered {
		itemsKey = "orderedItems"
	}

	if totalItems, ok := objMap["totalItems"]; ok {
		err = json.Unmarshal(totalItems, &c.TotalItems)
		if err != nil {
			return err
		}
	}

	if current, ok := objMap["current"]; ok {
		var unmarshalledCurrent CollectionPage
		err = unmarshalledCurrent.UnmarshalEntity(u, current)
		if err != nil {
			return err
		}

		c.Current = either.Left[*CollectionPage, LinkIface](&unmarshalledCurrent)
	}

	if first, ok := objMap["first"]; ok {
		var unmarshalledFirst CollectionPage
		err = unmarshalledFirst.UnmarshalEntity(u, first)
		if err != nil {
			return err
		}

		c.First = either.Left[*CollectionPage, LinkIface](&unmarshalledFirst)
	}

	if last, ok := objMap["last"]; ok {
		var unmarshalledLast CollectionPage
		err = unmarshalledLast.UnmarshalEntity(u, last)
		if err != nil {
			return err
		}

		c.Last = either.Left[*CollectionPage, LinkIface](&unmarshalledLast)
	}

	if items, ok := objMap[itemsKey]; ok {
		var jsonItems []json.RawMessage
		err = json.Unmarshal(items, &jsonItems)
		if err != nil {
			return err
		}

		for _, jsonItem := range jsonItems {
			item, err := u.UnmarshalEntity(jsonItem)
			if err != nil {
				return err
			}

			c.Items = append(c.Items, item)
		}
	}

	return nil
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

func (c *CollectionPage) UnmarshalEntity(u *EntityUnmarshaler, b []byte) error {
	err := c.Collection.UnmarshalEntity(u, b)
	if err != nil {
		return err
	}

	var objMap map[string]json.RawMessage
	err = json.Unmarshal(b, &objMap)
	if err != nil {
		return nil
	}

	if partOf, ok := objMap["partOf"]; ok {
		var unmarshalledPartOf Collection
		err = unmarshalledPartOf.UnmarshalEntity(u, partOf)
		if err != nil {
			return err
		}

		c.PartOf = either.Left[*Collection, LinkIface](&unmarshalledPartOf)
	}

	if next, ok := objMap["next"]; ok {
		var unmarshalledNext CollectionPage
		err = unmarshalledNext.UnmarshalEntity(u, next)
		if err != nil {
			return err
		}

		c.Next = either.Left[*CollectionPage, LinkIface](&unmarshalledNext)
	}

	if prev, ok := objMap["prev"]; ok {
		var unmarshalledPrev CollectionPage
		err = unmarshalledPrev.UnmarshalEntity(u, prev)
		if err != nil {
			return err
		}

		c.Prev = either.Left[*CollectionPage, LinkIface](&unmarshalledPrev)
	}

	return nil
}
