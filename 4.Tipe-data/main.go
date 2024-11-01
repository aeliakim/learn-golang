package main

import "fmt"

func main() {
	// Tipe data numerik non-desimal
	/*
		Tipe data	Cakupan bilangan
		uint8		0 ↔ 255
		uint16		0 ↔ 65535
		uint32		0 ↔ 4294967295
		uint64		0 ↔ 18446744073709551615
		uint		sama dengan uint32 atau uint64 (tergantung nilai)
		byte		sama dengan uint8
		int8		-128 ↔ 127
		int16		-32768 ↔ 32767
		int32		-2147483648 ↔ 2147483647
		int64		-9223372036854775808 ↔ 9223372036854775807
		int			sama dengan int32 atau int64 (tergantung nilai)
		rune		sama dengan int32
	*/
	// jangan asal menggunakan tipe data, karena efeknya ke alokasi memori

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// Tipe data numerik desimal
	/*
		-float32
		-float64
		perbedaan keduanya yaitu lebar cakupan nilai desimal yang bisa ditampung
		untuk lebih jelasnya: https://www.h-schmidt.net/FloatConverter/IEEE754.html
	*/

	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)   // default yg ditampilkan 6 digit belakang koma
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) // yg ditampilkan 3 digit belakang koma

	// Tipe data boolean
	/*
		-true
		-false
	*/
	// biasa digunakan untuk perulangan dan percabangan

	var exist bool = true
	fmt.Printf("exist? %t \n", exist) // %t untuk memformat data bool

	// Tipe data string
	/*
		nilainya diapit oleh petik dua ("")
	*/

	var message string = "halo"
	fmt.Printf("message: %s \n", message)
	// bisa menggunakan backticks (``) sehingga apapun yg ada di dalamnya terdeteksi sebagai string dan tidak perlu diescape (\n)
	var pesan = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar Golang".`

	fmt.Println(pesan)
}
