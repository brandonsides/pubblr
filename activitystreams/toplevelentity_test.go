package activitystreams_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandonsides/pubblr/activitystreams"
)

var _ = Describe("Toplevelentity", func() {
	actual := activitystreams.TopLevelEntity{
		EntityIface: &activitystreams.Like{
			TransitiveActivity: activitystreams.TransitiveActivity{
				IntransitiveActivity: activitystreams.IntransitiveActivity{
					Object: activitystreams.Object{
						Entity: activitystreams.Entity{
							Id: "http://example.org/like/1",
							AttributedTo: []activitystreams.EntityIface{
								&activitystreams.Person{
									Actor: activitystreams.Actor{
										Object: activitystreams.Object{
											Entity: activitystreams.Entity{
												Id:   "http://sally.example.org",
												Name: "Sally",
											},
										},
									},
								},
							},
						},
						Summary: "Sally liked a repubbed note",
						To: []activitystreams.EntityIface{
							&activitystreams.Person{
								Actor: activitystreams.Actor{
									Object: activitystreams.Object{
										Entity: activitystreams.Entity{
											Id:   "http://joe.example.org",
											Name: "Joe",
										},
									},
								},
							},
						},
						Audience: []activitystreams.EntityIface{
							&activitystreams.Collection{
								Object: activitystreams.Object{
									Entity: activitystreams.Entity{
										Id: "http://sally.example.org/followers",
									},
								},
							},
						},
					},
					Actor: &activitystreams.Person{
						Actor: activitystreams.Actor{
							Object: activitystreams.Object{
								Entity: activitystreams.Entity{
									Id:   "http://sally.example.org",
									Name: "Sally",
								},
							},
						},
					},
				},
				Object: &activitystreams.Announce{
					TransitiveActivity: activitystreams.TransitiveActivity{
						IntransitiveActivity: activitystreams.IntransitiveActivity{
							Object: activitystreams.Object{
								Entity: activitystreams.Entity{
									Id: "http://john.example.org/repubs/1",
								},
							},
						},
						Object: &activitystreams.Note{
							Object: activitystreams.Object{
								Entity: activitystreams.Entity{
									Id: "http://joe.example.org/note/1",
								},
							},
						},
					},
				},
			},
		},
		Context: "https://www.w3.org/ns/activitystreams",
	}
	expected := map[string]interface{}{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type":     "Like",
		"id":       "http://example.org/like/1",
		"actor":    "http://sally.example.org",
		"object":   "http://john.example.org/repubs/1",
		"summary":  "Sally liked a repubbed note",
		"to": []interface{}{
			"http://joe.example.org",
		},
		"audience": []interface{}{
			"http://sally.example.org/followers",
		},
		"attributedTo": []interface{}{
			"http://sally.example.org",
		},
	}

	Describe("MarshalJSON", func() {
		It("should correctly marshal fully populated type", func() {
			jsonObject, err := actual.MarshalJSON()
			Expect(err).ToNot(HaveOccurred())
			var actual map[string]interface{}
			err = json.Unmarshal(jsonObject, &actual)
			Expect(err).ToNot(HaveOccurred())
			for key, value := range expected {
				Expect(actual[key]).To(Equal(value))
			}
			Expect(actual).To(Equal(expected))
		})

		It("should fail to marshal zero value", func() {
			actual := activitystreams.TopLevelEntity{}

			_, err := actual.MarshalJSON()
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("EntityUnmarshaler", func() {
		It("should correctly unmarshal fully populated type", func() {
			jsonObject, err := json.Marshal(expected)
			Expect(err).ToNot(HaveOccurred())

			unmarshalled, err := activitystreams.DefaultEntityUnmarshaler.UnmarshalEntity(jsonObject)
			Expect(err).ToNot(HaveOccurred())
			Expect(unmarshalled).To(Equal(actual))
		})
	})

	Describe("Type", func() {
		It("should return correct type", func() {
			Expect(actual.Type()).To(Equal("Like"))
		})
	})
})
