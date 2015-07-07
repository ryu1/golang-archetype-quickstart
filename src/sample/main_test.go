package main

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestMain(t *testing.T) {

	Describe(t, "fibonacci", func() {

		Context("case argument 20", func() {

			It("result should be 6765", func() {
				Expect(fibonacci(20)).To(Equal, 6765)
			})

		})

		Context("case argument 1", func() {

			It("result should be 1", func() {
				Expect(fibonacci(1)).To(Equal, 1)
			})

		})

	})
}
