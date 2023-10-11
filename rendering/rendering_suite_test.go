package rendering_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRendering(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rendering Suite")
}
