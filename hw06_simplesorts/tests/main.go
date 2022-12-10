package main

import (
	"log"
	"time"
)

//go:generate bash ./generate/generate_tests.sh generate/tests.template

var sortTimeout = time.Second * 120

// TODO run tests conccurency
func main() {
	rTable := make(ResultTable)

	for _, r := range runTestsCases(sortTimeout) {
		fillResultTable(r, &rTable)
	}

	if err := saveResultTable("README.md", &rTable); err != nil {
		log.Fatal(err)
	}
}
