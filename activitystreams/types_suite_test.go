package activitystreams_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestActivitystreams(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Activitystreams Suite")
}
