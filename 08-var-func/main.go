package main

import (
	"fmt"
)

func main() {
	addItem(1, 2, 3, 4, 5, 6, 7, 8, 9)

	var list = []int{10, 11, 12, 13}
	addItem(14, list...)

	change(list...)
	fmt.Println(list)

	var list2 = [...]int{1, 2, 3, 4}
	result := sumArray(list2, 0)
	fmt.Println(result)
}

func addItem(item int, list ...int) {
	fmt.Printf("%T\n", list)
	list = append(list, item)
	fmt.Println(list)
}

func change(list ...int) {
	list[0] = 999
}

// Truyền vào 1 array thì nó phải cùng 1 số lượng phần tử
func sumArray(list [4]int, sum int) int {
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	return sum
}
