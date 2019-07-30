package main

import "fmt"

func main(){
	m := make(map[string]int)

	fmt.Println("map:", m)
	
	m["sunday"]= 1
	m["monday"]=2
	m["tuesday"]=3

	fmt.Println(m["sunday"])
	fmt.Println("map length:", len(m))
	fmt.Println(m["friday"])
	

	day:= map[string]string{
		"today" :"tuesday",
		"tomorrow" : "wednesday",
		"yesterday" : "monday",
	}
	fmt.Println(day["today"])
	fmt.Println("map length:", len(day))
	fmt.Println(day)
	fmt.Println(day["tonight"])

	days, exist := day["friday"]
	days1, exist1 := day["friday"]
	fmt.Println(days, exist)




	// m1 := make(map[string]int, 20)
	// fmt.Println("map:", m1)
	// fmt.Println("map1 length:",len(m1))
	
}