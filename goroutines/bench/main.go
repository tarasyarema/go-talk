package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// MaxT defines the max. number of goroutines
const MaxT int = 1000

// Loops is the number of loops will do every goroutine
const Loops int = 1000

type foo struct {
	mutex   sync.Mutex
	counter int
}

func example(f foo, wg sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < Loops; i++ {
	}

	f.mutex.Lock()
	f.counter++
	f.mutex.Unlock()
}

func main() {
	n := MaxT

	if len(os.Args) > 1 {
		tmp, _ := strconv.Atoi(os.Args[1])
		n = tmp
	}

	var wg sync.WaitGroup
	f := foo{sync.Mutex{}, 0}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := 0; i < 1e+3; i++ {
				continue
			}

			f.mutex.Lock()
			f.counter++
			f.mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Printf("> the counter is %d\n", f.counter)
}
