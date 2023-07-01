package types_test

import (
	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/activitystreams/testutil"
	"github.com/brandonsides/pubblr/activitystreams/types"
	"github.com/brandonsides/pubblr/util/either"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Activity", func() {
	Describe("IntransitiveActivity", func() {
		actualIntransitiveActivity := types.IntransitiveActivity{
			Object: types.Object{
				Entity: activitystreams.Entity{
					Id: "http://example.org/john/activities/1",
				},
				Attachment: []activitystreams.EntityIface{
					&types.Image{
						Object: types.Object{
							Entity: activitystreams.Entity{
								Id: "http://example.org/john/images/1",
							},
							URL: either.Left[string, types.LinkIface]("http://example.org/john/images/1.jpg"),
						},
					},
					&types.Link{
						Entity: activitystreams.Entity{
							Id: "http://example.org/john/images/2",
						},
						Href: "http://example.org/john/images/2.jpg",
					},
				},
			},
			Actor: &types.Person{
				Object: types.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/john",
					},
				},
			},
			Target: &types.Link{
				Entity: activitystreams.Entity{
					Id: "http://example.org/john/objects/1",
				},
				Href: "http://example.org/john/objects/1",
			},
			Result: &types.Object{
				Entity: activitystreams.Entity{
					Id: "http://example.org/john/activities/1/result",
				},
			},
			Origin: &types.Link{
				Entity: activitystreams.Entity{
					Id: "http://example.org/john/activities/1/origin",
				},
				Href: "http://example.org/john/activities/1/origin",
			},
			Instrument: &types.Object{
				Entity: activitystreams.Entity{
					Id: "http://example.org/john/activities/1/instrument",
				},
			},
		}
		expectedIntransitiveActivityMap := map[string]interface{}{
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

		testutil.CheckActivityStreamsEntity("IntransitiveActivity", &actualIntransitiveActivity, expectedIntransitiveActivityMap)

		Describe("Arrive", func() {
			actualArrive := types.Arrive{actualIntransitiveActivity}
			expectedArriveMap := expectedIntransitiveActivityMap

			BeforeEach(func() {
				expectedArriveMap["type"] = "Arrive"
			})

			AfterEach(func() {
				expectedArriveMap["type"] = "IntransitiveActivity"
			})

			testutil.CheckActivityStreamsEntity("Arrive", &actualArrive, expectedArriveMap)
		})

		Describe("Listen", func() {
			actualListen := types.Listen{actualIntransitiveActivity}
			expectedListenMap := expectedIntransitiveActivityMap

			BeforeEach(func() {
				expectedListenMap["type"] = "Listen"
			})

			AfterEach(func() {
				expectedListenMap["type"] = "IntransitiveActivity"
			})

			testutil.CheckActivityStreamsEntity("Listen", &actualListen, expectedListenMap)
		})

		Describe("Read", func() {
			actualRead := types.Read{actualIntransitiveActivity}
			expectedReadMap := expectedIntransitiveActivityMap

			BeforeEach(func() {
				expectedReadMap["type"] = "Read"
			})

			AfterEach(func() {
				expectedReadMap["type"] = "IntransitiveActivity"
			})

			testutil.CheckActivityStreamsEntity("Read", &actualRead, expectedReadMap)
		})

		Describe("Travel", func() {
			actualTravel := types.Travel{actualIntransitiveActivity}
			expectedTravelMap := expectedIntransitiveActivityMap

			BeforeEach(func() {
				expectedTravelMap["type"] = "Travel"
			})

			AfterEach(func() {
				expectedTravelMap["type"] = "IntransitiveActivity"
			})

			testutil.CheckActivityStreamsEntity("Travel", &actualTravel, expectedTravelMap)
		})

		Describe("Question", func() {
			actualQuestion := types.Question{actualIntransitiveActivity}
			expectedQuestionMap := expectedIntransitiveActivityMap

			BeforeEach(func() {
				expectedQuestionMap["type"] = "Question"
			})

			AfterEach(func() {
				expectedQuestionMap["type"] = "IntransitiveActivity"
			})

			testutil.CheckActivityStreamsEntity("Question", &actualQuestion, expectedQuestionMap)

			Describe("SingleAnswerQuestion", func() {
				actualSingleAnswerQuestion := types.SingleAnswerQuestion{
					Question: actualQuestion,
					OneOf: []activitystreams.EntityIface{
						&types.Object{
							Entity: activitystreams.Entity{
								Id: "http://example.org/john/objects/2",
							},
							Content: "Hello world!",
						},
						&types.Link{
							Entity: activitystreams.Entity{
								Id: "http://example.org/john/objects/3",
							},
							Href: "http://example.org/john/objects/3",
						},
					},
				}
				expectedSingleAnswerQuestionMap := expectedQuestionMap

				BeforeEach(func() {
					expectedSingleAnswerQuestionMap["oneOf"] = []interface{}{
						map[string]interface{}{
							"type":    "Object",
							"id":      "http://example.org/john/objects/2",
							"content": "Hello world!",
						},
						map[string]interface{}{
							"type": "Link",
							"id":   "http://example.org/john/objects/3",
							"href": "http://example.org/john/objects/3",
						},
					}
				})

				AfterEach(func() {
					delete(expectedSingleAnswerQuestionMap, "oneOf")
				})

				testutil.CheckActivityStreamsEntity("Question", &actualSingleAnswerQuestion,
					expectedSingleAnswerQuestionMap)
			})

			Describe("MultiAnswerQuestion", func() {
				actualMultiAnswerQuestion := types.MultiAnswerQuestion{
					Question: actualQuestion,
					AnyOf: []activitystreams.EntityIface{
						&types.Object{
							Entity: activitystreams.Entity{
								Id: "http://example.org/john/objects/2",
							},
							Content: "Hello world!",
						},
						&types.Link{
							Entity: activitystreams.Entity{
								Id: "http://example.org/john/objects/3",
							},
							Href: "http://example.org/john/objects/3",
						},
					},
				}
				expectedMultiAnswerQuestionMap := expectedQuestionMap

				BeforeEach(func() {
					expectedMultiAnswerQuestionMap["anyOf"] = []interface{}{
						map[string]interface{}{
							"type":    "Object",
							"id":      "http://example.org/john/objects/2",
							"content": "Hello world!",
						},
						map[string]interface{}{
							"type": "Link",
							"id":   "http://example.org/john/objects/3",
							"href": "http://example.org/john/objects/3",
						},
					}
				})

				AfterEach(func() {
					delete(expectedMultiAnswerQuestionMap, "anyOf")
				})

				testutil.CheckActivityStreamsEntity("Question", &actualMultiAnswerQuestion,
					expectedMultiAnswerQuestionMap)
			})

			Describe("ClosedQuestion", func() {
				actualClosedQuestion := types.ClosedQuestion{
					Question: actualQuestion,
					Closed: &types.Object{
						Entity: activitystreams.Entity{
							Id: "http://example.org/john/objects/2",
						},
						Content: "Hello world!",
					},
				}
				expectedClosedQuestionMap := expectedQuestionMap

				BeforeEach(func() {
					expectedClosedQuestionMap["closed"] = map[string]interface{}{
						"type":    "Object",
						"id":      "http://example.org/john/objects/2",
						"content": "Hello world!",
					}
				})

				AfterEach(func() {
					delete(expectedClosedQuestionMap, "closed")
				})

				testutil.CheckActivityStreamsEntity("Question", &actualClosedQuestion,
					expectedClosedQuestionMap)
			})
		})

		Describe("Activity", func() {
			actualActivity := types.Activity{
				IntransitiveActivity: actualIntransitiveActivity,
				Object: &types.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.org/john/objects/2",
					},
					Content: "Hello world!",
				},
			}
			expectedActivityMap := map[string]interface{}{
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

			testutil.CheckActivityStreamsEntity("Activity", &actualActivity, expectedActivityMap)

			Describe("Accept", func() {
				actualAccept := types.Accept{actualActivity}
				expectedAcceptMap := expectedActivityMap

				BeforeEach(func() {
					expectedAcceptMap["type"] = "Accept"
				})

				AfterEach(func() {
					expectedAcceptMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Accept", &actualAccept, expectedAcceptMap)

				Describe("TentativeAccept", func() {
					actualTentativeAccept := types.TentativeAccept{actualAccept}
					expectedTentativeAcceptMap := expectedAcceptMap

					BeforeEach(func() {
						expectedTentativeAcceptMap["type"] = "TentativeAccept"
					})

					AfterEach(func() {
						expectedTentativeAcceptMap["type"] = "Accept"
					})

					testutil.CheckActivityStreamsEntity("TentativeAccept", &actualTentativeAccept, expectedTentativeAcceptMap)
				})
			})

			Describe("Add", func() {
				actualAdd := types.Add{actualActivity}
				expectedAddMap := expectedActivityMap

				BeforeEach(func() {
					expectedAddMap["type"] = "Add"
				})

				AfterEach(func() {
					expectedAddMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Add", &actualAdd, expectedAddMap)
			})

			Describe("Create", func() {
				actualCreate := types.Create{actualActivity}
				expectedCreateMap := expectedActivityMap

				BeforeEach(func() {
					expectedCreateMap["type"] = "Create"
				})

				AfterEach(func() {
					expectedCreateMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Create", &actualCreate, expectedCreateMap)
			})

			Describe("Delete", func() {
				actualDelete := types.Delete{actualActivity}
				expectedDeleteMap := expectedActivityMap

				BeforeEach(func() {
					expectedDeleteMap["type"] = "Delete"
				})

				AfterEach(func() {
					expectedDeleteMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Delete", &actualDelete, expectedDeleteMap)
			})

			Describe("Follow", func() {
				actualFollow := types.Follow{actualActivity}
				expectedFollowMap := expectedActivityMap

				BeforeEach(func() {
					expectedFollowMap["type"] = "Follow"
				})

				AfterEach(func() {
					expectedFollowMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Follow", &actualFollow, expectedFollowMap)
			})

			Describe("Ignore", func() {
				actualIgnore := types.Ignore{actualActivity}
				expectedIgnoreMap := expectedActivityMap

				BeforeEach(func() {
					expectedIgnoreMap["type"] = "Ignore"
				})

				AfterEach(func() {
					expectedIgnoreMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Ignore", &actualIgnore, expectedIgnoreMap)

				Describe("Block", func() {
					actualBlock := types.Block{actualIgnore}
					expectedBlockMap := expectedIgnoreMap

					BeforeEach(func() {
						expectedBlockMap["type"] = "Block"
					})

					AfterEach(func() {
						expectedBlockMap["type"] = "Ignore"
					})

					testutil.CheckActivityStreamsEntity("Block", &actualBlock, expectedBlockMap)
				})
			})

			Describe("Join", func() {
				actualJoin := types.Join{actualActivity}
				expectedJoinMap := expectedActivityMap

				BeforeEach(func() {
					expectedJoinMap["type"] = "Join"
				})

				AfterEach(func() {
					expectedJoinMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Join", &actualJoin, expectedJoinMap)
			})

			Describe("Leave", func() {
				actualLeave := types.Leave{actualActivity}
				expectedLeaveMap := expectedActivityMap

				BeforeEach(func() {
					expectedLeaveMap["type"] = "Leave"
				})

				AfterEach(func() {
					expectedLeaveMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Leave", &actualLeave, expectedLeaveMap)
			})

			Describe("Like", func() {
				actualLike := types.Like{actualActivity}
				expectedLikeMap := expectedActivityMap

				BeforeEach(func() {
					expectedLikeMap["type"] = "Like"
				})

				AfterEach(func() {
					expectedLikeMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Like", &actualLike, expectedLikeMap)
			})

			Describe("Offer", func() {
				actualOffer := types.Offer{actualActivity}
				expectedOfferMap := expectedActivityMap

				BeforeEach(func() {
					expectedOfferMap["type"] = "Offer"
				})

				AfterEach(func() {
					expectedOfferMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Offer", &actualOffer, expectedOfferMap)

				Describe("Invite", func() {
					actualInvite := types.Invite{actualOffer}
					expectedInviteMap := expectedOfferMap

					BeforeEach(func() {
						expectedInviteMap["type"] = "Invite"
					})

					AfterEach(func() {
						expectedInviteMap["type"] = "Offer"
					})

					testutil.CheckActivityStreamsEntity("Invite", &actualInvite, expectedInviteMap)
				})
			})

			Describe("Reject", func() {
				actualReject := types.Reject{actualActivity}
				expectedRejectMap := expectedActivityMap

				BeforeEach(func() {
					expectedRejectMap["type"] = "Reject"
				})

				AfterEach(func() {
					expectedRejectMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Reject", &actualReject, expectedRejectMap)

				Describe("TentativeReject", func() {
					actualTentativeReject := types.TentativeReject{actualReject}
					expectedTentativeRejectMap := expectedRejectMap

					BeforeEach(func() {
						expectedTentativeRejectMap["type"] = "TentativeReject"
					})

					AfterEach(func() {
						expectedTentativeRejectMap["type"] = "Reject"
					})

					testutil.CheckActivityStreamsEntity("TentativeReject", &actualTentativeReject, expectedTentativeRejectMap)
				})
			})

			Describe("Remove", func() {
				actualRemove := types.Remove{actualActivity}
				expectedRemoveMap := expectedActivityMap

				BeforeEach(func() {
					expectedRemoveMap["type"] = "Remove"
				})

				AfterEach(func() {
					expectedRemoveMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Remove", &actualRemove, expectedRemoveMap)
			})

			Describe("Undo", func() {
				actualUndo := types.Undo{actualActivity}
				expectedUndoMap := expectedActivityMap

				BeforeEach(func() {
					expectedUndoMap["type"] = "Undo"
				})

				AfterEach(func() {
					expectedUndoMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Undo", &actualUndo, expectedUndoMap)
			})

			Describe("Update", func() {
				actualUpdate := types.Update{actualActivity}
				expectedUpdateMap := expectedActivityMap

				BeforeEach(func() {
					expectedUpdateMap["type"] = "Update"
				})

				AfterEach(func() {
					expectedUpdateMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Update", &actualUpdate, expectedUpdateMap)
			})

			Describe("View", func() {
				actualView := types.View{actualActivity}
				expectedViewMap := expectedActivityMap

				BeforeEach(func() {
					expectedViewMap["type"] = "View"
				})

				AfterEach(func() {
					expectedViewMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("View", &actualView, expectedViewMap)
			})

			Describe("Move", func() {
				actualMove := types.Move{actualActivity}
				expectedMoveMap := expectedActivityMap

				BeforeEach(func() {
					expectedMoveMap["type"] = "Move"
				})

				AfterEach(func() {
					expectedMoveMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Move", &actualMove, expectedMoveMap)
			})

			Describe("Announce", func() {
				actualAnnounce := types.Announce{actualActivity}
				expectedAnnounceMap := expectedActivityMap

				BeforeEach(func() {
					expectedAnnounceMap["type"] = "Announce"
				})

				AfterEach(func() {
					expectedAnnounceMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Announce", &actualAnnounce, expectedAnnounceMap)
			})

			Describe("Flag", func() {
				actualFlag := types.Flag{actualActivity}
				expectedFlagMap := expectedActivityMap

				BeforeEach(func() {
					expectedFlagMap["type"] = "Flag"
				})

				AfterEach(func() {
					expectedFlagMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Flag", &actualFlag, expectedFlagMap)
			})

			Describe("Dislike", func() {
				actualDislike := types.Dislike{actualActivity}
				expectedDislikeMap := expectedActivityMap

				BeforeEach(func() {
					expectedDislikeMap["type"] = "Dislike"
				})

				AfterEach(func() {
					expectedDislikeMap["type"] = "Activity"
				})

				testutil.CheckActivityStreamsEntity("Dislike", &actualDislike, expectedDislikeMap)
			})
		})
	})
})
