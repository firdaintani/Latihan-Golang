package main

import "fmt"
import "reflect"

//Car is interface
type Car interface {
	Description()
}

type truck struct {
	model       string
	brand       string
	weight      int
	totalMuatan int
}

type standarCar struct {
	model  string
	brand  string
	color  string
	weight int
}

func (s standarCar) Description() {
	fmt.Println("Model : ", s.model)
	fmt.Println("Brand : ", s.brand)
	fmt.Println("Color : ", s.color)
	fmt.Println("Weight : ", s.weight)
}

func (t truck) Description() {
	fmt.Println("Model : ", t.model)
	fmt.Println("Brand : ", t.brand)
	fmt.Println("Weight : ", t.weight)
	fmt.Println("total muatan :", t.totalMuatan)

}

func (t truck) ShowWeight() {
	fmt.Println(t.weight)
}

func showDescription(c Car) {
	c.Description()
}

//
func printApaaja(a interface{}) {
	fmt.Printf("data %v dan tipenya %T \n", a, a)
}

func main() {
	var colt Car
	colt = truck{"colt diesel 343", "Hino", 200, 1000}
	// avanza := standarCar{"avanza", "toyota", "black", 100}
	// colt.Description()
	colt.Description()
	colt.(truck).ShowWeight()

	var apaaja interface{}
	apaaja = "great"
	fmt.Println("tipenya :", reflect.TypeOf(apaaja))
	printApaaja(1)
	printApaaja("sdsds")
	printApaaja(colt)
	showDescription(colt)
}
