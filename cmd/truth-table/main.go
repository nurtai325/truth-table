package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			fmt.Println("error: ", err.Error())
		}
	}()

	panic(fmt.Errorf("hello %s", "world"))
}
