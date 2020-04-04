package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	for i := 0; i < n; i++ {
		fmt.Println(i, nums[i])
	}
}
