package main

import "fmt"

func main(){
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Eyzel"
	names[2] = "Dreaska"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		100,
	}

	fmt.Println(values)

	// fungsi len() menghitung panjang array yg di deklarasikan, bukan jumlah data yang ada di array
	fmt.Println(len(names))
	fmt.Println(len(values))

}