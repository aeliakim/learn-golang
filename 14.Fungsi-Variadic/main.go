package main

import (
	"fmt"
	"strings"
)

/*
* Variadic Function: Fungsi dengan parameter yang bisa menampung nilai sejenis yang tidak terbatas jumlahnya
* Parameter variadic memiliki sifat yang mirip dengan slice, dimana nilai dari parameter-parameter yang disisipkan memiliki tipe data yang sama dan cukup ditampung di satu variabel saja. Cara pengaksesan juga sama, yaitu menggunakan index.
 */

/*
*Penerapan fungsi variadic
 */

// func main() {
// 	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
// sama seperti fmt.Printf(), hanya saja tidak menampilkan nilai tetapi mengembalikan nilai dalam bentuk string
// 	var msg = fmt.Sprintf("Rata-rata: %.2f", avg)
// 	fmt.Println(msg)
// }

// Deklarasi parameter variadic menggunakan tanda ... sebelum tipe data dari parameter
// func calculate(numbers ...int) float64 {
// 	var total int = 0
// 	for _, number := range numbers {
// 		total += number
// 	}

// Tipe data jika digunakan sebagai fungsi (ada tanda kurungnya) menandakan tipe data tersebut digunakan untuk casting
// Casting -> teknik untuk konversi tipe sebuah data ke tipe lain
// 	var avg = float64(total) / float64(len(numbers))
// 	return avg
// }

/*
* Pengisian parameter fungsi variadic menggunakan data slice
 */

// func main() {
// 	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
// tulis nama variabel slice beserta tanda titik tiga kali
// var avg = calculate(numbers...)
// atau
// var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
// 	var msg = fmt.Sprintf("Rata-rata: %.2f", avg)

// 	fmt.Println(msg)
// }

/*
* Fungsi dengan parameter biasa dan variadic
 */

// parameter variadic bisa digabungkan dengan parameter biasa dengan syarat parameter variadic harus ditempatkan di akhir
func yourHobbies(name string, hobbies...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func main() {
	// cara pemanggilan seperti fungsi biasa
	yourHobbies("wick", "sleeping", "eating")

	// jika parameter kedua dan seterusnya ingin diisi dengan data slice, gunakan tanda titik tiga kali
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}
