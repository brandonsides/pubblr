package testutil

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/activitystreams/entity"
	"github.com/go-test/deep"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func CheckActivityStreamsEntity(objectType string, actual entity.EntityIface, expected map[string]interface{}) {
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
			actual := reflect.New(reflect.TypeOf(actual).Elem()).Interface().(entity.EntityIface)

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

	Describe("EntityUnmarshaler", func() {
		It("should correctly unmarshal fully populated type", func() {
			jsonObject, err := json.Marshal(expected)
			Expect(err).ToNot(HaveOccurred())

			unmarshalled, err := activitystreams.DefaultEntityUnmarshaler.UnmarshalInterface(jsonObject)
			Expect(err).ToNot(HaveOccurred())
			diff := deep.Equal(unmarshalled, actual)
			if diff != nil {
				fmt.Println(diff)
			}
			Expect(unmarshalled).To(Equal(actual))
		})
	})

	Describe("Type", func() {
		It("should return correct type", func() {
			Expect(actual.Type()).To(Equal(objectType))
		})
	})
}
