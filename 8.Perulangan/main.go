package main

import "fmt"

func main() {
	// Perulangan di Golang hanya ada "for"

	// Cara 1: memasukkan variabel counter perulangan beserta kondisinya setelah keyword
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// cara 2: menuliskan kondisi setelah keyword for (mirip while)
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	// cara 3: for ditulis tanpa kondisi. Pemberhentian perulangan dilakukan dengan menggunakan keyword break
	// for {
	// 	fmt.Println("Angka", i)
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// cara 4: menggunakan kombinasi keyword for dan range. digunakan untuk me-looping data gabungan (string, array, slice, map)
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index =", i, "Value =", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} //array
	for _, v := range ys {
		fmt.Println("Value = ", v)
	}

	var zs = ys[0:2] //slice
	for _, v := range zs {
		fmt.Println("Value = ", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} //map
	for k, v := range kvs {
		fmt.Println("Key =", k, "Value =", v)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Print(i) //01234
	}
}
