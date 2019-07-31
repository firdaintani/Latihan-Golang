package main

import "fmt"
import "math/rand"

func addNumber(a int, b int) int  {
	return a+b
	
}

func getRandomNumber()int{
	return rand.Intn(100)
}

func main(){
	fmt.Println(addNumber(1,2))
	fmt.Println(getRandomNumber())
}

