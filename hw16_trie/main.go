package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	trie_map "github.com/komarovn654/OTUS_Alg_Hw/hw16_trie/map"
)

const (
	charSet = "abcdefghijklmnopqrstuvwxyz"
)

var (
	stringSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GenerateRandomString(length int, charset string) string {
	b := make([]byte, length)
	stringSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = charset[stringSeed.Intn(len(charset)-1)]
	}
	return string(b)
}

func GenerateRandomItems(len int) []trie_map.Item {
	res := make([]trie_map.Item, len)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range res {
		res[i].Key = GenerateRandomString(r.Intn(10)+1, charSet)
		res[i].Value = r.Intn(500)
	}

	return res
}

func main() {
	goMap := make(map[string]any)
	trieMap := trie_map.Constructor()
	items := GenerateRandomItems(100_000)

	tmr := time.Now()
	for _, i := range items {
		trieMap.Insert(i)
	}
	for _, i := range items {
		if _, ok := trieMap.Search(i.Key); !ok {
			log.Fatalf("trieMap: key doesnt exist")
		}
	}
	fmt.Println(time.Since(tmr))

	tmr = time.Now()
	for _, i := range items {
		goMap[i.Key] = i.Value
	}
	for _, i := range items {
		if _, ok := goMap[i.Key]; !ok {
			log.Fatalf("goMap: key doesnt exist")
		}
	}
	fmt.Println(time.Since(tmr))
}