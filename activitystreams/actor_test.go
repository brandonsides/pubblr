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
	actualActor := activitystreams.Actor{
		Object: activitystreams.Object{
			Entity: activitystreams.Entity{
				Id: "http://example.org/~john",
				AttributedTo: []activitystreams.EntityIface{
					&activitystreams.Person{
						Actor: activitystreams.Actor{
							Object: activitystreams.Object{
								Entity: activitystreams.Entity{
									Id: "http://example.org/~john",
								},
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
					Actor: activitystreams.Actor{
						Object: activitystreams.Object{
							Entity: activitystreams.Entity{
								Id: "http://example.org/~alice",
							},
						},
					},
				},
			},
			Bto: []activitystreams.EntityIface{
				&activitystreams.Person{
					Actor: activitystreams.Actor{
						Object: activitystreams.Object{
							Entity: activitystreams.Entity{
								Id: "http://example.org/~bob",
							},
						},
					},
				},
			},
			Cc: []activitystreams.EntityIface{
				&activitystreams.Person{
					Actor: activitystreams.Actor{
						Object: activitystreams.Object{
							Entity: activitystreams.Entity{
								Id: "http://example.org/~eve",
							},
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
						Id:   "http://example.org/~john/avatar",
					},
					URL: either.Left[string, activitystreams.LinkIface]("http://example.org/~john/avatar.jpg"),
				},
			},
			Image: &activitystreams.Image{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id:   "http://example.org/~john/header",
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
							Id:   "http://example.org/~john/location",
						},
					},
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
					Actor: activitystreams.Actor{
						Object: activitystreams.Object{
							Entity: activitystreams.Entity{
								Id: "http://example.org/~alice",
							},
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
		},
	}
	expectedActorMap := map[string]interface{}{
		"id": "http://example.org/~john",
		"attachment": []interface{}{
			"http://example.org/~john/picture",
			"http://example.org/~john/profile",
		},
		"attributedTo": []interface{}{
			"http://example.org/~john",
		},
		"bcc": []interface{}{
			"http://example.org/~alice",
		},
		"bto": []interface{}{
			"http://example.org/~bob",
		},
		"cc": []interface{}{
			"http://example.org/~eve",
		},
		"context":   "http://example.org/contexts/1",
		"generator": "http://example.org/generator",
		"icon":      "http://example.org/~john/avatar",
		"image":     "http://example.org/~john/header",
		"inReplyTo": []interface{}{
			"http://example.org/posts/1",
		},
		"location": []interface{}{
			"http://example.org/~john/location",
		},
		"preview": "http://example.org/~john/preview",
		"replies": "http://example.org/~john/replies",
		"tag": []interface{}{
			"http://example.org/tags/1",
		},
		"to": []interface{}{
			"http://example.org/~alice",
		},
		"url":       "http://example.org/~john",
		"content":   "This is a simple note",
		"name":      "A Simple Note",
		"duration":  duration.String(),
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
