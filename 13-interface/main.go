package main

import "fmt"

type Animal interface {
	speak()
	run()
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("gau gau")
}
func (d Dog) run() {
	fmt.Println("run run")
}

func main() {
	var dog Dog
	dog.speak()

	var animal Animal

	animal = Dog{}

	animal.speak()
	animal.run()
}
