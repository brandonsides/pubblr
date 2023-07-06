package activitystreams_test

import (
	. "github.com/onsi/ginkgo/v2"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/activitystreams/testutil"
)

var _ = Describe("Toplevelentity", func() {
	topLevelEntity := activitystreams.TopLevelEntity{
		EntityIface: &activitystreams.Like{
			Activity: activitystreams.Activity{
				IntransitiveActivity: activitystreams.IntransitiveActivity{
					Object: activitystreams.Object{
						Entity: activitystreams.Entity{
							Id: "http://example.org/like/1",
							AttributedTo: []activitystreams.EntityIface{
								&activitystreams.Person{
									Object: activitystreams.Object{
										Entity: activitystreams.Entity{
											Id:   "http://sally.example.org",
											Name: "Sally",
										},
									},
								},
							},
						},
						Summary: "Sally liked a repubbed note",
						To: []activitystreams.EntityIface{
							&activitystreams.Person{
								Object: activitystreams.Object{
									Entity: activitystreams.Entity{
										Id:   "http://joe.example.org",
										Name: "Joe",
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
						Object: activitystreams.Object{
							Entity: activitystreams.Entity{
								Id:   "http://sally.example.org",
								Name: "Sally",
							},
						},
					},
				},
				Object: &activitystreams.Announce{
					Activity: activitystreams.Activity{
						IntransitiveActivity: activitystreams.IntransitiveActivity{
							Object: activitystreams.Object{
								Entity: activitystreams.Entity{
									Id: "http://sally.example.org/repubs/1",
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
	expectedTopLevelEntityMap := map[string]interface{}{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type":     "Like",
		"id":       "http://example.org/like/1",
		"actor": map[string]interface{}{
			"type": "Person",
			"id":   "http://sally.example.org",
			"name": "Sally",
		},
		"object": map[string]interface{}{
			"type": "Announce",
			"id":   "http://sally.example.org/repubs/1",
			"object": map[string]interface{}{
				"type": "Note",
				"id":   "http://joe.example.org/note/1",
			},
		},
		"summary": "Sally liked a repubbed note",
		"to": []interface{}{
			map[string]interface{}{
				"type": "Person",
				"id":   "http://joe.example.org",
				"name": "Joe",
			},
		},
		"audience": []interface{}{
			map[string]interface{}{
				"type": "Collection",
				"id":   "http://sally.example.org/followers",
			},
		},
		"attributedTo": []interface{}{
			map[string]interface{}{
				"type": "Person",
				"id":   "http://sally.example.org",
				"name": "Sally",
			},
		},
	}

	testutil.CheckActivityStreamsEntity("Like", &topLevelEntity, expectedTopLevelEntityMap)
})
