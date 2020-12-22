package demo

import (
	"fmt"
	"go-module/sumnumber"
)

func MyPrintln(str string) {
	Gretting()
	fmt.Println(str)
}

func MyPrintln2() {
	fmt.Println("okok")
	fmt.Println("From demo:",sumnumber.SumFromTwoNumBer(1,1))
}


