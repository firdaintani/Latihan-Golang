package main

import "fmt"

func main(){
	var a = 3
	if a> 5 {
		fmt.Println("lebih besar dari 5")
	}else if a>2 && a<=5{
		fmt.Println("kurang dari 5 dan lebih besar dari 2")
	}else {
		fmt.Println("kurang dari sama dengan 2")
	
	}
	
	var c = 13
	if( c==12 ) {
		fmt.Println("c sama degn 12")
		
	}else if(c>5 && c<15){
	fmt.Println("c lebih besar dari 5 dan c lebih kecil dari 15")
	}else if( c>5 || c<2){
		fmt.Println("c lebih besar dari 5 atau c lebih kecil dari 2")
	}
		
	

}