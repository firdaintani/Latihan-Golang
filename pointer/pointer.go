package main

import "fmt"

func main() {
	var first int
	first = 1
	var second *int
	second = &first

	fmt.Println("first :", first)
	fmt.Println("second :", second)

}
