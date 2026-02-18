package main

import "fmt"

func main() {
	
	// int
	var age int = 25
	// float64
	var price float64 = 99.99

	// string
	var name string = "Manoj"

	// bool
	var isValid bool = true

	// array (fixed size)
	var arr [3]int = [3]int{1, 2, 3}

	// slice (dynamic array)
	var nums []int = []int{10, 20, 30}

	// map
	var user map[string]int = map[string]int{
		"age": 25,
	}

	// struct
	type User struct {
		Name string
		Age  int
	}
	var u User = User{"Rahul", 30}

	// pointer
	var x int = 100
	var ptr *int = &x

	fmt.Println(
		age,
		price,
		name,
		isValid,
		arr,
		nums,
		user,
		u,
		ptr,
	)
}
