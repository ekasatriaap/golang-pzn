package main

import "fmt"

func main(){
	var person map[string]string = map[string]string{
		"name": "Eyzel",
		"address": "Pekanbaru",
	}
	person["title"] = "Baby"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	// membuat map baru
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eyzel"
	book["ups"] = "Error"
	fmt.Println(book)
	// menghapus key pada map
	delete(book, "ups")
	fmt.Println(book)
}