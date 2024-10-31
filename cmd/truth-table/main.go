package main

import (
	"flag"
	"fmt"
)

func main() {
	exp := flag.String("e", "", "specify the `logic expression` to be parsed.")
	flag.Parse()
	fmt.Println(*exp)
}
