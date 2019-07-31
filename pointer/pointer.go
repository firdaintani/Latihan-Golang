package main

import "fmt"

func main() {
	var first int
	first = 1
	var second *int
	second = &first

	fmt.Println("first :", first)
	fmt.Println("second :", second)
	*second = 2
	fmt.Println("second :", *second)
	fmt.Println("first :", first)
	*second = 3
	fmt.Println("second :", second)
	var angka1 *int
	angkatemp := 1
	angka2 := 2
	angka1 = &angkatemp

	addP1(angka1, angka2)
}

func addP1(angka1 *int, angka2 int) {

	*angka1 += angka2
	fmt.Println(*angka1)
}
