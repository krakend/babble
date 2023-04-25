package babble

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Babbler struct {
	Count     int
	Separator string
	Words     []string
	mu        sync.Mutex
	rand      *rand.Rand
}

func NewBabbler() Babbler {
	return NewBabblerWithRand(rand.New(rand.NewSource(time.Now().UnixNano())))
}

func NewBabblerWithRand(rnd *rand.Rand) Babbler {
	return NewBabblerWithDictionary(rnd, readAvailableDictionary())
}

func NewBabblerWithDictionary(rnd *rand.Rand, dict []string) Babbler {
	return Babbler{
		Count:     2,
		Separator: "-",
		Words:     dict,
		rand:      rnd,
	}
}

func (b *Babbler) Babble() string {
	pieces := []string{}
	b.mu.Lock()
	for i := 0; i < b.Count; i++ {
		pieces = append(pieces, b.Words[b.rand.Int()%len(b.Words)])
	}
	b.mu.Unlock()
	return strings.Join(pieces, b.Separator)
}
