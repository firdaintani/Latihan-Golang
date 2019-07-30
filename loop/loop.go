package main

import (
	"fmt"
)

func main() {
//while loop
	n := 1
	for n<5{
		fmt.Println("sebelum ",n)
		n*=2
		
		fmt.Println("sesudah ",n)
		
	}
	fmt.Println(n)

//for-each range loop
	strings := []string{"hello","kawan", "selamat","siang"}
	for i, s := range strings{
		fmt.Println(i,s)
	}

//exit loop
	total :=1
	for i:=0;i<=9;i++{
	fmt.Println("indeks looping",i)
		if i%2 == 0{
		continue 
		}
		fmt.Println("indeks",i)
		total*=2
		
	}
	
	fmt.Println(total)


	for j:=0;j<5;j++{
		for k:=0;k<=j;k++{
			fmt.Print("@")
		}
		fmt.Print("\n")
	}
	outer:
	for l := 0; l < 3; l++ {
        for m := 1; m < 4; m++ {
            fmt.Printf("l = %d , m = %d\n", l, m)
            if l == m {
                break outer
            }
        }
	
	
	

    }

}