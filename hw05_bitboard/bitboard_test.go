package hw05bitboard

import (
	"fmt"
	"testing"
)

func TestKing(t *testing.T) {
	k := Queen{}
	fmt.Println(k.GetMoves(33))
	/*cache := Cache{}
	cache.Init()

	fmt.Println(BitsCountCache(770, &cache))*/
}
