package babble

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBabble(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Babble Suite")
}
