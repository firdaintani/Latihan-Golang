package main

import "fmt"

func main() {
	totalNilai := jumlah(1, 23, 4, 5, 2, 1, 4)
	fmt.Println(totalNilai)

	slice1 := []string{
		"senin",
		"selasa",
		"rabu",
	}

	fmt.Println(slice1)
	slice2 := []string{
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	// fmt.Println(slice2)
	slice1 = append(slice1, slice2...)
	// fmt.Println(slice1)
	sliceBaru := printTeks(slice1...)
	// fmt.Printf("alamat memori slice1 %p \n", &slice1)
	// fmt.Printf("alamat memori slice dari return function %p", &sliceBaru)
	fmt.Println(slice1)

	fmt.Println(sliceBaru)
	names := []string{"dina", "ananta"}
	changeName(names...)
	fmt.Println("setelah function dipanggil :", names)

	totalHitung := operasi("perkalian", 1, 2, 3, 4)
	fmt.Println(totalHitung)
}

func printTeks(teks ...string) []string {

	var newSlice = []string{}
	for i := 0; i < len(teks)-1; i++ {
		// fmt.Println(teks[i])
		newSlice = append(newSlice, teks[i])
	}
	return newSlice
}

func jumlah(nilai ...int) int {
	total := 0
	for _, val := range nilai {
		total += val
	}
	return total
}

func changeName(name ...string) {

	namanya := name[0:1]
	namanya[0] = "den"
	fmt.Println(namanya)
}

// func jumlah(nilai ...int) total int {
// 	total := 0
// 	for _, val := range nilai {
// 		total += val
// 	}
// 	return
// }

func operasi(jenis string, angka ...int) int {
	hasil := 0
	if jenis == "perkalian" {
		hasil = 1
		for _, val := range angka {
			hasil *= val
		}
	} else if jenis == "penjumlahan" {
		for _, val := range angka {
			hasil += val
		}
	}
	return hasil
}
