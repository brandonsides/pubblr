package activitystreams_test

import (
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
				}

				jsonActivity, err := activity.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())
				Expect(string(jsonActivity)).To(Equal(`{"actor":{"id":"http://example.org/john","type":"Person"},"attachment":[{"type":"Image","id":"http://example.org/john/images/1","url":"http://example.org/john/images/1.jpg"},{"@type":"Link","id":"http://example.org/john/images/2","href":"http://example.org/john/images/2.jpg"}],"id":"http://example.org/john/activities/1"}`))
			})
		})
	})
})
