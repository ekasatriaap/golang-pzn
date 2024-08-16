package main

import "fmt"

func main(){
	var first_name string
	first_name = "Muhammad"
	fmt.Println(first_name)

	var middle_name = "Eyzel"
	fmt.Println(middle_name)

	last_name := "Dreaska"
	fmt.Println(last_name)

	var (
		nama_depan = "Muhammad"
		nama_tengah = "Eyzel"
		nama_belakang = "Dreaska"
	)
	fmt.Println(nama_depan)
	fmt.Println(nama_tengah)
	fmt.Println(nama_belakang)
}