package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	nums := make([][3]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &nums[i][0], &nums[i][1], &nums[i][2])
	}

	substr := ""

	start := time.Now()

	for i := 0; i < n; i++ {
		f, _ := factorial(nums[i][0])

		s := f.String()
		l := len(s)

		begin := l - nums[i][1] + 1
		end := l - nums[i][2]

		substr += s[end:begin]
	}

	elapsed := time.Since(start)

	fmt.Println(substr)
	fmt.Println(elapsed)

	substr = ""
	var wg sync.WaitGroup

	slices := make([]string, n)

	start = time.Now()

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			f, _ := factorial(nums[i][0])

			s := f.String()
			l := len(s)

			begin := l - nums[i][1] + 1
			end := l - nums[i][2]

			slices[i] = s[end:begin]
		}(i)
	}

	wg.Wait()

	elapsed = time.Since(start)

	for i := 0; i < n; i++ {
		fmt.Printf("%s", slices[i])
	}
	fmt.Printf("\n")

	fmt.Println(elapsed)
}
