package main

import "fmt"

func main() {
	// var myArray [3]int
	// fmt.Println(myArray)

	// myArray2 := [3]int{1, 2, 3}
	// fmt.Println(myArray2)

	// fmt.Println(len(myArray2))

	myArray3 := [...]int{1, 2, 3}
	// fmt.Println(myArray3)

	for i := 0; i < len(myArray3); i++ {
		fmt.Println(myArray3[i])
	}

	for index, value := range myArray3 {
		fmt.Printf("Key: %d - Value: %d \n", index, value)
	}

	myArray4 := [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	fmt.Println(myArray4)

}
