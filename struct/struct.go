package main

import "fmt"

func main() {
	fmt.Println("struct")

	type student struct {
		name  string
		grade int
	}
	var s1 student
	s1.name = "john wick"
	s1.grade = 2
	// fmt.Println(s1.name)
	// fmt.Println(s1.grade)
	fmt.Println(s1)
}
