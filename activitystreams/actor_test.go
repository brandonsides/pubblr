package activitystreams_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	//	. "github.com/onsi/gomega"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/activitystreams/testutil"
	"github.com/brandonsides/pubblr/util/either"
)

var _ = Describe("Actor", func() {
	duration := time.Second * 60
	startTime := time.Date(2023, 6, 18, 9, 47, 0, 0, time.UTC)
	endTime := startTime.Add(duration)
	published := time.Date(2023, 6, 18, 9, 46, 0, 0, time.UTC)
	updated := time.Date(2023, 6, 18, 9, 46, 30, 0, time.UTC)
	actualActor := activitystreams.Object{
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
				URL: either.Left[string, activitystreams.LinkIface]("http://example.org/~john/avatar.jpg"),
			},
		},
		Image: &activitystreams.Image{
			Object: activitystreams.Object{
				Entity: activitystreams.Entity{
					Name: "John's Header",
				},
				URL: either.Left[string, activitystreams.LinkIface]("http://example.org/~john/header.jpg"),
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
		URL:       either.Left[string, activitystreams.LinkIface]("http://example.org/~john"),
		Content:   "This is a simple note",
		Duration:  &duration,
		EndTime:   &endTime,
		Published: &published,
		StartTime: &startTime,
		Summary:   "A simple note",
		Updated:   &updated,
	}
	expectedActorMap := map[string]interface{}{
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

	Describe("Application", func() {
		actualApplication := activitystreams.Application{actualActor}
		expectedApplicationMap := expectedActorMap

		BeforeEach(func() {
			expectedApplicationMap["type"] = "Application"
		})

		AfterEach(func() {
			delete(expectedApplicationMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Application", &actualApplication, expectedApplicationMap)
	})

	Describe("Group", func() {
		actualGroup := activitystreams.Group{actualActor}
		expectedGroupMap := expectedActorMap

		BeforeEach(func() {
			expectedGroupMap["type"] = "Group"
		})

		AfterEach(func() {
			delete(expectedGroupMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Group", &actualGroup, expectedGroupMap)
	})

	Describe("Organization", func() {
		actualOrganization := activitystreams.Organization{actualActor}
		expectedOrganizationMap := expectedActorMap

		BeforeEach(func() {
			expectedOrganizationMap["type"] = "Organization"
		})

		AfterEach(func() {
			delete(expectedOrganizationMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Organization", &actualOrganization, expectedOrganizationMap)
	})

	Describe("Person", func() {
		actualPerson := activitystreams.Person{actualActor}
		expectedPersonMap := expectedActorMap

		BeforeEach(func() {
			expectedPersonMap["type"] = "Person"
		})

		AfterEach(func() {
			delete(expectedPersonMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Person", &actualPerson, expectedPersonMap)
	})

	Describe("Service", func() {
		actualService := activitystreams.Service{actualActor}
		expectedServiceMap := expectedActorMap

		BeforeEach(func() {
			expectedServiceMap["type"] = "Service"
		})

		AfterEach(func() {
			delete(expectedServiceMap, "type")
		})

		testutil.CheckActivityStreamsEntity("Service", &actualService, expectedServiceMap)
	})
})
