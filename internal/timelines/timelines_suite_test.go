package timelines_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTimelines(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Timelines Suite")
}
