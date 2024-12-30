package main

import (
	"fmt"
)

/*
* Fungsi Closure
* Fungsi tanpa nama yang disimpan di dalam variabel. Biasa dimanfaatkan untuk membungkus suatu proses yang hanya dijalankan sekali saja atau hanya dipakai pada blok tertentu.
 */

/*
* Closure disimpan sebagai variabel
 */

// func main() {
// Fungsi closure sebagai variabel: fungsi ditulis tanpa nama dan ditampung ke variabel
// 	var getMinMax = func(n []int) (int, int) {
// 		var min, max int
// 		for i, e := range n {
// 			switch {
// 			case i == 0:
// 				max, min = e, e
// 			case e > max:
// 				max = e
// 			case e < min:
// 				min = e
// 			}
// 		}
// 		return min, max
// 	}

// 	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
// pemanggilan variabel closure dilakukan dengan pemanggilan fungsi biasa
// 	var min, max = getMinMax(numbers)
// 	fmt.Printf("data: %v\nmin: %v\nmax: %v\n", numbers, min, max) // %v digunakan untuk menampilkan data tanpa melihat tipe datanya. Biasa digunakan untuk debugging
// }

/*
* Closure jenis IIFE (Immediately-Invoked Function Expression)
* Biasa dilakukan untuk membungkus proses yang dilakukan hanya sekali.
 */

// func main() {
// var numbers = []int{2, 3, 0, 4, 3, 2,0, 4, 2, 0, 3}

// eksekusi fungsi closure dilakukan saat deklarasi variabel
// yg ditampung adalah nilai baliknya, bukan body fungsi atau closure nya.
// var newNumbers = func(min int) []int {
// 	var r []int
// 	for _, e := range numbers {
// 		if e < min {
// 			continue
// 		}
// 		r = append(r, e)
// 	}
// bisa memiliki nilai balik atau tidak
// return r
// } (3)  // ciri khas closure IIFE

// fmt.Println("original number: ", numbers)
// fmt.Println("filtered number: ", newNumbers)
// }

// Closure dengan gaya manifest typing
// var closure (func(string, int, []string) int)
// closure = func(a string, b int, c []string) int {
// }

/*
* Closure sebagai nilai balik fungsi
 */

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	// nilai balik berupa closure dengan skema func() []int
	return len(res), func() []int {
		return res
	}

	// fungsi closure yang akan dikembalikan boleh disimpan pada variabel terlebih dahulu
	// var getNumbers = func() []int {
	// return res
	// }
	// return len(res), getNumbers
}

func main() {
	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)
}


