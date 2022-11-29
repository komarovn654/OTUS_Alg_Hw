package hw05bitboard

import (
	"fmt"
	"testing"
)

func TestKing(t *testing.T) {
	k := King{}
	fmt.Println(k.GetMoves(24))
}
