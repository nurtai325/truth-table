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
	debug := flag.Bool("d", false, "`debug` mode")
	flag.Parse()

	if *exp != "" {
		handleExp(exp, *debug)
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
		handleExp(&line, *debug)
	}
}

func handleExp(exp *string, debug bool) {
	table, err := parser.Parse(exp, debug)
	if err != nil {
		log.Fatal(err)
	}
	table.Print()
}
