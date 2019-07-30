package main

import "fmt"


func main(){
	i :="e"
	switch i {
	case "a" :
		fmt.Println("i = a")
	case "b" : 
		fmt.Println("i = b")
	case "c" : 
		fmt.Println("i=c")
	case "d", "e" : 
		fmt.Println("i=d atau i=e")
	default : 
		fmt.Println("gaada")
	}
	b:=5
	switch  {
		case b==1 :
		fmt.Println("1")
		case b==2 :
		fmt.Println("2")
		case b>3 :
		fmt.Println("lebih besar dari 3")
        fallthrough
		case b<6 :
		fmt.Println("kurang dari 6")
		default : 
		fmt.Println("selain 1 dan 2")
	}
	 
}