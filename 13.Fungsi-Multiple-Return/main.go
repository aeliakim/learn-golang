package main

import (
	"fmt"
	"math"
)

/*
* Penerapan fungsi multiple return
 */
// tulis semua tipe data nilai return yang ingin digunakan di dalam kurung.
// func calculate(d float64) (float64, float64) {
// hitung luas
// 	var area = math.Pi * math.Pow(d / 2, 2)
// hitung keliling
// 	var circumference = math.Pi * d

// kembalikan 2 nilai
// tulis semua data yang ingin dikembalikan
// 	return area, circumference
// }

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

/*
* Fungsi dengan multiple return
 */
// di Go, variabel yang digunakan sebagai nilai balik bisa didefinisikan di awal
func calculate(d float64) (area float64, circumference float64){
	// math.Pow(): operasi pangkat nilai. math.Pow(2, 3) artinya 2 pangkat 3, hasilnya 8
	// math.Pi: konstanta bawaan package math yang merepresentasikan Pi atau 22/7
	area = math.Pi * math.Pow(d / 2, 2)
	circumference = math.Pi * d

	return // karena variabel nilai balik sudah ditentukan di awal, tidak perlu memanggil variabel lagi.
}
