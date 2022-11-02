package main

import "fmt"

func main() {
	// fmt.Println(sum(1, 2))
	// a, b := multi(1, 2)
	// fmt.Println(a)
	// fmt.Println(b)
	num1, num2, isSum, result := sumFor2Number(1, -1)

	if isSum {
		fmt.Println(result)
	} else {
		fmt.Println(num1)
		fmt.Println(num2)
	}
}

func sum(a int, b int) int {
	return (a + b)
}

// trả về nhiều giá trị
func multi(a, b int) (int, float32) {
	a++
	b++
	return a, float32(b) * 1.00
}

func sumFor2Number(a int, b int) (number1 int, number2 int, isSum bool, result int) {
	result = a + b
	if result != 0 {
		isSum = true
	} else {
		isSum = false
	}
	return a, b, isSum, result
}

func addTwoNumber(a int, b float32) float32 {
	return float32(a) + b
}
