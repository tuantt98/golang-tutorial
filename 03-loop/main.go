package main

import "fmt"

func main() {
	// fnFro()
	// fnForToWhile()
	result := sumTwoNumber(1, 2)
	fmt.Println(result)
}

func fnFro() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func fnForToWhile() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func sumTwoNumber(a int, b int) int {
	defer fmt.Println("ok")
	return (a + b)
}
