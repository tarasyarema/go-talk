package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hi, %s!", name)
}

func main() {
	go greet("2pac")
}
