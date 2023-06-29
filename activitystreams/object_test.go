package activitystreams_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	//	. "github.com/onsi/gomega"

	"github.com/brandonsides/pubblr/activitystreams"
)

var _ = Describe("Object", func() {
	duration := time.Second * 60
	startTime := time.Date(2023, 6, 18, 9, 47, 0, 0, time.UTC)
	endTime := startTime.Add(duration)
	published := time.Date(2023, 6, 18, 9, 46, 0, 0, time.UTC)
	updated := time.Date(2023, 6, 18, 9, 46, 30, 0, time.UTC)
	actualObject := activitystreams.Object{
		Entity: activitystreams.Entity{
			Id: "http://example.org/~john",
			AttributedTo: []activitystreams.EntityIface{
				&activitystreams.Person{
					Object: activitystreams.Object{
						Entity: activitystreams.Entity{
							Id: "http://example.org/~john",
						},
					},
				},
			},
			Name: "A Simple Note",
		},
		Attachment: []activitystreams.EntityIface{
			&activitystreams.Image{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id:        "http://example.org/~john/picture",
						MediaType: "image/jpeg",
					},
				},
			},
			&activitystreams.Link{
				Entity: activitystreams.Entity{
					Id: "http://example.org/~john/profile",
				},
				Href: "http://example.org/~john/profile",
			},
		},
		Bcc: []activitystreams.EntityIface{
			&activitystreams.Person{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/~alice",
					},
				},
			},
		},
		Bto: []activitystreams.EntityIface{
			&activitystreams.Person{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/~bob",
					},
				},
			},
		},
		Cc: []activitystreams.EntityIface{
			&activitystreams.Person{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/~eve",
					},
				},
			},
		},
		Context: &activitystreams.Object{
			Entity: activitystreams.Entity{
				Id: "http://example.org/contexts/1",
			},
		},
		Generator: &activitystreams.Object{
			Entity: activitystreams.Entity{
				Id: "http://example.org/generator",
			},
		},
		Icon: &activitystreams.Image{
			Object: activitystreams.Object{
				Entity: activitystreams.Entity{
					Name: "John's Avatar",
				},
				URL: activitystreams.Left[string, activitystreams.LinkIface]("http://example.org/~john/avatar.jpg"),
			},
		},
		Image: &activitystreams.Image{
			Object: activitystreams.Object{
				Entity: activitystreams.Entity{
					Name: "John's Header",
				},
				URL: activitystreams.Left[string, activitystreams.LinkIface]("http://example.org/~john/header.jpg"),
			},
		},
		InReplyTo: []activitystreams.EntityIface{
			&activitystreams.Object{
				Entity: activitystreams.Entity{
					Id: "http://example.org/posts/1",
				},
			},
		},
		Location: []activitystreams.EntityIface{
			&activitystreams.Place{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Name: "Work",
					},
				},
			},
		},
		Preview: &activitystreams.Object{
			Entity: activitystreams.Entity{
				Id: "http://example.org/~john/preview",
			},
		},
		Replies: &activitystreams.Collection{
			Object: activitystreams.Object{
				Entity: activitystreams.Entity{
					Id: "http://example.org/~john/replies",
				},
			},
		},
		Tag: []activitystreams.EntityIface{
			&activitystreams.Object{
				Entity: activitystreams.Entity{
					Id: "http://example.org/tags/1",
				},
			},
		},
		To: []activitystreams.EntityIface{
			&activitystreams.Person{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/~alice",
					},
				},
			},
		},
		URL:       activitystreams.Left[string, activitystreams.LinkIface]("http://example.org/~john"),
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

		CheckActivityStreamsEntity("Object", &actualObject, expectedObjectMap)
	})

	Describe("Relationship", func() {
		BeforeEach(func() {
			expectedObjectMap["type"] = "Relationship"
		})

		AfterEach(func() {
			delete(expectedObjectMap, "type")
		})

		CheckActivityStreamsEntity("Relationship", &activitystreams.Relationship{
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

		CheckActivityStreamsEntity("Article", &activitystreams.Article{
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

		CheckActivityStreamsEntity("Document", &activitystreams.Document{
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

		CheckActivityStreamsEntity("Audio", &activitystreams.Audio{
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

		CheckActivityStreamsEntity("Image", &activitystreams.Image{
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

		CheckActivityStreamsEntity("Video", &activitystreams.Video{
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

		CheckActivityStreamsEntity("Note", &activitystreams.Note{
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

		CheckActivityStreamsEntity("Page", &activitystreams.Page{
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

		CheckActivityStreamsEntity("Event", &activitystreams.Event{
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

		CheckActivityStreamsEntity("Place", &activitystreams.Place{
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

		CheckActivityStreamsEntity("Profile", &activitystreams.Profile{
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

		CheckActivityStreamsEntity("Tombstone", &activitystreams.Tombstone{
			Object: actualObject,
		}, expectedObjectMap)
	})
})
