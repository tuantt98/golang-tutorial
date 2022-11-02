// func => 1 chức năng riêng
// method => thuộc về đối tượng nào

package main

import "fmt"

type Student struct {
	name string
}

func (s Student) getName() string {
	return s.name
}

// value receiver
// => không làm thay đổi giá trị
// vì nó tạo ra 1 bãn copy và thay đổi trên đó
func (s Student) changeName() {
	s.name = "kazuto"
}

// pointer receiver
func (s *Student) changeName2() {
	s.name = "kazuto"
}

func main() {
	st1 := &Student{name: "kuro"}
	// name := st1.getName()
	st1.changeName()
	fmt.Println("value receiver: ", st1.name)

	st1.changeName2()
	fmt.Println("pointer receiver: ", st1.name)
}
