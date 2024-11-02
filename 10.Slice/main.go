package main

import "fmt"

func main() {
	// Slice : reference elemen array. Slice merupakan tipe data reference

	// Inisialisasi slice sama seperti array, bedanya tidak perlu menuliskan jumlah elemen
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	// slice bisa dibentuk dari 2 array yang sudah didefinisikan dengan teknik 2 index
	var newFruits = fruits[0:2] // pengaksesan elemen dalam slice fruits yang dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-2
	fmt.Println(newFruits)
	/*
		Kode		Output							Penjelasan
		fruits[0:2]	[apple, grape]					semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
		fruits[0:4]	[apple, grape, banana, melon]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
		fruits[0:0]	[]								menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
		fruits[4:4]	[]								menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
		fruits[4:0]	[]								error, pada penulisan fruits[a:b] nilai a harus lebih kecil atau sama dengan b
		fruits[:]	[apple, grape, banana, melon]	semua elemen
		fruits[2:]	[banana, melon]					semua elemen mulai indeks ke-2
		fruits[:2]	[apple, grape]					semua elemen hingga sebelum indeks ke-2
	*/

	/*
		jika ada slice baru yang terbentuk dari slice lama, maka data elemen slice yang baru akan memiliki alamat memori yang sama dengan elemen slice lama. Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.
	*/
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	baFruits[0] = "pinnaple" // Buah "grape" diubah menjadi "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon] -> perubahan pada slice terbaru berdampak pada slice asli
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	// fungsi len() digunakan untuk menghitung jumlah slice
	fmt.Println(len(fruits))

	// fungsi cap() digunakan untuk menghitung kapasitas maksimum slice
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	/*
		Kode			Output					len()	cap()
		fruits[0:4]		[buah buah buah buah]	4		4
		aFruits[0:3]	[buah buah buah ----]	3		4
		bFruits[1:4]	---- [buah buah buah]	3		3

		fruits[x:y]
		jika slice dimulai dari indeks 0, lebar kapasitasnya sama dengan slice aslinya.
		jika slice dimulai dari indeks x dimana nilai x adalah lebih dari 0, indeks x tersebut menjadi indeks 0 di slice baru, sehingga kapasitas slice berubah.
	*/

	// fungsi append() digunakan untuk menambahkan elemen pada slice dengan memposisikan elemen baru setelah indeks paling akhir
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)  // [apple pinnaple banana melon]
	fmt.Println(cFruits) // ["apple", "pinnaple", "banana", "melon", "papaya"]

	/*
		-jika append() digunakan pada slice baru dan jumlah elemen dan lebar kapasitas sebelum append sama dengan slice asli, elemen baru yg diappend pada slice baru adalah elemen referensi baru.
		-jika append digunakan pada slice baru yang yang jumlah elemen lebih kecil daripada lebar kapasitas, elemen baru ditempatkan pada cakupan kapasitas. misalnya jika pada slice asli indeks ke-3 adalah melon, dibuat slice baru dengan indeks terakhir banana dan digunakan append pada slice tersebut dengan nilai papaya, nilai melon pada slice asli akan berubah menjadi papaya.

		var fruits = []string{"apple", "grape", "banana"}
		var bFruits = fruits[0:2]

		fmt.Println(cap(bFruits)) // 3
		fmt.Println(len(bFruits)) // 2

		fmt.Println(fruits)  // ["apple", "grape", "banana"]
		fmt.Println(bFruits) // ["apple", "grape"]

		var cFruits = append(bFruits, "papaya")

		fmt.Println(fruits)  // ["apple", "grape", "papaya"]
		fmt.Println(bFruits) // ["apple", "grape"]
		fmt.Println(cFruits) // ["apple", "grape", "papaya"]
	*/

	// fungsi copy() digunakan untuk mencopy elemen slice pada source ke destination
	// copy(dst, src)
	dst := make([]string, 3) // lebar 3 elemen
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple -> yg tercopy hanya 3 karena copy() hanya mencopy elemen sebanyak len(dst)
	fmt.Println(src)
	fmt.Println(n) // menampilkan informasi angka (jumlah elemen yang berhasil dicopy)

	/*
		dst := []string{"potato", "potato", "potato"}
		src := []string{"watermelon", "pinnaple"}
		n := copy(dst, src)

		fmt.Println(dst) // watermelon pinnaple potato -> semua elemen src tercopy ke slice dst karena len(src) kurang dari len(dst) dan mengubah nilai
							indeks 0 dan 1
		fmt.Println(src) // watermelon pinnaple
		fmt.Println(n)   // 2
	*/

	// Elemen slice dapat diakses dengan teknik index 3 dengan menambahkan kapasitas slice baru. Kapasitas slice baru tidak boleh melebihi kapasitas slice asli
	aFruits = fruits[0:2]
	bFruits = fruits[0:2:2]

	fmt.Println(fruits)      // [apple pinnaple banana melon]
	fmt.Println(len(fruits)) // 4
	fmt.Println(cap(fruits)) // 4

	fmt.Println(aFruits)      // [apple pinnaple]
	fmt.Println(len(aFruits)) // 2
	fmt.Println(cap(aFruits)) // 4

	fmt.Println(bFruits)      // [apple pinnaple]
	fmt.Println(len(bFruits)) // 2
	fmt.Println(cap(bFruits)) // 2
}
