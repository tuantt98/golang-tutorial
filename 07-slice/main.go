package main

import "fmt"

func main() {
	// var mySlice []int
	// fmt.Println(mySlice)

	// var mySlice2 = []int{1, 2, 3, 4}
	// fmt.Println(mySlice2)

	// var array = [4]int{1, 2, 3, 4}
	// mySlice3 := array[1:]

	// mySlice3[0] = 10
	// fmt.Println(mySlice3)
	// fmt.Println(array)

	// var mySlice4 = []int{1, 2, 3, 4}
	// mySlice5 := mySlice4[2:3]

	// fmt.Println(len(mySlice5))
	// fmt.Println(cap(mySlice5))

	// slice := make([]int, 2, 5)
	// fmt.Println(slice)
	// fmt.Println(cap(slice))

	// slice2 := make([]int, 2)
	// slice2 = append(slice2, 100)
	// fmt.Println(slice2)

	// var slice3 []int
	// slice3 = append(slice3, 100)
	// fmt.Println(slice3)

	slice4 := make([]int, 2)

	fmt.Println(cap(slice4))
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	numberCopy := copy(slice4, src)

	fmt.Println(numberCopy)
	fmt.Println(slice4)

}
