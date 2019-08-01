package main

import "fmt"

//Book is a struct about name author and year of a book
type Book struct {
	Name   string
	author string
	year   int
}

func (b *Book) SetName(name string) {
	b.name = name
}

//Library is the struct about library
type Library struct {
	location string
	year     int
	Book
}

func X(b *Book) {
	b.name = "asd"
} 

func main() {
	b := Book{}
	X(&b)
	l = []Book{}
	for i, b := range l {
	} 		l[i].Name = 123

	f
	
	
	// fmt.Println("struct")

	// type student struct {
	// 	name  string
	// 	grade int
	// }
	// var s1 student
	// s1.name = "john wick"
	// s1.grade = 2
	// fmt.Println(s1)

	// type Book struct {
	// 	bookName string
	// 	year     int
	// 	author   string
	// }

	// var HarPot Book
	// HarPot.bookName = "Harry potter"
	// HarPot.year = 1998
	// HarPot.author = "JK Rowling"

	// fmt.Println(HarPot)

	// var conan = Book{"Conan Edogawa", 1990, "Aoyama Gosho"}
	// fmt.Println(conan)

	// type Car struct {
	// 	model, name, color string
	// 	weightInKg         float64
	// }

	// c := Car{
	// 	name:  "Ferarri",
	// 	model: "gt48",
	// 	color: "red",
	// }

	// fmt.Println(c.weightInKg)

	// d := Car{"toyota", "red", "avanza", 1221}

	// fmt.Println(d)

	// var e = Car{}
	// fmt.Println(e.name)
	// f := Car{color: "blue"}
	// f.weightInKg = 1201
	// fmt.Println(f)

	// d := struct {
	// 	name string
	// 	age  int
	// }{
	// 	name: "tina",
	// }
	// fmt.Println(d)
	harry := &Book{
		name:   "Harry Potter",
		author: "JK Rowling",
		year:   1998,
	}
	fmt.Println("name :", (*harry).name)
	fmt.Println("name :", harry.name)
	fmt.Println(&harry)

	conan := &Book{
		name:   "Conan",
		author: "Aoyama Gosho",
		year:   1990,
	}
	fmt.Println("name :", (*conan).name)
	fmt.Println("name :", conan.name)
	fmt.Println(conan)
	fmt.Println(&conan)

	// naruto := Book{
	// 	name:   "naruto",
	// 	author: "asdasd",
	// 	year:   2012,
	// }
	// fmt.Println("name :", naruto.name)
	// fmt.Printf("alamat :%p", &naruto)
	// var libraryJakarta = Library{}
	// libraryJakarta.location = "Jakarta"
	// libraryJakarta.author = "Aoyama Gosho"
	// libraryJakarta.name = "Conan"
	// libraryJakarta.year = 1999
	// libraryJakarta.Book.year = 2000
	// fmt.Println(libraryJakarta)

}
