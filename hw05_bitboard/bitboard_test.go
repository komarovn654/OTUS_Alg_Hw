package hw05bitboard

import (
	"fmt"
	"testing"
)

func TestKing(t *testing.T) {
	k := Rook{}
	fmt.Println(k.GetMoves(35))
	/*cache := Cache{}
	cache.Init()

	fmt.Println(BitsCountCache(770, &cache))*/
}
