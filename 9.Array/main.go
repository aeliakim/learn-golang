package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	// tidak boleh melebihi kapasitas array. jika lebih, terjadi error
	// names[4] = "ez"

	fmt.Println(names[0], names[1], names[2], names[3])

	// Pengisian array bisa dilakukan pada saat deklarasi variabel
	// cara horizontal
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	/*
		fruits = [4]string{
		"apple"
		"grape",
		"banana",
		"melon", --> tanda koma wajib ditulis di setiap elemen termasuk elemen terakhir agar tidak terjadi syntax error
		}
	*/

	fmt.Println("Jumlah semua element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// Inisialisasi nilai di awal boleh tidak ditulis jumlah lebar arraynya dengan menggunakan tanda 3 titik (...)
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// Array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}} // Khusus penulisan array yang merupakan subdimensi/elemen, boleh tidak dituliskan jumlah datanya

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// Perulangan array menggunakan for
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	// Perulangan array menggunakan for-range
	for i, fruit := range fruits { // i = nilai indeks, fruit = nilai yg ditampung dari array fruits sesuai indeks yg diakses
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// tanda underscore dapat digunakan untuk menampung nilai dari variabel yg tidak ingin digunakan
	for _, fruit := range fruits { // indeks tidak digunakan sehingga bisa diganti _
		fmt.Printf("nama buah: %s\n", fruit)
	}
	/*
		jika yg dibutuhkan hanya indeks, bisa memakai underscore atau tulis 1 variabel penampung nilai indeks setelah for
		for i, _:= range fruits {}
		atau
		for i := range fruits {}
	*/

	// Deklarasi sekaligus alokasi kapasitas array juga bisa dilakukan lewat keyword make
	var buahs = make([]string, 2) // make([]tipe data, lebar array)
	buahs[0] = "apple"
	buahs[1] = "mango"

	fmt.Println(buahs)
}
