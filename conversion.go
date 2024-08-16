package main

import "fmt"

func main(){

	// number
	var number32 int32 = 10000
	var number64 int64 = int64(number32)
	var number16 int16 = int16(number32)
	// jika dilakukan konveri ke tipe data yang melewati batas, maka hitungan akan dimulai dari titik terendah
	var number8 int8 = int8(number32)
	// hasilnya ketika sudah sampai di 127, maka angka akan dihitung lagi dari -128

	fmt.Println(number8)
	fmt.Println(number16)
	fmt.Println(number32)
	fmt.Println(number64)

	// string
	var name string = "Eyzel"
	var e byte = name[0]
	var e_string = string(e)

	fmt.Println(e_string)
}