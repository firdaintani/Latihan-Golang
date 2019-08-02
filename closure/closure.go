package main

import "fmt"

func main() {
	fmt.Println("closure")

	func1 := returnFunc()

	fmt.Println(func1())
	fmt.Println(func1())
	fmt.Println(func1())

	func2 := returnFunc()
	fmt.Println(func2())
	fmt.Println(func2())
	fmt.Println(func2())
	// anotherFunc := func() int {
	// 	return rand.Intn(100)
	// }

	// fmt.Println(anotherFunc())
}

func returnFunc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
