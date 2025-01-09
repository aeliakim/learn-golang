package main

import "fmt"

/*
* Pointer -> reference / alamat memori
* Variabel pointer -> variabel yang berisi alamat memori suatu nilai
* Ex: sebuah variabel integer memiliki nilai 4, yg dimaksud pointer adalah alamat memori di mana nilai 4 disimpan, bukan nilai 4 itu sendiri
* Variabel-variabel yang memiliki reference atau alamat pointer yang sama saling berhubungan satu sama lain dan nilainya pasti sama. Jika nilainya berubah, maka akan memberikan efek kepada variabel lain yang referensinya sama, yaitu nilainya ikut berubah.
* Nilai default pointer adalah nil (kosong). Variabel pointer tidak bisa menampung nilai yang bukan pointer, dan sebaliknya variabel biasa tidak bisa menampung nilai pointer.
 */

/*
* Penerapan pointer
 */

// func main() {
// 	var numberA int = 4
// 	var numberB *int = &numberA // variabel pointer yang berisi alamat variabel numberA

// 	fmt.Println("numberA (value)	:", numberA) // 4
// 	fmt.Println("numberA (address)	:", &numberA) // mengambil nilai pointer dari variabel biasa dengan menambahkan tanda & sebelum variabel (referencing)
// 	fmt.Println("numberB (value)	:", *numberB) // mengambil nilai asli dari variabel pointer dengan menambahkan tanda * sebelum variabel (dereferencing)
// 	fmt.Println("numberB (address)	:", numberB) // alamat variabel numberA
// }

/*
* Efek perubahan nilai pointer
 */

// func main() {
// 	var numberA int = 4
// 	var numberB *int = &numberA

// 	fmt.Println("numberA (value)	:", numberA)
// 	fmt.Println("numberA (address)	:", &numberA)
// 	fmt.Println("numberB (value)	:", *numberB)
// 	fmt.Println("numberB (address)	:", numberB)

// 	fmt.Println("")

// 	numberA = 5

// 	fmt.Println("numberA (value)	:", numberA)
// 	fmt.Println("numberA (address)	:", &numberA)
// 	fmt.Println("numberB (value)	:", *numberB) // value numberB ikut berubah karena mereferensi alamat yg sama dengan numberA
// 	fmt.Println("numberB (address)	:", numberB)
// }

/*
* Parameter pointer
 */

func main(){
	var number = 4
	fmt.Println("before:", number) // 4

	change(&number, 10)
	fmt.Println("after:", number) // 10
}

// original = variabel pointer
func change(original *int, value int) {
	*original = value // nilai dari variabel pointer original diubah
}
