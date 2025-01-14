package main

import "fmt"

/*
* Struct
* Kumpulan definisi variabel (atau property) atau fungsi (atau method) yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda
* Variabel baru bisa dibuat dari sebuah struct yang memiliki atribut sesuai skema struct tersebut. Variabel tersebut bisa dipanggil dengan istilah object atau variabel object
 */

/*
* Deklarasi struct
* Kombinasi keyword type dan struct digunakan untuk deklarasi struct
 */

// type student struct {
// property -> variabel yang menempel di struct
// 	name string // name -> property dengan tipe data string
// 	grade int
// }

/*
* Penerapan struct untuk membuat object
 */

// func main() {
// Cara membuat variabel objek sama dengan membuat variabel biasa. [nama variabel] [nama struct]
// 	var s1 student
// Property variabel objek bisa diakses nilainya dengan menggunakan notasi titik. Nilai propertynya juga bsia diubah.
// 	s1.name = "john wick"
// 	s1.grade = 2

// 	fmt.Println("name	:", s1.name)
// 	fmt.Println("grade	:", s1.grade)
// }

/*
* Inisialisasi object struct
 */

// func main() {
// cara inisialisasi variabel objek yaitu dengan menuliskan nama struct yang telah dibuat diikuti dengan tanda kurung kurawal.
// variabel s1 menampung objek cetakan student kemudian di-set nilai propertynya.
// var s1 = student{}
// s1.name = "wick"
// s1.grade = 2

// dideklarasikan dengan metode yang sama dengan s1, bedanya nilai property s2 diisi langsung ketika deklarasi
// var s2 = student{"ethan", 2}

// pengisian property dilakukan ketika pencetakan objek. bedanya, yang diisi hanya property name saja.
// var s3 = student{name: "jason"}
// dengan cara yg sama, penentuan nilai property bisa dilakukan degan tidak berurutan
// var s4 = student{name: "wayne", grade: 2}
// var s5 = student{grade: 2, name: "bruce"}

// 	fmt.Println("student 1: ", s1.name)
// 	fmt.Println("student 2: ", s2.name)
// 	fmt.Println("student 3: ", s3.name)
// }

/*
* Variabel objek pointer
 */

// func main() {
// 	var s1 = student{name:"wick", grade: 2}

// objek yang dibuat dari tipe struct bisa diambil nilai pointernya dan bisa disimpan pada variabel objek yang bertipe struct pointer
// 	var s2 *student = &s1 // s2 menampung nilai referensi s1
// 	fmt.Println("student 1, name: ", s1.name)
// 	fmt.Println("student 4, name: ", s2.name) // property s2 tetap dapat diakses seperti biasa tanpa di-dereferensi

// 	s2.name = "ethan" // pengisian nilai bisa langsung menggunakan nilai asli
// 	fmt.Println("student 1, name: ", s1.name) // nilai s2 berubah -> nilai s1 juga berubah karena mereferensi alamat asli
// 	fmt.Println("student 4, name: ", s2.name)
// }

/*
* Embedded Struct
* Mekanisme untuk menempelkan sebuah struct sebagai properti struct lain
* Mutable -> nilai propertynya bisa diubah
 */

// type person struct {
// 	name string
// 	age int
// }

// type student struct {
// 	grade int
// 	person // struct person di-embed ke dalam struct student
// }

// func main() {
// 	var s1 = student{}
// 	s1.name = "wick"
// 	s1.age = 21
// 	s1.grade = 2

// 	fmt.Println("name :", s1.name)
// 	fmt.Println("age :", s1.age)
// Khusus untuk properti yang bukan merupakan properti asli (melainkan properti turunan dari struct lain), pengaksesannya dilakukan dengan cara mengakses struct parent-nya terlebih dahulu
// 	fmt.Println("age :", s1.person.age)
// 	fmt.Println("grade :", s1.grade)
// }

/*
* Embedded Struct dengan nama property yang sama
* Jika salah satu nama properti sebuah struct memiliki kesamaan dengan properti milik struct lain yang di-embed, maka pengaksesan property-nya harus dilakukan secara eksplisit atau jelas.
 */

// type student struct {
// 	person
// 	age int
// 	grade int
// }

// func main() {
// 	var s1 = student{}
// 	s1.name = "wick"
// 	s1.age = 21 // age of student
// 	s1.person.age = 22 // age of person // untuk mengakses property age dari struct person

// 	fmt.Println(s1.name)
// 	fmt.Println(s1.age)
// 	fmt.Println(s1.person.age)
// }

/*
* Pengisian nilai Sub-Struct
 */

// func main() {
// 	var p1 = person{name: "wick", age: 21}
// 	var s1 = student{person: p1, grade: 2} // langsung masukan variabel objek yang tercetak dari struct yg sama

// 	fmt.Println("name: ", s1.name)
// 	fmt.Println("age: ", s1.age)
// 	fmt.Println("grade: ", s1.grade)
// }

/*
* Anonymous Struct
* Struct yang tidak dideklarasikan di awal sebagai tipe data baru, melainkan langsung ketika pembuatan objek.
 */

// type person struct {
// 	name string
// 	age int
// }

// func main() {
// variabel s1 langsung diisi objek anonymous struct yg memiliki property grade dan property person yang merupakan embedded struct
// pengisian anonymous struct tanpa pengisian property
// var s1 = struct {
// 	person
// 	grade int
// }{}

// pengisian anonymous struct dengan pengisian property
/*
	var s2 = struct {
		person
		grade int
		}{
			person: person{"wick", 21},
			grade: 2,
		}
*/

// 	s1.person = person{"wick", 21}
// 	s1.grade = 2

// 	fmt.Println("name: ", s1.person.name)
// 	fmt.Println("age: ", s1.person.age)
// 	fmt.Println("grade: ", s1.grade)
// }

/*
* Kombinasi slice & struct
 */

// type person struct {
// 	name string
// 	age int
// }

// func main() {
// kombinasi slice & struct
// 	var allStudents = []person{
// 		{name: "Wick", age: 23},
// 		{name: "Ethan", age: 23},
// 		{name: "Bourne", age: 22},
// 	}

// 	for _, student := range allStudents {
// 		fmt.Println(student.name, "age is", student.age)
// 	}
// }

/*
* Inisialisasi Slice Anonymous Struct
 */

// func main(){
// 	var allStudents = []struct {
// 		person
// 		grade int
// 	}{
// 		{person: person{"wick", 21}, grade: 2},
// 		{person: person{"ethan", 22}, grade: 3},
// 		{person: person{"bond", 21}, grade: 3},
// 	}

// 	for _, student := range allStudents {
// 		fmt.Println(student)
// 	}
// }

/*
* Deklarasi Anonymous Struct Menggunakan Keyword var
 */

// func main() {
// dicetak sebuah objek dari anonymous struct kemudian disimpan pada variabel bernama student
// var student struct {
// 	person
// 	grade int
// }

// inisialisasi data
// hanya deklarasi
/*
	var student struct {
		grade int
	}
*/

// deklarasi sekaligus inisialisasi
/*
	var student = struct {
		grade int
	}{
		12,
	}
*/

// 	student.person = person{"wick", 21}
// 	student.grade = 2
// }

/*
* Nested Struct
* Anonymous struct yang di-embed ke sebuah struct. Deklarasinya langsung di dalam struct peng-embed.
* Biasa digunakan ketika decoding data JSON yg struktur datanya cukup kompleks dengan proses decode hanya sekali.
 */

// type student struct {
// 	person struct {
// 		name string
// 		age int
// 	}
// 	grade int
// 	hobbies []string
// }

/*
* Deklarasi dan Inisialisasi Struck secara Horizontal
 */
// tanda semicolon digunakan sebagai pembatas deklarasi property yang akan dituliskan secara horizontal
// type person struct { name string; age int; hobbies []string }

// func main() {
// contoh inisialisasi nilai
// 	var p1 = struct {name string; age int;} {age: 22, name: "wick"}
// 	var p2 = struct {name string; age int;} {"ethan", 23}
// }

/*
* Tag property dalam struct
* Informasi opsional yang bisa ditambahkan pada property struct
* Tag biasa dimanfaatkan untuk keperluan encode/decode data. Informasi tag juga bisa diakses lewat reflect.
 */

// type person struct {
// 	name string `tag1`
// 	age int `tag2`
// }

/*
* Type alias
 */

type Person struct {
	name string
	age int
}
// membuat alias untuk sebuah struct
type People = Person

func main() {
	// var p1 = Person{"wick", 21}
	// fmt.Println(p1)

	// var p2 = People{"wick", 21}
	// fmt.Println(p2)

	// casting dari objek (yang dicetak lewat struct tertentu) ke tipe yang merupakan alias dari struct pencetak, hasilnya selalu valid. Berlaku juga sebaliknya
	people := People{"wick", 21}
	fmt.Println(Person(people))

	person := Person{"wick", 21}
	fmt.Println(People(person))

	// Pembuatan struct baru bisa dilakukan lewat teknik type alias
	/*
	type People1 struct {
		name string
		age int
	}
	type People2 = struct { -> alias dari anonymous struct
		name string
		age int
	}
	*/
}

