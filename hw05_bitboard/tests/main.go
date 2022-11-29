package main

import "log"

func init() {
	cache.Init()
}

func main() {
	_, err := kingTest("1.Bitboard - Король")
	if err != nil {
		log.Fatal(err)
	}

	_, err = knightTest("2.Bitboard - Конь")
	if err != nil {
		log.Fatal(err)
	}

	_, err = rookTest("3.Bitboard - Ладья")
	if err != nil {
		log.Fatal(err)
	}
}
