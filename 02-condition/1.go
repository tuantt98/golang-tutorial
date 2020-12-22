package main

import "fmt"

func main() {
	number := 10
	if number >= 10 {
		fmt.Println(number)
	} else {
		fmt.Println("Not ok")
	}

	if b := 10; b == 10 {
		fmt.Println(b)
	}
	fmt.Println("============================")
	SC(9)
	fmt.Println("============================")
	MFT(3)
}

func SC(number int) {
	// Nếu dùng toán tử so sánh thì sử dụng switch thôi -> không truyện tham số vào
	switch {
	// case 1:
	// 	fmt.Println("1")
	// case 2:
	// 	fmt.Println("2")
	case (number > 10):
		fmt.Println(">10")
	case (number < 10):
		fmt.Println("<10")
	default:
		fmt.Println("Not match")
	}
}

// fallthrough
func MFT(number int) {
	// fallthrough => thực hiện luôn các lệnh phía sau khi đã nhận dc đk đúng
	switch number {
	case 1:
		fmt.Println("=1")
		fallthrough
	case 2:
		fmt.Println("=2")
		fallthrough
	case 3:
		fmt.Println("=3")
		fallthrough
	default:
		fmt.Println("Not match")
	}
}
