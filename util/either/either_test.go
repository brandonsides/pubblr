package either_test

import (
	"github.com/brandonsides/pubblr/util/either"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Either", func() {
	var e either.Either[int, string]
	Context("Constructed with Left", func() {
		BeforeEach(func() {
			e = *either.Left[int, string](1)
		})

		It("should return the correct Left value", func() {
			Expect(*e.Left()).To(Equal(1))
		})

		It("should return nil for Right", func() {
			Expect(e.Right()).To(BeNil())
		})

		It("should marshal to the correct JSON", func() {
			Expect(e.MarshalJSON()).To(MatchJSON(`1`))
		})
	})

	Context("Constructed with Right", func() {
		BeforeEach(func() {
			e = *either.Right[int]("hello")
		})

		It("should return nil for Left", func() {
			Expect(e.Left()).To(BeNil())
		})

		It("should return the correct Right value", func() {
			Expect(*e.Right()).To(Equal("hello"))
		})

		It("should marshal to the correct JSON", func() {
			Expect(e.MarshalJSON()).To(MatchJSON(`"hello"`))
		})
	})

	Context("Unmarshaling", func() {
		Context("With a valid Left value", func() {
			BeforeEach(func() {
				data := []byte(`1`)
				Expect(e.UnmarshalJSON(data)).To(Succeed())
			})

			It("should return the correct Left value", func() {
				Expect(*e.Left()).To(Equal(1))
			})

			It("should return nil for Right", func() {
				Expect(e.Right()).To(BeNil())
			})
		})

		Context("With a valid Right value", func() {
			BeforeEach(func() {
				data := []byte(`"hello"`)
				Expect(e.UnmarshalJSON(data)).To(Succeed())
			})

			It("should return nil for Left", func() {
				Expect(e.Left()).To(BeNil())
			})

			It("should return the correct Right value", func() {
				Expect(*e.Right()).To(Equal("hello"))
			})
		})

		Context("With an invalid value", func() {
			data := []byte(`{}`)

			It("should return an error", func() {
				Expect(e.UnmarshalJSON(data)).ToNot(Succeed())
			})
		})
	})
})
