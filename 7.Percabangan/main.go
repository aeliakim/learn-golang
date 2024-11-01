package main

import "fmt"

func main() {
	// Percabangan if, else if, else
	var point = 8

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	// Variabel temporary di if-else
	/*
		variabel yang hanya bisa digunakan pada deretan blok seleksi kondisi di mana ia ditempatkan.
		Manfaat:
		-Scope atau cakupan variabel jelas, hanya bisa digunakan pada blok seleksi kondisi itu saja
		-Kode menjadi lebih rapi
		-Ketika nilai variabel tersebut didapat dari sebuah komputasi, perhitungan tidak perlu dilakukan di dalam blok masing-masing kondisi.
	*/

	var poin = 8840.0
	// Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan tanda :=. Penggunaan keyword var tidak diperbolehkan karena menyebabkan error.
	if percent := poin / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	}

	// Percabangan switch-case
	var nilai = 6

	switch nilai {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4: // dapat menampung banyak kondisi
		fmt.Println("awesome")
	default:
		// tanda kurung kurawal ({}) opsional, bisa diterapkan pada blok kondisi yang di dalamnya ada banyak statement, dengannya kode akan terlihat lebih rapi
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// switch dengan gaya if-else
	switch {
	case nilai == 8:
		fmt.Println("perfect")
	case (nilai < 8) && (nilai > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// Penggunaan keyword fallthrough dalam switch
	/*
		Keyword fallthrough digunakan untuk memaksa proses pengecekan tetap diteruskan ke case selanjutnya dengan tanpa menghiraukan nilai kondisinya, efeknya adalah case di pengecekan selanjutnya selalu dianggap true (meskipun aslinya bisa saja kondisi tersebut tidak terpenuhi, akan tetap dianggap true).
		Pada case dalam sebuah switch, diperbolehkan terdapat lebih dari satu fallthrough
	*/
	switch {
	case nilai == 8:
		fmt.Println("perfect")
	case (nilai < 8) && (nilai > 3):
		fmt.Println("awesome")
		fallthrough
	case nilai < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// Percabangan bersarang (nested conditions)
	/*
		Seleksi kondisi bersarang bisa dilakukan pada if - else, switch, ataupun kombinasi keduanya
	*/

	point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it!")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
