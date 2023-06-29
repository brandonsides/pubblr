package activitystreams_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandonsides/pubblr/activitystreams"
)

var _ = Describe("Collection", func() {
	actualCollection := activitystreams.Collection{
		Object: activitystreams.Object{
			Entity: activitystreams.Entity{
				Id: "http://example.org/collection",
			},
		},
		TotalItems: 2,
		Ordered:    false,
		Items: []activitystreams.Either[activitystreams.ObjectIface, activitystreams.LinkIface]{
			*activitystreams.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Note{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/note/1",
					},
				},
			}),
			*activitystreams.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Image{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/image/1",
					},
				},
			}),
		},
		Current: activitystreams.Left[*activitystreams.CollectionPage, activitystreams.LinkIface](&activitystreams.CollectionPage{
			Collection: activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/collection?page=1",
					},
				},
			},
		}),
		First: activitystreams.Left[*activitystreams.CollectionPage, activitystreams.LinkIface](&activitystreams.CollectionPage{
			Collection: activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/collection?page=1",
					},
				},
			},
		}),
		Last: activitystreams.Left[*activitystreams.CollectionPage, activitystreams.LinkIface](&activitystreams.CollectionPage{
			Collection: activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/collection?page=2",
					},
				},
			},
		}),
	}

	expectedCollectionMap := map[string]interface{}{
		"id":         "http://example.org/collection",
		"totalItems": 2.0,
		"current": map[string]interface{}{
			"id": "http://example.org/collection?page=1",
		},
		"first": map[string]interface{}{
			"id": "http://example.org/collection?page=1",
		},
		"last": map[string]interface{}{
			"id": "http://example.org/collection?page=2",
		},
	}

	Context("Unordered", func() {
		BeforeEach(func() {
			actualCollection.Ordered = false
			(*actualCollection.Current.Left()).Ordered = false
			(*actualCollection.First.Left()).Ordered = false
			(*actualCollection.Last.Left()).Ordered = false
			delete(expectedCollectionMap, "orderedItems")
			expectedCollectionMap["items"] = []interface{}{
				map[string]interface{}{
					"type": "Note",
					"id":   "http://example.org/note/1",
				},
				map[string]interface{}{
					"type": "Image",
					"id":   "http://example.org/image/1",
				},
			}
			expectedCollectionMap["type"] = "Collection"
			expectedCollectionMap["current"].(map[string]interface{})["type"] = "CollectionPage"
			expectedCollectionMap["first"].(map[string]interface{})["type"] = "CollectionPage"
			expectedCollectionMap["last"].(map[string]interface{})["type"] = "CollectionPage"
		})

		CheckActivityStreamsEntity("Collection", &actualCollection, expectedCollectionMap)
	})

	Context("Ordered", func() {
		BeforeEach(func() {
			actualCollection.Ordered = true
			(*actualCollection.Current.Left()).Ordered = true
			(*actualCollection.First.Left()).Ordered = true
			(*actualCollection.Last.Left()).Ordered = true
			delete(expectedCollectionMap, "items")
			expectedCollectionMap["orderedItems"] = []interface{}{
				map[string]interface{}{
					"type": "Note",
					"id":   "http://example.org/note/1",
				},
				map[string]interface{}{
					"type": "Image",
					"id":   "http://example.org/image/1",
				},
			}
			expectedCollectionMap["type"] = "OrderedCollection"
			expectedCollectionMap["current"].(map[string]interface{})["type"] = "OrderedCollectionPage"
			expectedCollectionMap["first"].(map[string]interface{})["type"] = "OrderedCollectionPage"
			expectedCollectionMap["last"].(map[string]interface{})["type"] = "OrderedCollectionPage"
		})

		Describe("MarshalJSON", func() {
			It("should correctly marshal fully populated type", func() {
				jsonObject, err := actualCollection.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())
				var actual map[string]interface{}
				err = json.Unmarshal(jsonObject, &actual)
				Expect(err).ToNot(HaveOccurred())
				Expect(actual).To(Equal(expectedCollectionMap))
			})
		})

		Describe("Type", func() {
			It("should return correct type", func() {
				Expect(actualCollection.Type()).To(Equal("OrderedCollection"))
			})
		})
	})
})
