package main

import "fmt"

func main() {
	fmt.Println(jumlahin(4))
	fmt.Println(fact(7))
}

func jumlahin(nilai int) int {
	if nilai == 0 {
		return 3
	}
	return 2*(jumlahin(nilai-1)) + 4
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
