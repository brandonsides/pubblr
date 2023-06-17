package activitystreams_test

import (
	"encoding/json"
	"reflect"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Activity", func() {
	Describe("IntransitiveActivity", func() {
		actualIntransitiveActivity := activitystreams.IntransitiveActivity{
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

		CheckActivityStreamsObject("IntransitiveActivity", &actualIntransitiveActivity, expectedIntransitiveActivityMap)

		Describe("Arrive", func() {
			actualArrive := activitystreams.Arrive{actualIntransitiveActivity}
			expectedArriveMap := expectedIntransitiveActivityMap

			BeforeEach(func() {
				expectedArriveMap["type"] = "Arrive"
			})

			AfterEach(func() {
				expectedArriveMap["type"] = "IntransitiveActivity"
			})

			CheckActivityStreamsObject("Arrive", &actualArrive, expectedArriveMap)
		})

		Describe("Listen", func() {
			actualListen := activitystreams.Listen{actualIntransitiveActivity}
			expectedListenMap := expectedIntransitiveActivityMap

			BeforeEach(func() {
				expectedListenMap["type"] = "Listen"
			})

			AfterEach(func() {
				expectedListenMap["type"] = "IntransitiveActivity"
			})

			CheckActivityStreamsObject("Listen", &actualListen, expectedListenMap)
		})

		Describe("Read", func() {
			actualRead := activitystreams.Read{actualIntransitiveActivity}
			expectedReadMap := expectedIntransitiveActivityMap

			BeforeEach(func() {
				expectedReadMap["type"] = "Read"
			})

			AfterEach(func() {
				expectedReadMap["type"] = "IntransitiveActivity"
			})

			CheckActivityStreamsObject("Read", &actualRead, expectedReadMap)
		})

		Describe("Travel", func() {
			actualTravel := activitystreams.Travel{actualIntransitiveActivity}
			expectedTravelMap := expectedIntransitiveActivityMap

			BeforeEach(func() {
				expectedTravelMap["type"] = "Travel"
			})

			AfterEach(func() {
				expectedTravelMap["type"] = "IntransitiveActivity"
			})

			CheckActivityStreamsObject("Travel", &actualTravel, expectedTravelMap)
		})

		Describe("Question", func() {
			actualQuestion := activitystreams.Question{actualIntransitiveActivity}
			expectedQuestionMap := expectedIntransitiveActivityMap

			BeforeEach(func() {
				expectedQuestionMap["type"] = "Question"
			})

			AfterEach(func() {
				expectedQuestionMap["type"] = "IntransitiveActivity"
			})

			CheckActivityStreamsObject("Question", &actualQuestion, expectedQuestionMap)

			Describe("SingleAnswerQuestion", func() {
				actualSingleAnswerQuestion := activitystreams.SingleAnswerQuestion{
					Question: actualQuestion,
					OneOf: []util.Either[activitystreams.ObjectIface, activitystreams.LinkIface]{
						*util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Object{
							Id:      "http://example.org/john/objects/2",
							Content: "Hello world!",
						}),
						*util.Right[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Link{
							Id:   "http://example.org/john/objects/3",
							Href: "http://example.org/john/objects/3",
						}),
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

				CheckActivityStreamsObject("Question", &actualSingleAnswerQuestion,
					expectedSingleAnswerQuestionMap)
			})

			Describe("MultiAnswerQuestion", func() {
				actualMultiAnswerQuestion := activitystreams.MultiAnswerQuestion{
					Question: actualQuestion,
					AnyOf: []util.Either[activitystreams.ObjectIface, activitystreams.LinkIface]{
						*util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Object{
							Id:      "http://example.org/john/objects/2",
							Content: "Hello world!",
						}),
						*util.Right[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Link{
							Id:   "http://example.org/john/objects/3",
							Href: "http://example.org/john/objects/3",
						}),
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

				CheckActivityStreamsObject("Question", &actualMultiAnswerQuestion,
					expectedMultiAnswerQuestionMap)
			})

			Describe("ClosedQuestion", func() {
				actualClosedQuestion := activitystreams.ClosedQuestion{
					Question: actualQuestion,
					Closed: util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Object{
						Id:      "http://example.org/john/objects/2",
						Content: "Hello world!",
					}),
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

				CheckActivityStreamsObject("Question", &actualClosedQuestion,
					expectedClosedQuestionMap)
			})
		})

		Describe("Activity", func() {
			actualActivity := activitystreams.Activity{
				IntransitiveActivity: actualIntransitiveActivity,
				Object: util.Left[activitystreams.ObjectIface, activitystreams.LinkIface](&activitystreams.Object{
					Id:      "http://example.org/john/objects/2",
					Content: "Hello world!",
				}),
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

			CheckActivityStreamsObject("Activity", &actualActivity, expectedActivityMap)

			Describe("Accept", func() {
				actualAccept := activitystreams.Accept{actualActivity}
				expectedAcceptMap := expectedActivityMap

				BeforeEach(func() {
					expectedAcceptMap["type"] = "Accept"
				})

				AfterEach(func() {
					expectedAcceptMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Accept", &actualAccept, expectedAcceptMap)

				Describe("TentativeAccept", func() {
					actualTentativeAccept := activitystreams.TentativeAccept{actualAccept}
					expectedTentativeAcceptMap := expectedAcceptMap

					BeforeEach(func() {
						expectedTentativeAcceptMap["type"] = "TentativeAccept"
					})

					AfterEach(func() {
						expectedTentativeAcceptMap["type"] = "Accept"
					})

					CheckActivityStreamsObject("TentativeAccept", &actualTentativeAccept, expectedTentativeAcceptMap)
				})
			})

			Describe("Add", func() {
				actualAdd := activitystreams.Add{actualActivity}
				expectedAddMap := expectedActivityMap

				BeforeEach(func() {
					expectedAddMap["type"] = "Add"
				})

				AfterEach(func() {
					expectedAddMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Add", &actualAdd, expectedAddMap)
			})

			Describe("Create", func() {
				actualCreate := activitystreams.Create{actualActivity}
				expectedCreateMap := expectedActivityMap

				BeforeEach(func() {
					expectedCreateMap["type"] = "Create"
				})

				AfterEach(func() {
					expectedCreateMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Create", &actualCreate, expectedCreateMap)
			})

			Describe("Delete", func() {
				actualDelete := activitystreams.Delete{actualActivity}
				expectedDeleteMap := expectedActivityMap

				BeforeEach(func() {
					expectedDeleteMap["type"] = "Delete"
				})

				AfterEach(func() {
					expectedDeleteMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Delete", &actualDelete, expectedDeleteMap)
			})

			Describe("Follow", func() {
				actualFollow := activitystreams.Follow{actualActivity}
				expectedFollowMap := expectedActivityMap

				BeforeEach(func() {
					expectedFollowMap["type"] = "Follow"
				})

				AfterEach(func() {
					expectedFollowMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Follow", &actualFollow, expectedFollowMap)
			})

			Describe("Ignore", func() {
				actualIgnore := activitystreams.Ignore{actualActivity}
				expectedIgnoreMap := expectedActivityMap

				BeforeEach(func() {
					expectedIgnoreMap["type"] = "Ignore"
				})

				AfterEach(func() {
					expectedIgnoreMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Ignore", &actualIgnore, expectedIgnoreMap)

				Describe("Block", func() {
					actualBlock := activitystreams.Block{actualIgnore}
					expectedBlockMap := expectedIgnoreMap

					BeforeEach(func() {
						expectedBlockMap["type"] = "Block"
					})

					AfterEach(func() {
						expectedBlockMap["type"] = "Ignore"
					})

					CheckActivityStreamsObject("Block", &actualBlock, expectedBlockMap)
				})
			})

			Describe("Join", func() {
				actualJoin := activitystreams.Join{actualActivity}
				expectedJoinMap := expectedActivityMap

				BeforeEach(func() {
					expectedJoinMap["type"] = "Join"
				})

				AfterEach(func() {
					expectedJoinMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Join", &actualJoin, expectedJoinMap)
			})

			Describe("Leave", func() {
				actualLeave := activitystreams.Leave{actualActivity}
				expectedLeaveMap := expectedActivityMap

				BeforeEach(func() {
					expectedLeaveMap["type"] = "Leave"
				})

				AfterEach(func() {
					expectedLeaveMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Leave", &actualLeave, expectedLeaveMap)
			})

			Describe("Like", func() {
				actualLike := activitystreams.Like{actualActivity}
				expectedLikeMap := expectedActivityMap

				BeforeEach(func() {
					expectedLikeMap["type"] = "Like"
				})

				AfterEach(func() {
					expectedLikeMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Like", &actualLike, expectedLikeMap)
			})

			Describe("Offer", func() {
				actualOffer := activitystreams.Offer{actualActivity}
				expectedOfferMap := expectedActivityMap

				BeforeEach(func() {
					expectedOfferMap["type"] = "Offer"
				})

				AfterEach(func() {
					expectedOfferMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Offer", &actualOffer, expectedOfferMap)

				Describe("Invite", func() {
					actualInvite := activitystreams.Invite{actualOffer}
					expectedInviteMap := expectedOfferMap

					BeforeEach(func() {
						expectedInviteMap["type"] = "Invite"
					})

					AfterEach(func() {
						expectedInviteMap["type"] = "Offer"
					})

					CheckActivityStreamsObject("Invite", &actualInvite, expectedInviteMap)
				})
			})

			Describe("Reject", func() {
				actualReject := activitystreams.Reject{actualActivity}
				expectedRejectMap := expectedActivityMap

				BeforeEach(func() {
					expectedRejectMap["type"] = "Reject"
				})

				AfterEach(func() {
					expectedRejectMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Reject", &actualReject, expectedRejectMap)

				Describe("TentativeReject", func() {
					actualTentativeReject := activitystreams.TentativeReject{actualReject}
					expectedTentativeRejectMap := expectedRejectMap

					BeforeEach(func() {
						expectedTentativeRejectMap["type"] = "TentativeReject"
					})

					AfterEach(func() {
						expectedTentativeRejectMap["type"] = "Reject"
					})

					CheckActivityStreamsObject("TentativeReject", &actualTentativeReject, expectedTentativeRejectMap)
				})
			})

			Describe("Remove", func() {
				actualRemove := activitystreams.Remove{actualActivity}
				expectedRemoveMap := expectedActivityMap

				BeforeEach(func() {
					expectedRemoveMap["type"] = "Remove"
				})

				AfterEach(func() {
					expectedRemoveMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Remove", &actualRemove, expectedRemoveMap)
			})

			Describe("Undo", func() {
				actualUndo := activitystreams.Undo{actualActivity}
				expectedUndoMap := expectedActivityMap

				BeforeEach(func() {
					expectedUndoMap["type"] = "Undo"
				})

				AfterEach(func() {
					expectedUndoMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Undo", &actualUndo, expectedUndoMap)
			})

			Describe("Update", func() {
				actualUpdate := activitystreams.Update{actualActivity}
				expectedUpdateMap := expectedActivityMap

				BeforeEach(func() {
					expectedUpdateMap["type"] = "Update"
				})

				AfterEach(func() {
					expectedUpdateMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Update", &actualUpdate, expectedUpdateMap)
			})

			Describe("View", func() {
				actualView := activitystreams.View{actualActivity}
				expectedViewMap := expectedActivityMap

				BeforeEach(func() {
					expectedViewMap["type"] = "View"
				})

				AfterEach(func() {
					expectedViewMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("View", &actualView, expectedViewMap)
			})

			Describe("Move", func() {
				actualMove := activitystreams.Move{actualActivity}
				expectedMoveMap := expectedActivityMap

				BeforeEach(func() {
					expectedMoveMap["type"] = "Move"
				})

				AfterEach(func() {
					expectedMoveMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Move", &actualMove, expectedMoveMap)
			})

			Describe("Announce", func() {
				actualAnnounce := activitystreams.Announce{actualActivity}
				expectedAnnounceMap := expectedActivityMap

				BeforeEach(func() {
					expectedAnnounceMap["type"] = "Announce"
				})

				AfterEach(func() {
					expectedAnnounceMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Announce", &actualAnnounce, expectedAnnounceMap)
			})

			Describe("Flag", func() {
				actualFlag := activitystreams.Flag{actualActivity}
				expectedFlagMap := expectedActivityMap

				BeforeEach(func() {
					expectedFlagMap["type"] = "Flag"
				})

				AfterEach(func() {
					expectedFlagMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Flag", &actualFlag, expectedFlagMap)
			})

			Describe("Dislike", func() {
				actualDislike := activitystreams.Dislike{actualActivity}
				expectedDislikeMap := expectedActivityMap

				BeforeEach(func() {
					expectedDislikeMap["type"] = "Dislike"
				})

				AfterEach(func() {
					expectedDislikeMap["type"] = "Activity"
				})

				CheckActivityStreamsObject("Dislike", &actualDislike, expectedDislikeMap)
			})
		})
	})
})

func CheckActivityStreamsObject(objectType string, actual activitystreams.ObjectIface, expected map[string]interface{}) {
	Describe("MarshalJSON", func() {
		It("should correctly marshal fully populated type", func() {
			jsonObject, err := actual.MarshalJSON()
			Expect(err).ToNot(HaveOccurred())
			var actual map[string]interface{}
			err = json.Unmarshal(jsonObject, &actual)
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(Equal(expected))
		})

		It("should correctly marshal zero value", func() {
			actual := reflect.New(reflect.TypeOf(actual).Elem()).Interface().(activitystreams.ObjectIface)

			expected := map[string]interface{}{
				"type": objectType,
			}

			jsonObject, err := actual.MarshalJSON()
			Expect(err).ToNot(HaveOccurred())
			var actualMap map[string]interface{}
			err = json.Unmarshal(jsonObject, &actualMap)
			Expect(err).ToNot(HaveOccurred())
			Expect(actualMap).To(Equal(expected))
		})
	})

	Describe("Type", func() {
		It("should return correct type", func() {
			Expect(actual.Type()).To(Equal(objectType))
		})
	})
}
