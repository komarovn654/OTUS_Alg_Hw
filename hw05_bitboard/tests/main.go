package main

import (
	"log"

	"github.com/komarovn654/OTUS_Alg_Hw/hw05bitboard"
)

var (
	cache = hw05bitboard.Cache{}
)

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
