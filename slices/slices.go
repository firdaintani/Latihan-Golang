package main

import (
	"fmt"
)

func main() {
	
	slices := []string{"dream","catcher"}
	fmt.Println(slices)
	slices=append(slices, "nightmare")
	slices=append(slices, "night")
	slices=append(slices, "mare")
	slices=append(slices, "good", "bye", "day", "night")
	fmt.Println("length slices:",len(slices))
	fmt.Println("cap slices:",cap(slices))
	
	
	
	
	
	fmt.Println(slices)
	slices=append(slices, "night")
	fmt.Println("cap slices after append string night:",cap(slices))
	another_slices := slices[0:2]
	
	fmt.Println(another_slices)
	fmt.Println("len :",len(another_slices))
	fmt.Println("cap:",cap(another_slices))
	fmt.Println("cap slices:",cap(slices))
	
	another_slices[0] = "dread"
	fmt.Println("slice awal : ",slices)
	fmt.Println("another slice :",another_slices)
	
	
	a := make([]int, 5,6)
	
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	c:= make([]string, len(slices))
	copy(c, slices)
	fmt.Println("copy:",c)
	
	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n) 

	
}