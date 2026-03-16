package main

import "fmt"

func main() {

    // 1. Integer
    var age int = 21
    fmt.Println("Integer:", age)

    // 2. Float
    var price float64 = 19.99
    fmt.Println("Float:", price)

    // 3. Boolean
    var isStudent bool = true
    fmt.Println("Boolean:", isStudent)

    // 4. String
    var name string = "Golang"
    fmt.Println("String:", name)

    // 5. Array
    numbers := [3]int{1, 2, 3}
    fmt.Println("Array:", numbers)

    // 6. Slice
    fruits := []string{"Apple", "Banana", "Mango"}
    fmt.Println("Slice:", fruits)

    // 7. Map
    student := map[string]int{
        "Alice": 90,
        "Bob":   85,
    }
    fmt.Println("Map:", student)

    // 8. Struct
    type Person struct {
        Name string
        Age  int
    }

    p := Person{Name: "John", Age: 25}
    fmt.Println("Struct:", p)

    // 9. Byte
    var b byte = 65
    fmt.Println("Byte:", b)

    // 10. Rune
    var r rune = 'A'
    fmt.Println("Rune:", r)

    // 11. uintptr
    var x int = 10
    ptr := &x
    var address uintptr = uintptr(fmt.Sprintf("%p", ptr)[0])
    fmt.Println("uintptr example:", address)
}