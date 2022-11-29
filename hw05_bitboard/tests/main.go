package main

import "fmt"

func main() {
	tc, err := parseKingTC("1.Bitboard - Король")
	k := hw05bitboard.King{}
	fmt.Println(tc, err)
}
