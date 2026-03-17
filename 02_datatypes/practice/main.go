package main

import "fmt"

func main() {
	// Primitive Data types
	var name string = "manoj"
	var age int = 12
	var isValid bool = true
	var price float32 = 12.3
	fmt.Println(name, age, isValid, price)

	// Non Primitive Datatypes
	arr := []int{1, 2, 3}
	students := []int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println(students)

	data := map[string]string{
		"name":  "Manoj Rawat",
		"email": "manoj@gmail.com",
	}
	fmt.Println(data)

	var b byte = 65
	fmt.Println(b)
	var r rune = 'A'
	fmt.Println("Rune:", r)

	var num int = 12
	p := &num
	fmt.Println(p)

	//
}
