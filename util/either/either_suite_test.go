package either_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEither(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Either Suite")
}
