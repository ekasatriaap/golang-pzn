package main

import "fmt"

func main(){
	// langsung
	var result = 10 + 10
	fmt.Println(result)

	// masukkan dari variable
	var a = 10
	var b = 10
	var c = a * b // bisa dengan operasi lain seperti -, +, /, %
	fmt.Println(c)

	// assigment
	var i = 10
	i += 10 // i = i + 10, bisa juga dengan operator lain -=, *=, /=
	fmt.Println(i)

	// unary operator
	i++ // i = i + 1, operator lain --
	fmt.Println(i)
	var negative = -100
	var positive = +100 // default angka positif tidak perlu +
	fmt.Println(negative)
	fmt.Println(positive)
}