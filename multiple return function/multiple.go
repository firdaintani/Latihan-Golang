package main

import "fmt"
import "strings"

func main() {
	fmt.Println("multiple return function")
	fmt.Println(splitKata("hello world"))
	luasPersegi, kelilingPersegi := luasKeliling(8)
	fmt.Println(luasPersegi)
	fmt.Println(kelilingPersegi)
}

func splitKata(a string) (string, string) {
	splittedString := strings.Split(a, " ")

	return splittedString[0], splittedString[1]

}

func luasKeliling(a int) (luas int, keliling int) {
	luas = a * a
	keliling = a * 4
	return
}
