package main

import "fmt"

func main() {
	// Map: tipe data key-value --> key harus unik
	var chicken map[string]int // map[tipe data]tipe data
	chicken = map[string]int{} // inisialisasi nilai awal agar tidak error dengan menambahkan kurung kurawal ({})

	chicken["januari"] = 50
	chicken["februai"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// Inisialisasi nilai awal map
	/*
		var data map[string]int -> terjadi error karena bernilai nil
		data["one"] = 1
	*/
	var data = map[string]int{}
	data["one"] = 1

	// inisialisasi map cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	fmt.Println(chicken1)

	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40, // tambahkan koma di setiap elemen
	}
	fmt.Println(chicken2)

	// perulangan map menggunakan for-range
	var chicken3 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken3 {
		fmt.Println(key, " \t:", val)
	}

	// fungsi delete() digunakan untuk menghapus item dengan key tertentu pada map

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)

	delete(chicken, "januari") // delete(objek map, key)

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

	// deteksi keberadaan item dengan key
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item does not exists")
	}

	// kombinasi slice & map
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	/*
		bisa ditulis:
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},

		tiap elemen bisa memiliki key yg berbeda beda
		var data = []map[string]string{
			{"name": "chicken blue", "gender": "male", "color": "brown"},
			{"address": "mangga street", "id": "k001"},
			{"community": "chicken lovers"},
		}
	*/

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
