package activitystreams_test

import (
	"encoding/json"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Activity", func() {
	Describe("IntransitiveActivity", func() {
		Describe("MarshalJSON", func() {
			It("should correctly marshal fully populated type", func() {
				activity := &activitystreams.IntransitiveActivity{
					Object: activitystreams.Object{
						Id: "http://example.org/john/activities/1",
						Attachment: []util.Either[activitystreams.ObjectIface, activitystreams.LinkIface]{
							*util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Image{
								Object: activitystreams.Object{
									Id:  "http://example.org/john/images/1",
									URL: util.Left[string, activitystreams.LinkIface]("http://example.org/john/images/1.jpg"),
								},
							}),
							*util.Right[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Link{
								Id:   "http://example.org/john/images/2",
								Href: "http://example.org/john/images/2.jpg",
							}),
						},
					},
					Actor: util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Person{
						Object: activitystreams.Object{
							Id: "http://example.org/john",
						},
					}),
					Target: util.Right[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Link{
						Id:   "http://example.org/john/objects/1",
						Href: "http://example.org/john/objects/1",
					}),
					Result: util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Object{
						Id: "http://example.org/john/activities/1/result",
					}),
					Origin: util.Right[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Link{
						Id:   "http://example.org/john/activities/1/origin",
						Href: "http://example.org/john/activities/1/origin",
					}),
					Instrument: util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Object{
						Id: "http://example.org/john/activities/1/instrument",
					}),
				}

				expected := map[string]interface{}{
					"type": "IntransitiveActivity",
					"actor": map[string]interface{}{
						"type": "Person",
						"id":   "http://example.org/john",
					},
					"attachment": []interface{}{
						map[string]interface{}{
							"type": "Image",
							"id":   "http://example.org/john/images/1",
							"url":  "http://example.org/john/images/1.jpg",
						},
						map[string]interface{}{
							"type": "Link",
							"id":   "http://example.org/john/images/2",
							"href": "http://example.org/john/images/2.jpg",
						},
					},
					"id": "http://example.org/john/activities/1",
					"instrument": map[string]interface{}{
						"id":   "http://example.org/john/activities/1/instrument",
						"type": "Object",
					},
					"origin": map[string]interface{}{
						"href": "http://example.org/john/activities/1/origin",
						"id":   "http://example.org/john/activities/1/origin",
						"type": "Link",
					},
					"result": map[string]interface{}{
						"id":   "http://example.org/john/activities/1/result",
						"type": "Object",
					},
					"target": map[string]interface{}{
						"href": "http://example.org/john/objects/1",
						"id":   "http://example.org/john/objects/1",
						"type": "Link",
					},
				}

				jsonActivity, err := activity.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())
				var actual map[string]interface{}
				err = json.Unmarshal(jsonActivity, &actual)
				Expect(err).ToNot(HaveOccurred())
				Expect(actual).To(Equal(expected))
			})

			It("should correctly marshal zero value", func() {
				activity := &activitystreams.IntransitiveActivity{}

				expected := map[string]interface{}{
					"type": "IntransitiveActivity",
				}

				jsonActivity, err := activity.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())
				var actual map[string]interface{}
				err = json.Unmarshal(jsonActivity, &actual)
				Expect(err).ToNot(HaveOccurred())
				Expect(actual).To(Equal(expected))
			})
		})

		Describe("Type", func() {
			It("should return correct type", func() {
				activity := &activitystreams.IntransitiveActivity{}
				Expect(activity.Type()).To(Equal("IntransitiveActivity"))
			})
		})
	})

	Describe("Activity", func() {
		Describe("MarshalJSON", func() {
			It("should correctly marshal fully populated type", func() {
				activity := &activitystreams.Activity{
					IntransitiveActivity: activitystreams.IntransitiveActivity{
						Object: activitystreams.Object{
							Id: "http://example.org/john/activities/1",
							Attachment: []util.Either[activitystreams.ObjectIface, activitystreams.LinkIface]{
								*util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Image{
									Object: activitystreams.Object{
										Id:  "http://example.org/john/images/1",
										URL: util.Left[string, activitystreams.LinkIface]("http://example.org/john/images/1.jpg"),
									},
								}),
								*util.Right[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Link{
									Id:   "http://example.org/john/images/2",
									Href: "http://example.org/john/images/2.jpg",
								}),
							},
						},
						Actor: util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Person{
							Object: activitystreams.Object{
								Id: "http://example.org/john",
							},
						}),
						Target: util.Right[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Link{
							Id:   "http://example.org/john/objects/1",
							Href: "http://example.org/john/objects/1",
						}),
						Result: util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Object{
							Id: "http://example.org/john/activities/1/result",
						}),
						Origin: util.Right[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Link{
							Id:   "http://example.org/john/activities/1/origin",
							Href: "http://example.org/john/activities/1/origin",
						}),
						Instrument: util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Object{
							Id: "http://example.org/john/activities/1/instrument",
						}),
					},
					Object: util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Object{
						Id:      "http://example.org/john/objects/2",
						Content: "Hello world!",
					}),
				}

				expected := map[string]interface{}{
					"type": "Activity",
					"actor": map[string]interface{}{
						"type": "Person",
						"id":   "http://example.org/john",
					},
					"attachment": []interface{}{
						map[string]interface{}{
							"type": "Image",
							"id":   "http://example.org/john/images/1",
							"url":  "http://example.org/john/images/1.jpg",
						},
						map[string]interface{}{
							"type": "Link",
							"id":   "http://example.org/john/images/2",
							"href": "http://example.org/john/images/2.jpg",
						},
					},
					"id": "http://example.org/john/activities/1",
					"instrument": map[string]interface{}{
						"id":   "http://example.org/john/activities/1/instrument",
						"type": "Object",
					},
					"origin": map[string]interface{}{
						"href": "http://example.org/john/activities/1/origin",
						"id":   "http://example.org/john/activities/1/origin",
						"type": "Link",
					},
					"result": map[string]interface{}{
						"id":   "http://example.org/john/activities/1/result",
						"type": "Object",
					},
					"target": map[string]interface{}{
						"href": "http://example.org/john/objects/1",
						"id":   "http://example.org/john/objects/1",
						"type": "Link",
					},
					"object": map[string]interface{}{
						"content": "Hello world!",
						"id":      "http://example.org/john/objects/2",
						"type":    "Object",
					},
				}

				jsonActivity, err := activity.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())
				var actual map[string]interface{}
				err = json.Unmarshal(jsonActivity, &actual)
				Expect(err).ToNot(HaveOccurred())
				Expect(actual).To(Equal(expected))
			})

			It("should correctly marshal zero value", func() {
				activity := &activitystreams.IntransitiveActivity{}

				expected := map[string]interface{}{
					"type": "IntransitiveActivity",
				}

				jsonActivity, err := activity.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())
				var actual map[string]interface{}
				err = json.Unmarshal(jsonActivity, &actual)
				Expect(err).ToNot(HaveOccurred())
				Expect(actual).To(Equal(expected))
			})
		})

		Describe("Type", func() {
			It("should return correct type", func() {
				activity := &activitystreams.IntransitiveActivity{}
				Expect(activity.Type()).To(Equal("IntransitiveActivity"))
			})
		})
	})
})
