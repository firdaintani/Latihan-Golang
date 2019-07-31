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

	v, ok := day["tomight"]
	if !ok {
		fmt.Println("not found")
		// return
	}
	fmt.Println("value: ", v)
	v2, ok2 := day["today"]
	if !ok2 {
		fmt.Println("not found")
		return
	}
	fmt.Println("value: ", v2)
	// days, exist := day["friday"]
	// days1, exist1 := day["today"]
	// fmt.Println(days, exist)
	// fmt.Println(days1, exist1)

	// delete(day, "yesterday")
	// fmt.Println(day)
	// delete(day, "friday")
	// fmt.Println(day)

	// for key, value := range day {
	// 	fmt.Println (key, " ", value)
	// }

	// data := map[string]int{}
	// data["one"]=1
	// fmt.Println(data["one"])
	
	// keyMap := []string{}
	// for _, value := range day {
	// 	keyMap = append(keyMap, value)
	// }
	// fmt.Println(keyMap)

	// copyMap := map[string]string{}
	// for key, value := range day{
	// 	copyMap[key] = value
	// }
	// fmt.Println(copyMap)
	

}