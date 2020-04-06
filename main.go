package main

/*
#include <stdio.h>
extern void lol(int);
extern int fib_c(int);
*/
import "C"
import (
	"fmt"
	"sync"
	"time"
)

//export Fib
func Fib(n int) C.int {
	if n <= 1 {
		return 1
	}

	a, b := 1, 1

	for i := 2; i <= n; i++ {
		tmp := b
		b += a
		a = tmp
	}

	return C.int(b)
}

func main() {
	n := 1000000

	t := time.Now()

	for i := 0; i < n; i++ {
		// C.lol(C.int(rand.Intn(100)))
		C.fib_c(C.int(i))
	}

	elapsed := time.Since(t)
	fmt.Println("a pelo", elapsed)

	var wg sync.WaitGroup

	wg.Add(n)
	t = time.Now()

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			// C.lol(C.int(rand.Intn(100)))
			C.fib_c(C.int(i))
		}(i)
	}

	wg.Wait()
	elapsed = time.Since(t)

	fmt.Println("goroutines", elapsed)
}
