package main 

import "fmt"

func main(){

	var a = 12
	var b = 2
	var c int
	
	c= a+b
	fmt.Println(" a+b=", c)	
	c=a-b
	fmt.Println(" a-b=", c)	
	c=a/b
	fmt.Println(" a/b=", c)
	c=a*b
	fmt.Println(" a*b=", c)
	a++
	fmt.Println(" a++", a)
	fmt.Println("(2*3)-3+(4/2)=",(2*3)-3+(4/2))
	
	d := 100
	e := 20
	//f:="100"
	
	fmt.Println("100 > 20", d>e)
	fmt.Println("100 < 20", d<e)
	fmt.Println("100 >= 20", d>=e)
	fmt.Println("100 <= 20", d<=e)
	fmt.Println("100 != 20", d!=e)
	fmt.Println("100 == 20", d==e)
	//mismatch
	//fmt.Println("d == f", d==f) 
	
	bool0 := false
	bool1 := true
	fmt.Println(bool0&&bool0)
	fmt.Println(bool0&&bool1)
	fmt.Println(bool1&&bool1)
	fmt.Println(bool1&&bool0)
	
	fmt.Println(bool0||bool0)
	fmt.Println(bool0||bool1)
	fmt.Println(bool1||bool1)
	fmt.Println(bool1||bool0)
	
	fmt.Println("(100>20) && (2>5) = ",(100>20) && (2>5))
	fmt.Println("(100<20) || (2==5) = ",(100<20) || (2==5))
	fmt.Println("((200>4) || (50<100)) && !(21<1)",((200>4) || (50<100)) && !(21<1))	
}