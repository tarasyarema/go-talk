package main

import (
	"fmt"
	"sync"
)

func print(s string) {
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go print("jeje")
	wg.Wait()
}
