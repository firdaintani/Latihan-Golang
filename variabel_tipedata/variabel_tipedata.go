package main

import "fmt"
import "reflect"

func main(){
	var a = "initial"
	fmt.Println(a)
	
	var b,c int = 1,2
	fmt.Println(b,c)
	
	var d= true
	fmt.Println(d)
	
	var e int
	fmt.Println(e)
	
	f:="apple"
	fmt.Println(f)

	var name string = "Firda"
	fmt.Println("nama saya :",name)
	fmt.Printf("nama saya : %s\n", name)
	fmt.Println(`Hello today is happy day`)
	fmt.Printf(`Hello today is happy day \n`)
	
	
	var date, year int
	date = 5
	year = 2019
	fmt.Print("tanggal ", date, " tahun ", year, "\n")
	
	var pi = 3.14
	fmt.Println(reflect.TypeOf(pi))
	var nilai = reflect.ValueOf(pi)
	
	fmt.Println("nilai variabel pi :",nilai, " dan tipe datanya : ", nilai.Kind())
	
	var day1, day2, day3 = "senin", "selasa", "rabu"
	fmt.Printf("%s, %s, %s \n", day1, day2,day3)
	
	day4 := "kamis"
	fmt.Println(day4)
	
	fmt.Println("2.4 + 3.5=", 2.4+3.5)

}