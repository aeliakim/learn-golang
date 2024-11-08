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
	var j = 0
	for {
		fmt.Println("Angka", j)
		j++
		if j == 5 {
			break
		}
	}

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

	// penggunaan keyword break & Continue
	for i := 1; i <= 10; i++ {
		// ketika kondisi terpenuhi, dipaksa lanjut ke perulangan selanjutnya
		if i%2 == 1 {
			continue
		}

		// ketika i melebihi 8, perulangan berhenti
		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// perulangan bersarang (nested loops)
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, "")
		}

		fmt.Println()
	}

	// pemanfaatan label dalam perulangan
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
