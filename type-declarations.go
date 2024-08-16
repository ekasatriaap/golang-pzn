package main

import "fmt"

func main(){
	type noKTP string
	type married bool

	var no_ktp_eka noKTP = "1401033006950005"
	var status_eka married = true

	fmt.Println(no_ktp_eka)
	fmt.Println(status_eka)
}