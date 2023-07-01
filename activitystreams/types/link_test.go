package types_test

import (
	. "github.com/onsi/ginkgo/v2"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/activitystreams/testutil"
	"github.com/brandonsides/pubblr/activitystreams/types"
)

var _ = Describe("Link", func() {
	height := uint64(100)
	width := uint64(200)
	actualLink := types.Link{
		Entity: activitystreams.Entity{
			Id: "http://example.com/abc",
			AttributedTo: []activitystreams.EntityIface{
				types.ObjectIface(&types.Person{
					Object: types.Object{
						Entity: activitystreams.Entity{
							Id: "http://example.com/~john",
						},
					},
				}),
			},
			Name:      "A Link",
			MediaType: "text/html",
		},
		Href:     "http://example.com/abc",
		HrefLang: "en",
		Preview: &types.Object{
			Entity: activitystreams.Entity{
				Id: "http://example.com/abc/preview",
			},
		},
		Height: &height,
		Width:  &width,
		Rel:    []string{"me"},
	}
	expectedLinkMap := map[string]interface{}{
		"id":           "http://example.com/abc",
		"attributedTo": []interface{}{map[string]interface{}{"id": "http://example.com/~john", "type": "Person"}},
		"preview":      map[string]interface{}{"id": "http://example.com/abc/preview", "type": "Object"},
		"name":         "A Link",
		"height":       float64(100),
		"width":        float64(200),
		"href":         "http://example.com/abc",
		"mediaType":    "text/html",
		"hreflang":     "en",
		"rel":          []interface{}{"me"},
	}

	Describe("Link", func() {
		BeforeEach(func() {
			expectedLinkMap["type"] = "Link"
		})

		AfterEach(func() {
			delete(expectedLinkMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Link", &actualLink, expectedLinkMap)
	})

	Describe("Mention", func() {
		actualMention := types.Mention{
			Link: actualLink,
		}

		BeforeEach(func() {
			expectedLinkMap["type"] = "Mention"
		})

		AfterEach(func() {
			delete(expectedLinkMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Mention", &actualMention, expectedLinkMap)
	})
})
