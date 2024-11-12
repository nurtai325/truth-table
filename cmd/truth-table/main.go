package main

import (
	"flag"
	"log"

	"github.com/nurtai325/truth-table/internal/parser"
)

func main() {
	log.SetFlags(0)
	exp := flag.String("e", "", "specify the `logic expression` to be parsed.")
	flag.Parse()

	if *exp != "" {
		handleExp(exp)
	}
}

func handleExp(exp *string) {
	table, err := parser.Parse(exp, true)
	if err != nil {
		log.Fatal(err)
	}
	table.Print()
}
