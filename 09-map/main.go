package main

import "fmt"

// map là tham chiếu
func main() {

	// var myMap = make(map[string]int)
	// // fmt.Println(myMap)
	// if myMap == nil {
	// 	fmt.Println("myMap nil")
	// } else {
	// 	fmt.Println("myMap not nil")
	// }
	// var myMap1 map[string]int
	// if myMap1 == nil {
	// 	fmt.Println("myMap1 nil")
	// }

	// constructor
	a := "123"
	myMap2 := map[string]int{
		a:       1,
		"key-2": 2,
		"key-3": 3,
	}
	fmt.Println(myMap2)

	// add
	myMap2["key-4"] = 4
	myMap2["key-5"] = 5

	fmt.Println(myMap2)

	// delete
	delete(myMap2, "key-3")
	fmt.Println(myMap2)

	// length
	fmt.Println(len(myMap2))

	value, found := myMap2["key-3"]
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("Khong tim thay")
	}
}
