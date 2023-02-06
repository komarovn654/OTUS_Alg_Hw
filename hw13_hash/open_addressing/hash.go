package openaddressing

import (
	"math/rand"
	"time"
)

type hash struct {
	a int
	b int
}

func (h *hash) init() {

}

func getHashCode(k key) int {
	switch kt := k.(type) {
	case string:
		return getStringHashCode(kt)
	case int:
		return getIntegerHashCode(kt)
	default:
		panic("unsupported key type")
	}
}

func getIntegerHashCode(k int) int {
	return 0
}

func getStringHashCode(k string) int {
	return 0
}

func initHash() hash {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	h := hash{}
	h.a = r1.Int()
	h.b = r1.Int()
	return h
}
