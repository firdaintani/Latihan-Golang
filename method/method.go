package main

import "fmt"

//People struct is about people profile
type People struct {
	name    string
	age     int
	address string
}

//Car is about car profile such as model, year, brand
type Car struct {
	brand string
	model string
	year  int
}

func (p People) describeMe() {
	fmt.Println("My name is", p.name)
	fmt.Println("I am", p.age, "years old")
	fmt.Println("I am from", p.address)

}

func (c Car) countYear() int {
	return 2019 - c.year
}

func (c Car) describeMe() {
	fmt.Println("model", c.model, "from brand", c.brand, "released at", c.year)
}

func (p *People) setProfile(name string, age int, address string) {
	p.name = name
	p.age = age
	p.address = address
}

func main() {
	dina := People{"Dina Ananta", 30, "Jakarta"}
	dina.describeMe()

	tina := People{}
	tina.setProfile("tina", 20, "Bogor")
	tina.describeMe()

	mt48 := Car{"Ferarri", "mt48", 2016}
	mt48.describeMe()
	fmt.Println("it is been", mt48.countYear(), "year since its release")

}
