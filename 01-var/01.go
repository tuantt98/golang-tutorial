package main

import "fmt"

func main() {
	var name string = "name"
	var age int = 22
	fmt.Println(name)
	fmt.Println(age)

	var (
		name1 string
		age1  int
	)
	name1 = "name1"
	age1 = 22
	fmt.Println(name1)
	fmt.Println(age1)

	var (
		name2 string = "name2"
		age2  int    = 22
	)
	fmt.Println(name2)
	fmt.Println(age2)

	name3, age3 := "name3", 22
	fmt.Println(name3)
	fmt.Println(age3)
}
