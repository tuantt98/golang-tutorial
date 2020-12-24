package main

import "fmt"

type StudentInfo struct {
	class string
	email string
	age   int
}

type Student struct {
	id   int
	name string
	info StudentInfo
}

type NoName struct {
	int
	string
}

func main() {
	// st1 := Student{
	// 	id:   1,
	// 	name: "kuro",
	// }
	// fmt.Println(st1)
	// fmt.Println(st1.id)
	// fmt.Println(st1.name)

	// st2 := Student{2, "kazuto"}
	// fmt.Println(st2)

	// poiter := &Student{3, "kkk"}
	// fmt.Println(poiter.id)
	// fmt.Println(poiter.name)

	st3 := NoName{4, "5"}
	fmt.Println(st3)

	st4 := Student{
		id:   5,
		name: "okok",
		info: StudentInfo{
			class: "A2",
			email: "admin@gmail.com",
			age:   22,
		},
	}
	st5 := Student{
		id:   5,
		name: "okok",
		info: StudentInfo{
			class: "A2",
			email: "admin@gmail.com",
			age:   22,
		},
	}
	fmt.Println(st4)
	fmt.Println(st4.info.class)
	fmt.Println(st4 == st5)

	var st6 Student
	fmt.Println(st6)
}
