package main

import "fmt"

func main(){
	var days = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jum'at",
		"Sabtu",
		"Minggu", 
	}
	var slice = days[3:6]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	// ketika mengupdate data di slice, maka array yg menjadi referensi ikut berubha
	slice[2] = "Sabtu update"
	fmt.Println(days)

	// append
	var slice2 = days[3:]
	// ketika menggunakan func append pada slice namun sudah melebihi kapasitas, maka slice akan membuat array baru
	var slice_append = append(slice2, "hari baru")
	// perubahan ini tidak akan merubah array days karena slice_append sudah membuat array baru
	slice_append[0] = "Kamis update"
	fmt.Println(slice_append)
	fmt.Println(days)

	// membuat slice baru
	var new_slice = make([]string, 2, 5)
	new_slice[0] = "Eyzel"
	new_slice[1] = "Dreaska"
	fmt.Println(new_slice)

	// copu slice
	var copy_slice = make([]string, len(new_slice), cap(new_slice))
	copy(copy_slice, new_slice)
	fmt.Println(copy_slice)

	// perbedaan penulisan array dan slice
	var ini_array = [...]int{2,3,4,1,4,1,1}
	var ini_slice = []int{3,2,2,1}
	fmt.Println(ini_array)
	fmt.Println(ini_slice)
}