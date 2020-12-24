package main

import "fmt"

func main() {
	// a := 100
	// var pointer *int
	// pointer = &a

	// fmt.Printf("%T\n", &pointer)
	// fmt.Println(&pointer)

	// b := 123
	// p2 := new(int)
	// p2 = &b
	// fmt.Println(p2)

	// c := 100
	// var pointer *int
	// pointer = &c
	// fmt.Println(c)
	// fmt.Println(pointer)
	// *pointer = 99

	// fmt.Println(c)
	// fmt.Println(pointer)
	const n int = 4
	array := [n]int{1, 2, 3}

	var pointer *[n]int

	pointer = &array

	fmt.Println(pointer)
	fmt.Printf("%T", pointer)

	pointer2 := new(int)
	b := 10
	pointer2 = &b

	changeValue(pointer2)
	fmt.Println(pointer2)
	fmt.Println(b)
}

func changeValue(pointer *int) {
	*pointer = 20
}
