package main

import (
	"log"
	"time"
)

var sortTimeout = time.Second * 120

func main() {
	if err := runTests("README.md", sortTimeout); err != nil {
		log.Fatal(err)
	}
}
