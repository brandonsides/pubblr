package util_test

import (
	"github.com/brandonsides/pubblr/activitystreams/util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Either", func() {
	var either util.Either[int, string]
	Context("Constructed with Left", func() {
		BeforeEach(func() {
			either = *util.Left[int, string](1)
		})

		It("should return the correct Left value", func() {
			Expect(*either.Left()).To(Equal(1))
		})

		It("should return nil for Right", func() {
			Expect(either.Right()).To(BeNil())
		})

		It("should marshal to the correct JSON", func() {
			Expect(either.MarshalJSON()).To(MatchJSON(`1`))
		})
	})

	Context("Constructed with Right", func() {
		BeforeEach(func() {
			either = *util.Right[int]("hello")
		})

		It("should return nil for Left", func() {
			Expect(either.Left()).To(BeNil())
		})

		It("should return the correct Right value", func() {
			Expect(*either.Right()).To(Equal("hello"))
		})

		It("should marshal to the correct JSON", func() {
			Expect(either.MarshalJSON()).To(MatchJSON(`"hello"`))
		})
	})

	Context("Unmarshaling", func() {
		Context("With a valid Left value", func() {
			BeforeEach(func() {
				data := []byte(`1`)
				Expect(either.UnmarshalJSON(data)).To(Succeed())
			})

			It("should return the correct Left value", func() {
				Expect(*either.Left()).To(Equal(1))
			})

			It("should return nil for Right", func() {
				Expect(either.Right()).To(BeNil())
			})
		})

		Context("With a valid Right value", func() {
			BeforeEach(func() {
				data := []byte(`"hello"`)
				Expect(either.UnmarshalJSON(data)).To(Succeed())
			})

			It("should return nil for Left", func() {
				Expect(either.Left()).To(BeNil())
			})

			It("should return the correct Right value", func() {
				Expect(*either.Right()).To(Equal("hello"))
			})
		})

		Context("With an invalid value", func() {
			data := []byte(`{}`)

			It("should return an error", func() {
				Expect(either.UnmarshalJSON(data)).ToNot(Succeed())
			})
		})
	})
})
