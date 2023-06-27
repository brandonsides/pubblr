package activitystreams_test

import (
	"encoding/json"
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandonsides/pubblr/activitystreams"
)

var _ = Describe("Entity", func() {
	actualEntity := activitystreams.Entity{
		Id: "http://example.com/thing",
		AttributedTo: []activitystreams.EntityIface{
			&activitystreams.Person{
				Object: activitystreams.Object{
					Entity: activitystreams.Entity{
						Id: "http://example.com/actor",
					},
				},
			},
		},

		Name:      "thing",
		MediaType: "text/plain",
	}

	Describe("MarshalEntity", func() {
		It("should correctly marshal non-entity embedded type", func() {
			actual := testStruct{
				Entity: activitystreams.Entity{
					Id: "http://example.com/thing",
				},
				TestEmbeddedStruct: TestEmbeddedStruct{
					A: "a",
					B: "b",
				},
				B: "c",
			}
			expectedMap := map[string]interface{}{
				"id":   "http://example.com/thing",
				"A":    "a",
				"B":    "c",
				"type": "testStruct",
			}

			actualJSON, err := activitystreams.MarshalEntity(&actual)
			Expect(err).ToNot(HaveOccurred())
			var actualMap map[string]interface{}
			err = json.Unmarshal(actualJSON, &actualMap)
			Expect(err).ToNot(HaveOccurred())
			Expect(actualMap).To(Equal(expectedMap))
		})
	})

	Describe("TopLevelEntity", func() {
		tle := activitystreams.TopLevelEntity{
			EntityIface: &activitystreams.Object{
				Entity: actualEntity,
			},
			Context: "https://www.w3.org/ns/activitystreams",
		}

		expectedTLEMap := map[string]interface{}{
			"id": "http://example.com/thing",
			"attributedTo": []interface{}{
				map[string]interface{}{
					"id":   "http://example.com/actor",
					"type": "Person",
				},
			},
			"name":      "thing",
			"mediaType": "text/plain",
			"type":      "Object",
			"@context":  "https://www.w3.org/ns/activitystreams",
		}

		Describe("MarshalJSON", func() {
			It("should correctly marshal fully populated type", func() {
				jsonObject, err := tle.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())
				var actual map[string]interface{}
				err = json.Unmarshal(jsonObject, &actual)
				Expect(err).ToNot(HaveOccurred())
				for key, value := range expectedTLEMap {
					Expect(actual[key]).To(Equal(value))
				}
				Expect(actual).To(Equal(expectedTLEMap))
			})

			It("should correctly marshal zero value", func() {
				actual := activitystreams.TopLevelEntity{}

				_, err := actual.MarshalJSON()
				Expect(err.Error()).To(Equal("No EntityIface set on TopLevelEntity"))
			})
		})

		Describe("Type", func() {
			It("should return correct type", func() {
				Expect(tle.Type()).To(Equal("Object"))
			})
		})
	})
})

func CheckActivityStreamsEntity(objectType string, actual activitystreams.EntityIface, expected map[string]interface{}) {
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

		It("should correctly marshal zero value", func() {
			actual := reflect.New(reflect.TypeOf(actual).Elem()).Interface().(activitystreams.EntityIface)

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

type TestEmbeddedStruct struct {
	A string
	B string
}
type testStruct struct {
	activitystreams.Entity
	TestEmbeddedStruct
	B string
}

func (t *testStruct) Type() (string, error) {
	return "testStruct", nil
}

func (t *testStruct) MarshalJSON() ([]byte, error) {
	return activitystreams.MarshalEntity(t)
}
