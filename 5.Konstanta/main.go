package main

import "fmt"

func main() {
	// Penggunaan konstanta sama seperti deklarasi variabel
	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	// Teknik type inference bisa diterapkan ke konstanta
	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")

	// deklarasi multi konstanta
	// deklarasi dengan tipe data dan nilai yg berbeda
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	// deklarasi dengan tipe dan nilai yg sama
	const (
		a = "konstanta"
		// Ketika tipe data dan nilai tidak dituliskan dalam deklarasi konstanta, maka tipe data dan nilai yang dipergunakan adalah sama seperti konstanta yang dideklarasikan diatasnya.
		b // b = "konstanta" - tipe data string
	)

	// gabungan keduanya
	const (
		today    string = "senin"
		sekarang        // sekarang = "senin"
		isToday2 = true
	)

	// multiple konstanta dalam 1 baris
	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"

}
