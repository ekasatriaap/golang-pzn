package main

import "fmt"

func main(){
	var name_1 = "Eka"
	var name_2 = "Eyzel"
	var result bool = name_1 == name_2
	fmt.Println(result)

	var value_1 = 100
	var value_2 = 300

	fmt.Println(value_1 > value_2)
	fmt.Println(value_1 < value_2)
	fmt.Println(value_1 == value_2)
	fmt.Println(value_1 != value_2)
}