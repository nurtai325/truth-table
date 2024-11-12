package main

import (
	"flag"
	"log"

	"github.com/chzyer/readline"
	"github.com/nurtai325/truth-table/internal/parser"
)

func main() {
	log.SetFlags(0)
	exp := flag.String("e", "", "specify the `logic expression` to be parsed.")
	flag.Parse()

	if *exp != "" {
		handleExp(exp)
		return
	}

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		handleExp(&line)
	}
}

func handleExp(exp *string) {
	table, err := parser.Parse(exp, false)
	if err != nil {
		log.Fatal(err)
	}
	table.Print()
}
