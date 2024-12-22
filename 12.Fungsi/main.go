package main

/*
* Penerapan fungsi
* func namaFungsi (bisaDiisiParameter int) {
* }
* parameter bersifat opsional, digunakan tergantung kebutuhan
* Data yang digunakan sebagai value parameter saat pemanggilan fungsi biasa disebut dengan argument parameter (atau argument)
 */

import (
	"fmt"
	// "strings"
	// "math/rand"
	// "time"
)

// fungsi utama saat program dijalankan
// func main() {
// 	var names = []string{"John", "Wick"}
// 	printMessage("halo", names)
// }

// fungsi yg dibuat dengan 2 parameter
// func printMessage(message string, arr[] string) {
// 	var nameString = strings.Join(arr, " ") // menggabungkan 2 string dengan pemisah spasi
// 	fmt.Println(message, nameString)
// }

/*
* Fungsi yang memiliki return value, saat deklarasinya harus ditentukan terlebih dahulu tipe data dari nilai baliknya.
* Fungsi yang tidak mengembalikan nilai apapun disebut dengan void function
 */

/*
 * rand.New(): membuat object randomizer untuk mendapatkan nilai random hasil generator. Fungsi ini membutuhkan argument yaitu random source seed yang bisa dibuat dengan statement
 * rand.NewSource(time.Now().Unix())
 * argument rand.NewSource() bisa diisi dengan nilai apapun, salah satunya adalah time.Now().Unix()
 * fungsi rand.New() merupakan bagian dari package time
 */
// var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

// func main() {
// 	var randomValue int

// 	randomValue = randomWithRange(2,10)
// 	fmt.Println("random number: ", randomValue)

// 	randomValue = randomWithRange(2,10)
// 	fmt.Println("random number: ", randomValue)

// 	randomValue = randomWithRange(2,10)
// 	fmt.Println("random number: ", randomValue)
// }

// penentuan tipe data yang akan dikembalikan dilakukan setelah deklarasi parameter
// untuk fungsi yang 2 atau lebih tipe data parameternya sama, bisa ditulis dengan nama seluruh parameter yang digunakan kemudian diikuti tipe datanya.
// func randomWithRange(min, max int) int {
// 	var value = randomizer.Int()%(max-min+1) + min
// 	return value // mengembalikan nilai dari variabel ke baris kode fungsi dipanggil
// }

func main() {
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
}

// tidak memiliki nilai balik namun menampilkan hasilnya di dalam fungsi
func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot be divided by %d\n", m, n)
		return // bisa dimanfaatkan untuk menghentikan proses dalam blok fungsi
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}


