package main

import (
	"fmt"
	"go-module/demo"
	"go-module/sumnumber"

	"github.com/leekchan/accounting"
)

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))

	demo.MyPrintln("kuro")
	demo.MyPrintln2()

	sumnumber.MyPrintln2()
	// result := sumnumber.SumFromTwoNumBer(1, 2)
	// fmt.Println("KQ: ", result)

}
