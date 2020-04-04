package main

import (
	"fmt"
)

func example(s string) {
	fmt.Println(s)
}

func main() {
	var wg sync.Waitgroup

	go example("foo")
}
