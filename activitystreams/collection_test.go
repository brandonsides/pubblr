package activitystreams_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/activitystreams/testutil"
	"github.com/brandonsides/pubblr/util/either"
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
		Items: []*either.Either[activitystreams.ObjectIface, activitystreams.LinkIface]{
			either.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Note{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/note/1",
					},
				},
			}),
			either.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Image{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/image/1",
					},
				},
			}),
		},
		Current: either.Left[*activitystreams.CollectionPage, activitystreams.LinkIface](&activitystreams.CollectionPage{
			Collection: activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/collection?page=1",
					},
				},
			},
		}),
		First: either.Left[*activitystreams.CollectionPage, activitystreams.LinkIface](&activitystreams.CollectionPage{
			Collection: activitystreams.Collection{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/collection?page=1",
					},
				},
			},
		}),
		Last: either.Left[*activitystreams.CollectionPage, activitystreams.LinkIface](&activitystreams.CollectionPage{
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
		"current":    "http://example.org/collection?page=1",
		"first":      "http://example.org/collection?page=1",
		"last":       "http://example.org/collection?page=2",
	}

	Context("Unordered", func() {
		BeforeEach(func() {
			actualCollection.Ordered = false
			(*actualCollection.Current.Left()).Ordered = false
			(*actualCollection.First.Left()).Ordered = false
			(*actualCollection.Last.Left()).Ordered = false
			delete(expectedCollectionMap, "orderedItems")
			expectedCollectionMap["items"] = []interface{}{
				"http://example.org/note/1",
				"http://example.org/image/1",
			}
			expectedCollectionMap["type"] = "Collection"
		})

		testutil.CheckActivityStreamsEntity("Collection", &actualCollection, expectedCollectionMap)
	})

	Context("Ordered", func() {
		BeforeEach(func() {
			actualCollection.Ordered = true
			(*actualCollection.Current.Left()).Ordered = true
			(*actualCollection.First.Left()).Ordered = true
			(*actualCollection.Last.Left()).Ordered = true
			delete(expectedCollectionMap, "items")
			expectedCollectionMap["orderedItems"] = []interface{}{
				"http://example.org/note/1",
				"http://example.org/image/1",
			}
			expectedCollectionMap["type"] = "OrderedCollection"
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
