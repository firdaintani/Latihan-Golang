package main

import "fmt"

//Employee is about employee data
type Employee struct {
	firstName, lastName string
	birthYear           int
	address             string
}

func (e *Employee) setProfile(first, last, addr string, year int) {
	e.firstName = first
	e.lastName = last
	e.birthYear = year
	e.address = addr
}

func (e Employee) printData() {
	fmt.Println("Nama ", e.firstName, e.lastName)
	fmt.Println("Tahun lahir ", e.birthYear)
	fmt.Println("alamat ", e.address)

}

func main() {
	var birthYear int
	fmt.Printf("Masukkan first Name: ")
	var firstName string
	fmt.Scanf("%s\n", &firstName)
	fmt.Printf("Masukkan last Name:")
	var lastName string
	fmt.Scanf("%s\n", &lastName)
	fmt.Printf("Masukkan tahun lahir:")
	fmt.Scanf("%d\n", &birthYear)
	fmt.Printf("Masukkan alamat:")
	var address string
	fmt.Scanf("%s\n", &address)

	user := Employee{}
	user.setProfile(firstName, lastName, address, birthYear)
	user.printData()

}
