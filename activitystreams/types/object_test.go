package types_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	//	. "github.com/onsi/gomega"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/activitystreams/testutil"
	"github.com/brandonsides/pubblr/activitystreams/types"
	"github.com/brandonsides/pubblr/util/either"
)

var _ = Describe("Object", func() {
	duration := time.Second * 60
	startTime := time.Date(2023, 6, 18, 9, 47, 0, 0, time.UTC)
	endTime := startTime.Add(duration)
	published := time.Date(2023, 6, 18, 9, 46, 0, 0, time.UTC)
	updated := time.Date(2023, 6, 18, 9, 46, 30, 0, time.UTC)
	actualObject := types.Object{
		Entity: activitystreams.Entity{
			Id: "http://example.org/~john",
			AttributedTo: []activitystreams.EntityIface{
				&types.Person{
					Object: types.Object{
						Entity: activitystreams.Entity{
							Id: "http://example.org/~john",
						},
					},
				},
			},
			Name: "A Simple Note",
		},
		Attachment: []activitystreams.EntityIface{
			&types.Image{
				Object: types.Object{
					Entity: activitystreams.Entity{
						Id:        "http://example.org/~john/picture",
						MediaType: "image/jpeg",
					},
				},
			},
			&types.Link{
				Entity: activitystreams.Entity{
					Id: "http://example.org/~john/profile",
				},
				Href: "http://example.org/~john/profile",
			},
		},
		Bcc: []activitystreams.EntityIface{
			&types.Person{
				Object: types.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/~alice",
					},
				},
			},
		},
		Bto: []activitystreams.EntityIface{
			&types.Person{
				Object: types.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/~bob",
					},
				},
			},
		},
		Cc: []activitystreams.EntityIface{
			&types.Person{
				Object: types.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/~eve",
					},
				},
			},
		},
		Context: &types.Object{
			Entity: activitystreams.Entity{
				Id: "http://example.org/contexts/1",
			},
		},
		Generator: &types.Object{
			Entity: activitystreams.Entity{
				Id: "http://example.org/generator",
			},
		},
		Icon: &types.Image{
			Object: types.Object{
				Entity: activitystreams.Entity{
					Name: "John's Avatar",
				},
				URL: either.Left[string, types.LinkIface]("http://example.org/~john/avatar.jpg"),
			},
		},
		Image: &types.Image{
			Object: types.Object{
				Entity: activitystreams.Entity{
					Name: "John's Header",
				},
				URL: either.Left[string, types.LinkIface]("http://example.org/~john/header.jpg"),
			},
		},
		InReplyTo: []activitystreams.EntityIface{
			&types.Object{
				Entity: activitystreams.Entity{
					Id: "http://example.org/posts/1",
				},
			},
		},
		Location: []activitystreams.EntityIface{
			&types.Place{
				Object: types.Object{
					Entity: activitystreams.Entity{
						Name: "Work",
					},
				},
			},
		},
		Preview: &types.Object{
			Entity: activitystreams.Entity{
				Id: "http://example.org/~john/preview",
			},
		},
		Replies: &types.Collection{
			Object: types.Object{
				Entity: activitystreams.Entity{
					Id: "http://example.org/~john/replies",
				},
			},
		},
		Tag: []activitystreams.EntityIface{
			&types.Object{
				Entity: activitystreams.Entity{
					Id: "http://example.org/tags/1",
				},
			},
		},
		To: []activitystreams.EntityIface{
			&types.Person{
				Object: types.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/~alice",
					},
				},
			},
		},
		URL:       either.Left[string, types.LinkIface]("http://example.org/~john"),
		Content:   "This is a simple note",
		Duration:  &duration,
		EndTime:   &endTime,
		Published: &published,
		StartTime: &startTime,
		Summary:   "A simple note",
		Updated:   &updated,
	}
	expectedObjectMap := map[string]interface{}{
		"id": "http://example.org/~john",
		"attachment": []interface{}{
			map[string]interface{}{
				"type":      "Image",
				"id":        "http://example.org/~john/picture",
				"mediaType": "image/jpeg",
			},
			map[string]interface{}{
				"type": "Link",
				"id":   "http://example.org/~john/profile",
				"href": "http://example.org/~john/profile",
			},
		},
		"attributedTo": []interface{}{
			map[string]interface{}{
				"type": "Person",
				"id":   "http://example.org/~john",
			},
		},
		"bcc": []interface{}{
			map[string]interface{}{
				"type": "Person",
				"id":   "http://example.org/~alice",
			},
		},
		"bto": []interface{}{
			map[string]interface{}{
				"type": "Person",
				"id":   "http://example.org/~bob",
			},
		},
		"cc": []interface{}{
			map[string]interface{}{
				"type": "Person",
				"id":   "http://example.org/~eve",
			},
		},
		"context": map[string]interface{}{
			"type": "Object",
			"id":   "http://example.org/contexts/1",
		},
		"generator": map[string]interface{}{
			"type": "Object",
			"id":   "http://example.org/generator",
		},
		"icon": map[string]interface{}{
			"name": "John's Avatar",
			"type": "Image",
			"url":  "http://example.org/~john/avatar.jpg",
		},
		"image": map[string]interface{}{
			"type": "Image",
			"name": "John's Header",
			"url":  "http://example.org/~john/header.jpg",
		},
		"inReplyTo": []interface{}{
			map[string]interface{}{
				"type": "Object",
				"id":   "http://example.org/posts/1",
			},
		},
		"location": []interface{}{
			map[string]interface{}{
				"type": "Place",
				"name": "Work",
			},
		},
		"preview": map[string]interface{}{
			"type": "Object",
			"id":   "http://example.org/~john/preview",
		},
		"replies": map[string]interface{}{
			"type": "Collection",
			"id":   "http://example.org/~john/replies",
		},
		"tag": []interface{}{
			map[string]interface{}{
				"type": "Object",
				"id":   "http://example.org/tags/1",
			},
		},
		"to": []interface{}{
			map[string]interface{}{
				"type": "Person",
				"id":   "http://example.org/~alice",
			},
		},
		"url":       "http://example.org/~john",
		"content":   "This is a simple note",
		"name":      "A Simple Note",
		"duration":  float64(time.Second * 60),
		"endTime":   "2023-06-18T09:48:00Z",
		"published": "2023-06-18T09:46:00Z",
		"startTime": "2023-06-18T09:47:00Z",
		"summary":   "A simple note",
		"updated":   "2023-06-18T09:46:30Z",
	}

	Describe("Object", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Object"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Object", &actualObject, expectedObjectMap)
	})

	Describe("Relationship", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Relationship"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Relationship", &types.Relationship{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Article", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Article"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Article", &types.Article{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Document", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Document"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Document", &types.Document{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Audio", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Audio"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Audio", &types.Audio{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Image", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Image"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Image", &types.Image{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Video", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Video"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Video", &types.Video{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Note", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Note"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Note", &types.Note{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Page", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Page"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Page", &types.Page{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Event", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Event"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Event", &types.Event{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Place", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Place"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Place", &types.Place{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Profile", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Profile"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Profile", &types.Profile{
			Object: actualObject,
		}, expectedObjectMap)
	})

	Describe("Tombstone", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Tombstone"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Tombstone", &types.Tombstone{
			Object: actualObject,
		}, expectedObjectMap)
	})
})
