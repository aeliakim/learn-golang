package main

import (
	"fmt"
	"strings"
)

/*
* Fungsi sebagai parameter
 */

// parameter callback merupakan closure yang bertipe func(string) bool
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

/*
1. Data array (yang didapat dari parameter pertama) akan di-looping.
2. Di tiap perulangannya, closure callback dipanggil, dengan disisipkan data tiap elemen perulangan sebagai parameter.
3. Closure callback berisikan kondisi filtering, dengan hasil bertipe bool yang kemudian dijadikan nilai balik dikembalikan.
4. Di dalam fungsi filter() sendiri, ada proses seleksi kondisi (yang nilainya didapat dari hasil eksekusi closure callback). Ketika kondisinya bernilai true, maka data elemen yang sedang diulang dinyatakan lolos proses filtering.
5. Data yang lolos ditampung variabel result. Variabel tersebut dijadikan sebagai nilai balik fungsi filter().
*/

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o") // fungsi untuk mendeteksi apakah sebuah substring merupakan bagian dari string
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli: [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf o: [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
	// filter jumlah huruf 5: [jason ethan]
}


/*
* Alias skema closure
* untuk fungsi yang skemanya cukup panjang, akan lebih baik jika menggunakan alias dalam pendefinisiannya
* membuat alias fungsi berarti menjadikan skema fungsi tersebut menjadi tipe data baru
*/

/*
type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string {
    ...
}
*/
