package main

import (
	"fmt"
)

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0], fruits[1], fruits[2], fruits[3])

	var num [3]int

	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
		num[i] = i
		fmt.Println(num[i])

	}

	var number = [2]int{1, 2}
	looping(number)

	stringskata := [2]string{"halo", "world"}
	fmt.Println(stringskata)

	names := [...]string{
		"kiki",
		"dita",
		"angga",
	}
	fmt.Println(names)

	date := [4][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{3, 4},
		[2]int{3, 4},
	}
	fmt.Println(date)

	dates := [...][2]int{
		[...]int{1, 2},
		[...]int{3, 4},
		[...]int{3, 4},
	}
	fmt.Println(dates)

}

func looping(arr [2]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])

	}
}
