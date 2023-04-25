package babble

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("babble", func() {
	var babbler Babbler
	BeforeEach(func() {
		babbler = Babbler{
			Count:     1,
			Words:     []string{"hello"},
			Separator: "☃",
			rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		}
	})

	It("returns a random word", func() {
		Expect(babbler.Babble()).To(Equal("hello"))
	})

	Describe("with multiple words", func() {
		It("concatenates strings", func() {
			babbler.Count = 2
			Expect(babbler.Babble()).To(Equal("hello☃hello"))
		})
	})
})
