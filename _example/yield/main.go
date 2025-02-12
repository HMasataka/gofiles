package main

import (
	"fmt"

	"github.com/HMasataka/gofiles"
)

func main() {
	iter, err := gofiles.ReadFileYield("example.txt")
	if err != nil {
		panic(err)
	}

	for s := range iter {
		fmt.Println(s)
	}
}
