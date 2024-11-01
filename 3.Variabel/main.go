package main

import "fmt"

func main() {
	// tipe manifest typing (deklarasi dengan tipe datanya)
	var firstName string = "john"

	// tipe inference (tipe data-nya diketahui secara otomatis dari data/nilai variabel)
	/* menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstName = "john" */
	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName := "wick"
	// Tanda := hanya digunakan sekali di awal pada saat deklarasi. Untuk assignment nilai selanjutnya harus menggunakan tanda =
	lastName = "ethan"
	lastName = "bourne"

	// deklarasi multi variabel
	/*
		var first, second, third string
		first, second, third = "satu", "dua", "tiga"
	*/

	// pengisian nilai bersamaan deklarasi
	/*
		var fourth, fifth, sixth string = "empat", "lima", "enam"
	*/

	// versi ringkas
	/*
		seventh, eight, ninth = "tujuh", "delapan", "sembilan"
	*/

	// variabel dengan tipe data berbeda
	/*
		one, isFriday, twoPointtwo, say := 1, true, 2.2, "hello"
	*/

	// reserved variable (underscore) --> menampung nilai yang tidak digunakan dan tidak akan ditampilkan
	_ = "belajar golang"
	_ = "golang itu mudah"
	name, _ := "john", "wick" // nilai wick tidak digunakan

	// deklarasi variabel menggunakan keyword new()
	nama := new(string) // membuat variabel pointer
	fmt.Println(nama)   // 0xc0000260a0 --> menampilkan alamat memori nilai
	fmt.Println(*nama)  // pakai * untuk dereference (menampilkan isi dari memori)

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println(name)
}
